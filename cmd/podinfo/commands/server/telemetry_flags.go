package server

import (
	"time"

	flagset "github.com/da-moon/podinfo/internal/cli/flagset"
	value "github.com/da-moon/podinfo/internal/cli/value"
	"github.com/da-moon/podinfo/sdk/api/metrics"
)

type TelemetryFlags struct {
	*flagset.FlagSet
	metricsPrefix           value.String
	statsiteAddr            value.String
	statsdAddr              value.String
	prometheusRetentionTime value.Duration
}

func (f *TelemetryFlags) init() {
	f.FlagSet = flagset.New("Telemetry", "")
	f.Var(&f.metricsPrefix, "metrics-prefix",
		[]string{
			"flag used to set default metrics prefix.",
			"This can also be specified via the 'PODINFO_METRICS_PREFIX' env variable",
		})
	f.Var(&f.statsdAddr, "statsd-addr",
		[]string{
			"flag used to set statsd address",
			"This can also be specified via the 'STATSD_ADDR' env variable",
		})
	f.Var(&f.statsiteAddr, "statsite-addr",
		[]string{
			"flag used to set statsite address.",
			"This can also be specified via the 'STATSITE_ADDR' env variable",
		})
	f.Var(&f.prometheusRetentionTime, "prometheus-retention-time",
		[]string{
			"flag used to set prometheus retention time.",
			"This can also be specified via the 'PODINFO_PROMETHEUS_RETENTION_TIME' env variable",
		})
}
func (f *TelemetryFlags) MetricsPrefix() string {
	result := f.metricsPrefix.Get()
	if result == "" {
		result = metrics.DefaultMetricsPrefix()
	}
	return result
}
func (f *TelemetryFlags) StatsiteAddr() string {
	result := f.statsiteAddr.Get()
	if result == "" {
		result = metrics.DefaultStatsiteAddr()
	}
	return result
}
func (f *TelemetryFlags) StatsdAddr() string {
	result := f.statsdAddr.Get()
	if result == "" {
		result = metrics.DefaultStatsdAddr()
	}
	return result
}
func (f *TelemetryFlags) PrometheusRetentionTime() time.Duration {
	result := f.prometheusRetentionTime.Get()
	if result == 0 {
		result = metrics.DefaultPrometheusRetentionTime()
	}
	return result
}
