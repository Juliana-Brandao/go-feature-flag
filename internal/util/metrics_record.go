package util

import (
	"github.com/prometheus/client_golang/prometheus"
)

type MetricsRecord interface {
	ResetGaugeFeatureFlag()
	WithLabelValues(flagName string, status string)
}

type metricsRecord struct {
	gaugeFeatureFlag *prometheus.GaugeVec
}

func (m *metricsRecord) WithLabelValues(flagName string, status string) {
	m.gaugeFeatureFlag.WithLabelValues(flagName, status).Set(1)
}

func (m *metricsRecord) ResetGaugeFeatureFlag() {
	m.gaugeFeatureFlag.Reset()
}

func NewMetricsRecord() MetricsRecord {
	flagGauge := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "feature_flag_status",
			Help: "Status das feature flags no cache local da instância (1 para ativo, 0 para inativo)",
		},
		[]string{"flag_name", "status"},
	)

	prometheus.MustRegister(flagGauge)

	return &metricsRecord{
		gaugeFeatureFlag: flagGauge,
	}
}
