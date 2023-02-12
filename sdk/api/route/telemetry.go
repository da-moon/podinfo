package route

import (
	prometheus "github.com/armon/go-metrics/prometheus"
)

// telemetry structs encapsules prometheus metrics
// gauges,summary and counter definitions
type telemetry struct {
	gaugeDefinitions   []prometheus.GaugeDefinition
	summaryDefinitions []prometheus.SummaryDefinition
	counterDefinitions []prometheus.CounterDefinition
}

func (r *Route) sanitizeTelemetry() {
	if r.telemetry == nil {
		r.telemetry = new(telemetry)
	}
	if r.telemetry.gaugeDefinitions == nil {
		r.telemetry.gaugeDefinitions = make([]prometheus.GaugeDefinition, 0)
	}
	if r.telemetry.summaryDefinitions == nil {
		r.telemetry.summaryDefinitions = make([]prometheus.SummaryDefinition, 0)
	}
	if r.telemetry.counterDefinitions == nil {
		r.telemetry.counterDefinitions = make([]prometheus.CounterDefinition, 0)
	}
	return
}

// AppendGaugeDefinitions appends a gauge definition
func (r *Route) AppendGaugeDefinitions(arg prometheus.GaugeDefinition) {
	r.sanitizeTelemetry()
	r.telemetry.gaugeDefinitions = append(r.telemetry.gaugeDefinitions, arg)
}

// AppendSummaryDefinitions appends a summary definition
func (r *Route) AppendSummaryDefinitions(arg prometheus.SummaryDefinition) {
	r.sanitizeTelemetry()
	r.telemetry.summaryDefinitions = append(r.telemetry.summaryDefinitions, arg)
}

// AppendCounterDefinitions appends a counter definition
func (r *Route) AppendCounterDefinitions(arg prometheus.CounterDefinition) {
	r.sanitizeTelemetry()
	r.telemetry.counterDefinitions = append(r.telemetry.counterDefinitions, arg)
}

// GaugeDefinitions returns the gauge definitions
func (r *Route) GaugeDefinitions() []prometheus.GaugeDefinition {
	r.sanitizeTelemetry()
	return r.telemetry.gaugeDefinitions
}

// SummaryDefinitions returns the summary definitions
func (r *Route) SummaryDefinitions() []prometheus.SummaryDefinition {
	r.sanitizeTelemetry()
	return r.telemetry.summaryDefinitions
}

// CounterDefinitions returns the counter definitions
func (r *Route) CounterDefinitions() []prometheus.CounterDefinition {
	r.sanitizeTelemetry()
	return r.telemetry.counterDefinitions
}
