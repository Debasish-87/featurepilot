package repository

import (
	"context"

	"github.com/google/uuid"

	"github.com/Debasish-87/featurepilot/internal/release/domain"
)

func (r *PostgresRepository) GetActive(
	ctx context.Context,
	projectID uuid.UUID,
) (*domain.Release, error) {

	query := `
		SELECT
			id,
			project_id,
			version,
			status,
			created_at
		FROM releases
		WHERE project_id = $1
		AND status = 'ACTIVE'
		LIMIT 1
	`

	var release domain.Release

	err := r.db.QueryRow(
		ctx,
		query,
		projectID,
	).Scan(
		&release.ID,
		&release.ProjectID,
		&release.Version,
		&release.Status,
		&release.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &release, nil
}