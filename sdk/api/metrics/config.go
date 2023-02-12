package metrics

import (
	"os"
	"sync"
	"time"

	logger "github.com/da-moon/northern-labs-interview/internal/logger"
	stacktrace "github.com/palantir/stacktrace"
)

// Config is the configuration for the metrics package.
//
//go:generate gomodifytags -override -file $GOFILE -struct Config -add-tags json,yaml,mapstructure -w -transform snakecase
type Config struct {
	mutex                   sync.RWMutex          `json:"lock" yaml:"lock" mapstructure:"lock"`                                                                //nolint:govet // a getter function is defined for this field
	log                     *logger.WrappedLogger `json:"log" yaml:"log" mapstructure:"log"`                                                                   //nolint:govet // a getter function is defined for this field
	metricsPrefix           string                `json:"metrics_prefix" yaml:"metrics_prefix" mapstructure:"metrics_prefix"`                                  //nolint:govet // a getter function is defined for this field
	statsiteAddr            string                `json:"statsite_addr" yaml:"statsite_addr" mapstructure:"statsite_addr"`                                     //nolint:govet // a getter function is defined for this field
	statsdAddr              string                `json:"statsd_addr" yaml:"statsd_addr" mapstructure:"statsd_addr"`                                           //nolint:govet // a getter function is defined for this field
	prometheusRetentionTime time.Duration         `json:"prometheus_retention_time" yaml:"prometheus_retention_time" mapstructure:"prometheus_retention_time"` //nolint:govet // a getter function is defined for this field
}

// New returns a new metrics exporter configuration
func New(log *logger.WrappedLogger, opts ...Option) (*Config, error) {
	if log == nil {
		err := stacktrace.NewError("no logger was provided")
		return nil, err
	}
	result := &Config{
		log:                     log,
		metricsPrefix:           DefaultMetricsPrefix(),
		statsiteAddr:            DefaultStatsiteAddr(),
		statsdAddr:              DefaultStatsdAddr(),
		prometheusRetentionTime: DefaultPrometheusRetentionTime(),
	}
	for _, opt := range opts {
		opt(result)
	}
	return result, nil
}

// ──────────────────────────────────────────────────────── I ──────────
//
//	:::::: D E F A U L T S : :  :   :    :     :        :          :
//
// ──────────────────────────────────────────────────────────────────

// DefaultMetricsPrefix returns the default metrics prefix
func DefaultMetricsPrefix() string {
	result := os.Getenv("PODINFO_METRICS_PREFIX")
	if result == "" {
		result = "podinfo_api"
	}
	return result
}

// DefaultStatsiteAddr returns the default statsite address
func DefaultStatsiteAddr() string {
	return os.Getenv("PODINFO_STATSITE_ADDR")
}

// DefaultStatsdAddr returns the default statsd address
func DefaultStatsdAddr() string {
	return os.Getenv("PODINFO_STATSD_ADDR")
}

// DefaultPrometheusRetentionTime returns the default prometheus retention time in seconds
func DefaultPrometheusRetentionTime() time.Duration {
	result := time.Second * 60
	resultStr := os.Getenv("PODINFO_PROMETHEUS_RETENTION_TIME")
	if resultStr == "" {
		parsed, err := time.ParseDuration(resultStr)
		if err == nil {
			result = parsed
		}
	}
	return result
}

// ────────────────────────────────────────────────────── I ──────────
//
//	:::::: S E T T E R S : :  :   :    :     :        :          :
//
// ────────────────────────────────────────────────────────────────

// SetMetricsPrefix sets the metrics prefix
func (c *Config) SetMetricsPrefix(arg string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.metricsPrefix = arg
}

// SetStatsiteAddr sets the statsite address
func (c *Config) SetStatsiteAddr(arg string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.statsiteAddr = arg
}

// SetStatsdAddr sets the statsd address
func (c *Config) SetStatsdAddr(arg string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.statsdAddr = arg
}

// SetPrometheusRetentionTime sets the prometheus retention time in seconds
func (c *Config) SetPrometheusRetentionTime(arg time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.prometheusRetentionTime = arg
}

// SetLog sets the logger
func (c *Config) SetLog(arg *logger.WrappedLogger) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.log = arg
}

//
// ────────────────────────────────────────────────────── I ──────────
//   :::::: G E T T E R S : :  :   :    :     :        :          :
// ────────────────────────────────────────────────────────────────
//

// MetricsPrefix returns the metrics prefix
func (c *Config) MetricsPrefix() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.metricsPrefix
}

// StatsiteAddr returns the statsite address
func (c *Config) StatsiteAddr() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.statsiteAddr
}

// StatsdAddr returns the statsd address
func (c *Config) StatsdAddr() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.statsdAddr
}

// PrometheusRetentionTime returns the prometheus retention time in seconds
func (c *Config) PrometheusRetentionTime() time.Duration {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.prometheusRetentionTime
}

// Log returns the logger
func (c *Config) Log() *logger.WrappedLogger {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.log
}

// ──────────────────────────────────────────────────────────── I ──────────
//
//	:::::: V A L I D A T I O N : :  :   :    :     :        :          :
//
// ──────────────────────────────────────────────────────────────────────
//

// Validate validates the configuration
func (c *Config) Validate() error {
	c.log.Info("metrics : validating configuration")
	var err error
	err = c.ValidateMetricsPrefix()
	if err != nil {
		err = stacktrace.Propagate(err, "could not validate metrics config")
		return err
	}
	c.log.Info("metrics : 'MetricsPrefix' configuration value was successfully validated")
	err = c.ValidateStatsiteAddr()
	if err != nil {
		err = stacktrace.Propagate(err, "could not validate metrics config")
		return err
	}
	c.log.Info("metrics : 'StatsiteAddr' configuration value was successfully validated")
	err = c.ValidateStatsdAddr()
	if err != nil {
		err = stacktrace.Propagate(err, "could not validate metrics config")
		return err
	}
	c.log.Info("metrics : 'StatsdAddr' configuration value was successfully validated")
	err = c.ValidatePrometheusRetentionTime()
	if err != nil {
		err = stacktrace.Propagate(err, "could not validate metrics config")
		return err
	}
	c.log.Info("metrics : 'PrometheusRetentionTime' configuration value was successfully validated")
	return nil
}

// ValidateMetricsPrefix validates the MetricsPrefix configuration value
func (c *Config) ValidateMetricsPrefix() error {
	c.log.Info("metrics : validating 'MetricsPrefix' configuration value")
	return nil
}

// ValidateStatsiteAddr validates the StatsiteAddr configuration value
func (c *Config) ValidateStatsiteAddr() error {
	c.log.Info("metrics : validating 'StatsiteAddr' configuration value")
	return nil
}

// ValidateStatsdAddr validates the StatsdAddr configuration value
func (c *Config) ValidateStatsdAddr() error {
	c.log.Info("metrics : validating 'StatsdAddr' configuration value")
	return nil
}

// ValidatePrometheusRetentionTime validates the PrometheusRetentionTime configuration value
func (c *Config) ValidatePrometheusRetentionTime() error {
	c.log.Info("metrics : validating 'PrometheusRetentionTime' configuration value")
	if c.prometheusRetentionTime.Nanoseconds() < 1 {
		err := stacktrace.NewError(
			"prometheus expiry duration must be greater than 1 Nanosecond",
		)
		return err
	}
	return nil
}
