package repository

import (
	"context"

	"github.com/google/uuid"

	"github.com/Debasish-87/featurepilot/internal/release_metric/domain"
)

func (r *PostgresRepository) GetLatestByRelease(
	ctx context.Context,
	releaseID uuid.UUID,
) (*domain.ReleaseMetric, error) {

	query := `
		SELECT
			id,
			release_id,
			error_rate,
			latency_ms,
			created_at
		FROM release_metrics
		WHERE release_id = $1
		ORDER BY created_at DESC
		LIMIT 1
	`

	var metric domain.ReleaseMetric

	err := r.db.QueryRow(
		ctx,
		query,
		releaseID,
	).Scan(
		&metric.ID,
		&metric.ReleaseID,
		&metric.ErrorRate,
		&metric.LatencyMS,
		&metric.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &metric, nil
}