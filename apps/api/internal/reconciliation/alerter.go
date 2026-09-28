package reconciliation

import (
	"time"
	"context"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Level represents the category or severity tier of a reconciliation finding.
type Level string

const (
	LevelBalance     Level = "balance"
	LevelTransaction Level = "transaction"
	LevelInvariant   Level = "invariant"
)

// Type represents the kind of reconciliation finding.
type Type string

const (
	TypeMismatch Type = "mismatch"
)

// Severity represents the urgency of a finding.
type Severity string

const (
	SeverityCritical Severity = "critical"
	SeverityWarning  Severity = "warning"
)

// Finding represents a single discrepancy found during reconciliation.
type Finding struct {
	ID            uuid.UUID
	RunID         uuid.UUID
	Level         Level
	Type          Type
	Severity      Severity
	EntityType    string
	EntityID      string
	RecordedValue *decimal.Decimal
	OnChainValue  *decimal.Decimal
	Difference    *decimal.Decimal
	Details       map[string]string
	ObservedAt    time.Time
}

// Alerter abstracts the notification mechanism for findings.
type Alerter interface {
	CriticalFinding(ctx context.Context, finding Finding) error
	WarningFinding(ctx context.Context, finding Finding) error
}
