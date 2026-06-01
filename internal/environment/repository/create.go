package repository

import (
	"context"

	"github.com/Debasish-87/featurepilot/internal/environment/domain"
)

func (r *PostgresRepository) Create(
	ctx context.Context,
	environment *domain.Environment,
) error {

	query := `
		INSERT INTO environments (
			id,
			project_id,
			name,
			created_at,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		environment.ID,
		environment.ProjectID,
		environment.Name,
		environment.CreatedAt,
		environment.UpdatedAt,
	)

	return err
}
