package repository

import (
	"context"

	"github.com/Debasish-87/featurepilot/internal/feature/domain"
)

func (r *PostgresRepository) Create(
	ctx context.Context,
	feature *domain.Feature,
) error {

	query := `
		INSERT INTO features (
			id,
			environment_id,
			key,
			name,
			description,
			enabled,
			created_at,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		feature.ID,
		feature.EnvironmentID,
		feature.Key,
		feature.Name,
		feature.Description,
		feature.Enabled,
		feature.CreatedAt,
		feature.UpdatedAt,
	)

	return err
}
