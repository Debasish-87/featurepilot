package repository

import (
	"context"

	"github.com/google/uuid"

	"github.com/Debasish-87/featurepilot/internal/environment/domain"
)

func (r *PostgresRepository) GetByID(
	ctx context.Context,
	id uuid.UUID,
) (*domain.Environment, error) {

	query := `
		SELECT
			id,
			project_id,
			name,
			created_at,
			updated_at
		FROM environments
		WHERE id = $1
	`

	var environment domain.Environment

	err := r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&environment.ID,
		&environment.ProjectID,
		&environment.Name,
		&environment.CreatedAt,
		&environment.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &environment, nil
}
