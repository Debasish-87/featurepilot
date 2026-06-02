package repository

import (
	"context"

	"github.com/Debasish-87/featurepilot/internal/release/domain"
)

func (r *PostgresRepository) Create(
	ctx context.Context,
	release *domain.Release,
) error {

	query := `
		INSERT INTO releases (
			id,
			project_id,
			version,
			status,
			created_at
		)
		VALUES ($1,$2,$3,$4,$5)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		release.ID,
		release.ProjectID,
		release.Version,
		release.Status,
		release.CreatedAt,
	)

	return err
}