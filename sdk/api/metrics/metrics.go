package metrics

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	metricscollector "github.com/armon/go-metrics"
	prometheuscollector "github.com/armon/go-metrics/prometheus"
	logger "github.com/da-moon/northern-labs-interview/internal/logger"
	response "github.com/da-moon/northern-labs-interview/sdk/api/response"
	stacktrace "github.com/palantir/stacktrace"
	prometheus "github.com/prometheus/client_golang/prometheus"
	promhttp "github.com/prometheus/client_golang/prometheus/promhttp"
)

// Metrics struct represents
// a server core engine that
// collects Metrics and stores it in
// various sinks.
type Metrics struct {
	mutex          sync.RWMutex
	once           sync.Once
	initialized    bool
	fanout         metricscollector.FanoutSink
	inmem          *metricscollector.InmemSink
	config         *Config
	log            *logger.WrappedLogger
	collector      *metricscollector.Metrics
	prometheusOpts *prometheuscollector.PrometheusOpts
}

// New is the Metrics object constructor
func (c *Config) New() (*Metrics, error) {
	err := c.Validate()
	if err != nil {
		err = stacktrace.Propagate(err, "could not create Metrics object")
		return nil, err
	}
	result := &Metrics{
		config: c,
		fanout: make([]metricscollector.MetricSink, 0),
		log:    c.Log(),
		prometheusOpts: &prometheuscollector.PrometheusOpts{
			Expiration: c.PrometheusRetentionTime(),
		},
	}
	result.sanitizePrometheusOpts()
	return result, nil
}

// Init starts all Metrics collectors
// TODO: simplify this function. split it into constructor and init
// nolint:gocognit,gocyclo // I can't think of a way to lower complexity
func (m *Metrics) Init(ctx context.Context) error {
	// ─── VALIDATION ─────────────────────────────────────────────────────────────────
	if m.prometheusOpts.Expiration.Nanoseconds() < 1 {
		err := stacktrace.NewError(
			"prometheus expiry duration must be greater than 1 Nanosecond",
		)
		return err
	}

	cErr := make(chan error, 1)
	m.once.Do(func() {
		if m.config.StatsiteAddr() != "" {
			err := m.Statsite()
			if err != nil {
				cErr <- err
				return
			}
			m.log.Info("statsite Metrics collector core engine successfully initialized")
		}
		if m.config.StatsdAddr() != "" {
			err := m.Statsd()
			if err != nil {
				cErr <- err
				return
			}
			m.log.Info("statsd Metrics collector core engine successfully initialized")
		}
		if m.prometheusOpts != nil {
			err := m.Prometheus()
			if err != nil {
				cErr <- err
				return
			}
			m.log.Info("prometheus Metrics collector core engine successfully initialized")
		}
	})
	select {
	case err, ok := <-cErr:
		if ok {
			if err != nil {
				return err
			}
		}
	case <-ctx.Done():
		errMsg := "Metrics engine initialization timed out"
		return stacktrace.Propagate(ctx.Err(), errMsg)

	default:
	}

	conf := metricscollector.DefaultConfig(m.config.MetricsPrefix())
	sinks := m.Sinks()
	// [ NOTE ] at least enable in-mem Metrics sink
	if len(sinks) == 0 {
		conf.EnableHostname = false
	}
	m.inmem = metricscollector.NewInmemSink(10*time.Second, time.Minute)
	sinks = append(
		sinks,
		m.inmem,
	)
	// TODO: maybe this is not needed
	m.fanout = sinks
	// TODO: see if collector can be removed
	collector, err := metricscollector.NewGlobal(conf, m.fanout)
	if err != nil {
		return stacktrace.Propagate(
			err,
			"failed to initialize Metrics collector engine",
		)
	}
	if collector == nil {
		err = stacktrace.NewError("no Metrics sink were specified")
		return err
	}
	m.collector = collector
	m.initialized = true
	return nil
}

// Handler returns a http.Handler that serves metrics.
func (m *Metrics) Handler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if m.config.PrometheusRetentionTime().Nanoseconds() < 1 {
			m.inmemHandler()(w, r)
			return
		}
		if enablePrometheusOutput(r) {
			m.prometheusHandler()(w, r)
			return
		}
		m.inmemHandler()(w, r)
		return
	}
}

// PrometheusHandler returns an http handler that exports prometheus compatible metrics
func (m *Metrics) prometheusHandler() func(http.ResponseWriter, *http.Request) {
	handlerOptions := promhttp.HandlerOpts{
		ErrorLog:      m.log.Unwrap(),
		ErrorHandling: promhttp.ContinueOnError,
	}
	handler := promhttp.HandlerFor(prometheus.DefaultGatherer, handlerOptions)
	return handler.ServeHTTP
}

// inmemHandler returns an http handler for non-prometheus ( inmem ) sink
func (m *Metrics) inmemHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := m.inmem.DisplayMetrics(w, r)
		if err != nil {
			payload := response.ErrInternalServerError()
			response.WriteErrorJSON(w, r, payload, fmt.Sprintf("root cause : %s", err.Error()))
			return
		}
		response.WriteSuccessfulJSON(w, r, resp)
		return
	}
}

// enablePrometheusOutput will look for Prometheus mime-type or format Query parameter
func enablePrometheusOutput(r *http.Request) bool {
	format := r.URL.Query().Get("format")
	// curl 127.0.0.1:2048/metrics?format=prometheus
	if strings.EqualFold(strings.ToLower(format), "prometheus") {
		return true
	}
	// checking to see if mime type is Prometheus-compatible
	acceptHeader := r.Header.Get("Accept")
	mimeTypes := strings.Split(acceptHeader, ",")
	for _, v := range mimeTypes {
		mimeInfo := strings.Split(v, ";")
		if len(mimeInfo) > 0 {
			rawMime := strings.ToLower(strings.Trim(mimeInfo[0], " "))
			if rawMime == "application/openmetrics-text" {
				return true
			}
			if rawMime == "text/plain" && (len(mimeInfo) > 1 && strings.Trim(mimeInfo[1], " ") == "version=0.4.0") {
				return true
			}
		}
	}
	return false
}

//
// ──────────────────────────────────────────────────────────── I ──────────
//   :::::: S I N K   S E T U P : :  :   :    :     :        :          :
// ──────────────────────────────────────────────────────────────────────
//

// Statsite function starts up statsite Metrics sink
func (m *Metrics) Statsite() error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if m.config.StatsiteAddr() == "" {
		return nil
	}
	addr := m.config.StatsdAddr()
	m.log.Info("Starting statsite Metrics collector core engine")
	if m.initialized {
		err := stacktrace.NewError("Metrics has already been initialized")
		return err
	}
	result, err := metricscollector.NewStatsiteSink(addr)
	if err != nil {
		err = stacktrace.Propagate(
			err,
			"failed to start statsite collector core engine",
		)
		return err
	}
	m.fanout = append(m.fanout, result)
	return nil
}

// Statsd starts up statsd Metrics sink
func (m *Metrics) Statsd() error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if m.config.StatsdAddr() != "" {
		return nil
	}
	m.log.Info("Starting statsd Metrics collector core engine")
	if m.initialized {
		err := stacktrace.NewError("Metrics has already been initialized")
		return err
	}
	addr := m.config.StatsdAddr()
	result, err := metricscollector.NewStatsdSink(addr)
	if err != nil {
		err = stacktrace.Propagate(
			err,
			"failed to start statsd collector core engine",
		)
		return err
	}
	m.fanout = append(m.fanout, result)
	return nil
}

// Prometheus starts up prometheus Metrics sink
func (m *Metrics) Prometheus() error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if m.prometheusOpts == nil {
		return nil
	}
	m.log.Info("Starting prometheus Metrics collector core engine")
	if m.initialized {
		err := stacktrace.NewError("Metrics has already been initialized")
		return err
	}
	// NOTE this code path should never get triggered
	if m.prometheusOpts.Expiration.Nanoseconds() < 1 {
		err := stacktrace.NewError(
			"prometheus expiry duration must be greater than 1 Nanosecond",
		)
		return err
	}

	result, err := prometheuscollector.NewPrometheusSinkFrom(*m.prometheusOpts)
	if err != nil {
		err = stacktrace.Propagate(
			err,
			"failed to start prometheus collector core engine",
		)
		return err
	}
	m.fanout = append(m.fanout, result)
	return nil
}

// Sinks returns the list of metrics sinks
func (m *Metrics) Sinks() metricscollector.FanoutSink {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.fanout
}

// ─── PROMETHEUS OPTS ────────────────────────────────────────────────────────────
func (m *Metrics) sanitizePrometheusOpts() {
	if m.prometheusOpts == nil {
		m.prometheusOpts = new(prometheuscollector.PrometheusOpts)
	}
	if m.prometheusOpts.GaugeDefinitions == nil {
		m.prometheusOpts.GaugeDefinitions = make([]prometheuscollector.GaugeDefinition, 0)
	}
	if m.prometheusOpts.SummaryDefinitions == nil {
		m.prometheusOpts.SummaryDefinitions = make([]prometheuscollector.SummaryDefinition, 0)
	}
	if m.prometheusOpts.CounterDefinitions == nil {
		m.prometheusOpts.CounterDefinitions = make([]prometheuscollector.CounterDefinition, 0)
	}
	return
}

// AppendGaugeDefinitions appends counter definitions to the prometheus collector
func (m *Metrics) AppendGaugeDefinitions(arg ...prometheuscollector.GaugeDefinition) error {
	if m.initialized {
		err := stacktrace.NewError("cannot append gauge definitions as metrics has already been initialized")
		return err
	}
	m.sanitizePrometheusOpts()
	if len(arg) > 0 {
		m.prometheusOpts.GaugeDefinitions = append(m.prometheusOpts.GaugeDefinitions, arg...)
	}
	return nil
}

// AppendSummaryDefinitions appends summary definitions to the prometheus collector
func (m *Metrics) AppendSummaryDefinitions(arg ...prometheuscollector.SummaryDefinition) error {
	if m.initialized {
		err := stacktrace.NewError("cannot append summary definitions as metrics has already been initialized")
		return err
	}

	m.sanitizePrometheusOpts()
	if len(arg) > 0 {
		m.prometheusOpts.SummaryDefinitions = append(m.prometheusOpts.SummaryDefinitions, arg...)
	}
	return nil
}

// AppendCounterDefinitions appends counter definitions to the prometheus collector
func (m *Metrics) AppendCounterDefinitions(arg ...prometheuscollector.CounterDefinition) error {
	if m.initialized {
		err := stacktrace.NewError("cannot append counter definitions as metrics has already been initialized")
		return err
	}
	m.sanitizePrometheusOpts()
	if len(arg) > 0 {
		m.prometheusOpts.CounterDefinitions = append(m.prometheusOpts.CounterDefinitions, arg...)
	}
	return nil
}
