package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	StellarOperationalAccountBalance = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: Namespace,
		Subsystem: "stellar",
		Name:      "operational_account_xlm_balance",
		Help:      "Current XLM balance of the operational account used for sponsorship and reserves.",
	})

	StellarOperationalAccountSafeReserve = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: Namespace,
		subsystem: "stellar",
		Name:      "operational_account_safe_reserve_xlm",
		Help:      "Configured safe reserve XLM threshold below which an alert is triggered.",
	})
)

func (m *Metrics) RegisterStellarReserves() error {
	if m == nil || m.registry == nil {
		return nil
	}
	if err := m.registry.Register(StellarOperationalAccountBalance); err != nil {
		return err
	}
	return m.registry.Register(StellarOperationalAccountSafeReserve)
}

func (m *Metrics) SetStellarOperationalBalance(balance float64) {
	if m != nil {
		StellarOperationalAccountBalance.Set(balance)
	}
}

func (m *Metrics) SetStellarOperationalSafeReserve(reserve float64) {
	if m != nil {
		StellarOperationalAccountSafeReserve.Set(reserve)
	}
}
