package repository

import (
	"context"

	"github.com/Debasish-87/featurepilot/internal/release_metric/domain"
)

func (r *PostgresRepository) Create(
	ctx context.Context,
	metric *domain.ReleaseMetric,
) error {

	query := `
		INSERT INTO release_metrics (
			id,
			release_id,
			error_rate,
			latency_ms,
			created_at
		)
		VALUES ($1,$2,$3,$4,$5)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		metric.ID,
		metric.ReleaseID,
		metric.ErrorRate,
		metric.LatencyMS,
		metric.CreatedAt,
	)

	return err
}