package repository

import (
	"context"

	"github.com/Debasish-87/featurepilot/internal/project/domain"
)

func (r *PostgresRepository) Create(
	ctx context.Context,
	project *domain.Project,
) error {

	query := `
		INSERT INTO projects (
			id,
			organization_id,
			name,
			created_at,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		project.ID,
		project.OrganizationID,
		project.Name,
		project.CreatedAt,
		project.UpdatedAt,
	)

	return err
}
