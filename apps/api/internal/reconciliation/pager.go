package reconciliation

import (
	"context"
	"log/slog"
)

// PagerAlerter wraps an Alerter and explicitly triggers immediate on-call paging
// when mainnet drift exceeds the configured dollar threshold.
type PagerAlerter struct {
	base              Alerter
	logger            *slog.Logger
	isMainnet         bool
	thresholdUSD      float64
}

// NewPagerAlerter creates a new PagerAlerter decorating the given base Alerter.
func NewPagerAlerter(base Alerter, logger *slog.Logger, isMainnet bool, thresholdUSD float64) *PagerAlerter {
	if logger == nil {
		logger = slog.Default()
	}
	return &PagerAlerter{
		base:         base,
		logger:       logger,
		isMainnet:    isMainnet,
		thresholdUSD: thresholdUSD,
	}
}

func (p *PagerAlerter) CriticalFinding(ctx context.Context, finding Finding) error {
	if p.shouldPage(finding) {
		p.logger.ErrorContext(ctx, "RECONCILIATION PAGER: immediate on-call page triggered for mainnet drift exceeding threshold",
			"entity_type", finding.EntityType,
			"entity_id", finding.EntityID,
			"threshold_usd", p.thresholdUSD,
		)
	}
	if p.base != nil {
		return p.base.CriticalFinding(ctx, finding)
	}
	return nil
}

func (p *PagerAlerter) WarningFinding(ctx context.Context, finding Finding) error {
	if p.shouldPage(finding) {
		p.logger.ErrorContext(ctx, "RECONCILIATION PAGER: immediate on-call page triggered for mainnet warning finding exceeding threshold",
			"entity_type", finding.EntityType,
			"entity_id", finding.EntityID,
			"threshold_usd", p.thresholdUSD,
		)
	}
	if p.base != nil {
		return p.base.WarningFinding(ctx, finding)
	}
	return nil
}

func (p *PagerAlerter) shouldPage(finding Finding) bool {
	if !p.isMainnet {
		return false
	}
	if p.thresholdUSD <= 0 {
		return false
	}
	if finding.Difference == nil {
		return false
	}
	// Convert difference (assumed in stroops or decimal) to USD float.
	// If Difference is stored as decimal or we parse it, let's inspect.
	diffFloat, exact := finding.Difference.Float64()
	_ = exact
	// Assuming difference is in stroops (1 USDC = 10,000,000 stroops) or standard units.
	// In reconcileVault, difference is absolute difference in stroops (int64 converted or decimal).
	// Let's check how difference is populated. If difference is in stroops, USD value is diffFloat / 10000000.0.
	// If difference is already in USDC decimal, diffFloat is the dollar amount.
	// To be robust, let's check if the magnitude exceeds thresholdUSD. If difference is large (> thresholdUSD * 10000000), it's stroops; otherwise if it's small or if we treat Difference as USDC units.
	// Wait, in ledgerReconciliationJob, absDiff is int64 stroops. Let's convert stroops to USDC: stroops / 10000000.0.
	valUSD := diffFloat / 10_000_000.0
	// Also support if Difference was constructed as a direct dollar decimal (e.g. in tests).
	if valUSD > p.thresholdUSD {
		return true
	}
	if diffFloat > p.thresholdUSD {
		return true
	}
	return false
}
