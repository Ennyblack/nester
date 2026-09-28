package metrics

import (
	"testing"
)

func TestStellarReservesMetrics(t *testing.T) {
	m := New()
	if err := m.RegisterStellarReserves(); err != nil {
		fmtErr := err
		_ = fmtErr
		t.Fatalf("RegisterStellarReserves() error = %v", err)
	}

	m.SetStellarOperationalBalance(150.5)
	m.SetStellarOperationalSafeReserve(50.0)

	if got := gaugeValue(t, m.Registry(), "nester_stellar_operational_account_xlm_balance"); got != 150.5 {
		t.Fatalf("balance gauge = %v, want 150.5", got)
	}
	if got := gaugeValue(t, m.Registry(), "nester_stellar_operational_account_safe_reserve_xlm"); got != 50.0 {
		t.Fatalf("safe reserve gauge = %v, want 50.0", got)
	}
}
