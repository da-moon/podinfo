package core

import (
	"context"
	"net/http"
	"strings"

	metricscollector "github.com/armon/go-metrics"
	prometheuscollector "github.com/armon/go-metrics/prometheus"
	middlewares "github.com/da-moon/podinfo/api/middlewares"
	buildver "github.com/da-moon/podinfo/build/go/version"
	metrics "github.com/da-moon/podinfo/sdk/api/metrics"
	route "github.com/da-moon/podinfo/sdk/api/route"
	stacktrace "github.com/palantir/stacktrace"
)

// initMetrics must be called after initRoutes
// TODO: add boolean flags for routes and ensure routes have been initialized prior to metrics
func (c *Config) Telemetry(routes *route.Collection) error {
	c.log.Info("initializing telementry")
	var err error
	conf, err := metrics.New(c.log)
	if err != nil {
		err = stacktrace.Propagate(err, "could not initialize telemetry")
		return err
	}
	if c.MetricsPrefix != "" {
		conf.SetMetricsPrefix(c.MetricsPrefix)
	}
	if c.StatsiteAddr != "" {
		conf.SetStatsiteAddr(c.StatsiteAddr)
	}
	if c.StatsdAddr != "" {
		conf.SetStatsdAddr(c.StatsdAddr)
	}
	if c.PrometheusRetentionTime.Nanoseconds() >= 1 {
		conf.SetPrometheusRetentionTime(c.PrometheusRetentionTime)
	}
	err = conf.Validate()
	if err != nil {
		err = stacktrace.Propagate(err, "could not initialize telemetry")
		return err
	}
	m, err := conf.New()
	if err != nil {
		err = stacktrace.Propagate(err, "could not initialize telemetry")
		return err
	}
	// NOTE this code path should never be hit, but if it is, we want to fail fast.
	if m == nil {
		err = stacktrace.NewError("telemetry returned nil")
		return err
	}
	// version gauge
	err = m.AppendGaugeDefinitions(c.VersionGauge())
	if err != nil {
		err = stacktrace.Propagate(err, "could not initialize telemetry")
		return err
	}
	// extracting route specific Telemetry
	gaugeDefinition := routes.GaugeDefinitions()
	if gaugeDefinition != nil && len(gaugeDefinition) > 0 {
		err = m.AppendGaugeDefinitions(gaugeDefinition...)
		if err != nil {
			err = stacktrace.Propagate(err, "could not initialize telemetry")
			return err
		}
	}
	summaryDefinitions := routes.SummaryDefinitions()
	if summaryDefinitions != nil && len(summaryDefinitions) > 0 {
		err = m.AppendSummaryDefinitions(summaryDefinitions...)
		if err != nil {
			err = stacktrace.Propagate(err, "could not initialize telemetry")
			return err
		}
	}
	counterDefinitions := routes.CounterDefinitions()
	if counterDefinitions != nil && len(counterDefinitions) > 0 {
		err = m.AppendCounterDefinitions(counterDefinitions...)
		if err != nil {
			err = stacktrace.Propagate(err, "could not initialize telemetry")
			return err
		}
	}
	// TODO: inject ctx from caller
	ctx := context.Background()
	err = m.Init(ctx)
	if err != nil {
		err = stacktrace.Propagate(err, "could not initialize telemetry")
		return err
	}
	// ─── INITIALIZING ROUTE ─────────────────────────────────────────────────────────
	name := "metrics"
	path := "/metrics"
	r := route.New()
	r.SetName(name)
	r.SetPath(path)
	r.SetMethod(http.MethodGet)
	r.SetHandlerFunc(m.Handler())
	r.AppendMiddleware(middlewares.Log(c.log))
	routes.AppendRoute("", *r)
	return nil
}

//
// ──────────────────────────────────────────────────────────────────── I ──────────
//   :::::: G L O B A L   M E T R I C S : :  :   :    :     :        :          :
// ──────────────────────────────────────────────────────────────────────────────
//

// VersionGauge call SetGaugeWithLabels outside
func (c *Config) VersionGauge() prometheuscollector.GaugeDefinition {
	c.log.Info("Setting version gauge")
	labels := make([]metricscollector.Label, 0)
	version := strings.TrimPrefix(buildver.Version, "v")

	if version != "" {
		label := metricscollector.Label{Name: "version", Value: version}
		labels = append(labels, label)
	}
	revision := buildver.Revision
	if revision != "" {
		label := metricscollector.Label{Name: "revision", Value: revision}
		labels = append(labels, label)
	}
	metricscollector.SetGaugeWithLabels([]string{"version"}, 1, labels)
	return prometheuscollector.GaugeDefinition{
		Name: []string{"version"},
		Help: "Represents Podinfo-API build version.",
	}
}
