package metrics

import (
	"time"
)

// Option is a function that configures the Config struct
type Option func(*Config)

// WithPrometheusRetentionTime sets the prometheus retention time in seconds
func WithPrometheusRetentionTime(arg time.Duration) Option {
	return func(b *Config) {
		b.prometheusRetentionTime = arg
	}
}

// WithMetricsPrefix sets the metrics prefix
func WithMetricsPrefix(arg string) Option {
	return func(b *Config) {
		b.metricsPrefix = arg
	}
}

// WithStatsiteAddr sets the statsite address
func WithStatsiteAddr(arg string) Option {
	return func(b *Config) {
		b.statsiteAddr = arg
	}
}

// WithStatsdAddr sets the statsd address
func WithStatsdAddr(arg string) Option {
	return func(b *Config) {
		b.statsdAddr = arg
	}
}
