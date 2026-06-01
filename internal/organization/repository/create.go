package repository

import (
	"context"

	"github.com/Debasish-87/featurepilot/internal/organization/domain"
)

func (r *PostgresRepository) Create(
	ctx context.Context,
	organization *domain.Organization,
) error {

	query := `
		INSERT INTO organizations (
			id,
			name,
			created_at,
			updated_at
		)
		VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		organization.ID,
		organization.Name,
		organization.CreatedAt,
		organization.UpdatedAt,
	)

	return err
}
