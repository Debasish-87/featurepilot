package repository

import (
	"context"

	"github.com/google/uuid"

	"github.com/Debasish-87/featurepilot/internal/feature/domain"
)

func (r *PostgresRepository) GetByID(
	ctx context.Context,
	id uuid.UUID,
) (*domain.Feature, error) {

	query := `
		SELECT
			id,
			environment_id,
			key,
			name,
			description,
			enabled,
			rollout_percentage,
			created_at,
			updated_at
		FROM features
		WHERE id = $1
	`

	var feature domain.Feature

	err := r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&feature.ID,
		&feature.EnvironmentID,
		&feature.Key,
		&feature.Name,
		&feature.Description,
		&feature.Enabled,
		&feature.RolloutPercentage,
		&feature.CreatedAt,
		&feature.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &feature, nil
}
