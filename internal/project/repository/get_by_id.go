package repository

import (
	"context"

	"github.com/google/uuid"

	"github.com/Debasish-87/featurepilot/internal/project/domain"
)

func (r *PostgresRepository) GetByID(
	ctx context.Context,
	id uuid.UUID,
) (*domain.Project, error) {

	query := `
		SELECT
			id,
			organization_id,
			name,
			created_at,
			updated_at
		FROM projects
		WHERE id = $1
	`

	var project domain.Project

	err := r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&project.ID,
		&project.OrganizationID,
		&project.Name,
		&project.CreatedAt,
		&project.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &project, nil
}
