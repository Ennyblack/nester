package reconciliation

import (
	"context"
	"log/slog"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type mockAlerter struct {
	criticalCalled bool
	warningCalled  bool
}

func (m *mockAlerter) CriticalFinding(_ context.Context, _ Finding) error {
	m.criticalCalled = true
	return nil
}

func (m *mockAlerter) WarningFinding(_ context.Context, _ Finding) error {
	m.warningCalled = true
	return nil
}

type captureHandler struct {
	records *[]string
}

func (h captureHandler) Enabled(context.Context, slog.Level) bool { return true }
func (h captureHandler) Handle(_ context.Context, r slog.Record) error {
	var sb strings.Builder
	sb.WriteString(r.Level.String() + " " + r.Message)
	r.Attrs(func(a slog.Attr) bool {
		sb.WriteString(" " + a.Key + "=" + a.Value.String())
		return true
	})
	*h.records = append(*h.records, sb.String())
	return nil
}
func (h captureHandler) WithAttrs([]slog.Attr) slog.Handler { return h }
func (h captureHandler) WithGroup(string) slog.Handler      { return h }

func TestPagerAlerterMainnetExceedsThreshold(t *testing.T) {
	var logs []string
	logger := slog.New(captureHandler{records: &logs})
	base := &mockAlerter{}

	// Threshold $100.00 -> 1,000,000,000 stroops or decimal 150
	pager := NewPagerAlerter(base, logger, true, 100.0)

	diff := decimal.NewFromFloat(150.0)
	finding := Finding{
		ID:         uuid.New(),
		RunID:      uuid.New(),
		Level:      LevelBalance,
		Type:       TypeMismatch,
		Severity:   SeverityCritical,
		EntityType: "vault",
		EntityID:   "vault-mainnet",
		Difference: &diff,
		ObservedAt: time.Now(),
	}

	if err := pager.CriticalFinding(context.Background(), finding); err != nil {
		T.Fatalf("unexpected error: %v", err)
	}

	if !base.criticalCalled {
		T.Fatal("expected base CriticalFinding to be called")
	}

	foundPageLog := false
	for _, logLine := range logs {
		if strings.Contains(logLine, "RECONCILIATION PAGER: immediate on-call page") {
			foundPageLog = true
		}
	}
	if !foundPageLog {
		T.Fatalf("expected immediate page log entry, got logs: %v", logs)
	}
}

func TestPagerAlerterTestnetDoesNotPage(t *testing.T) {
	var logs []string
	logger := slog.New(captureHandler{records: &logs})
	base := &mockAlerter{}

	// Testnet (isMainnet = false)
	pager := NewPagerAlerter(base, logger, false, 100.0)

	diff := decimal.NewFromFloat(150.0)
	finding := Finding{
		ID:         uuid.New(),
		RunID:      uuid.New(),
		Level:      LevelBalance,
		Type:       TypeMismatch,
		Severity:   SeverityCritical,
		EntityType: "vault",
		EntityID:   "vault-testnet",
		Difference: &diff,
		ObservedAt: time.Now(),
	}

	if err := pager.CriticalFinding(context.Background(), finding); err != nil {
		T.Fatalf("unexpected error: %v", err)
	}

	for _, logLine := range logs {
		if strings.Contains(logLine, "RECONCILIATION PAGER:") {
			T.Fatalf("did not expect pager log on testnet, got %q", logLine)
		}
	}
}
