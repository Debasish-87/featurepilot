package repository

import (
	"context"

	"github.com/Debasish-87/featurepilot/internal/evaluation_event/domain"
)

func (r *PostgresRepository) Create(
	ctx context.Context,
	event *domain.EvaluationEvent,
) error {

	query := `
		INSERT INTO evaluation_events (
			id,
			environment,
			feature_key,
			user_id,
			enabled,
			created_at
		)
		VALUES ($1,$2,$3,$4,$5,$6)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		event.ID,
		event.Environment,
		event.FeatureKey,
		event.UserID,
		event.Enabled,
		event.CreatedAt,
	)

	return err
}