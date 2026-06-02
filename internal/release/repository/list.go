package repository

import (
	"context"

	"github.com/Debasish-87/featurepilot/internal/release/domain"
)

func (r *PostgresRepository) List(
	ctx context.Context,
) ([]domain.Release, error) {

	query := `
		SELECT
			id,
			project_id,
			version,
			status,
			created_at
		FROM releases
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(
		ctx,
		query,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var releases []domain.Release

	for rows.Next() {

		var release domain.Release

		err := rows.Scan(
			&release.ID,
			&release.ProjectID,
			&release.Version,
			&release.Status,
			&release.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		releases = append(
			releases,
			release,
		)
	}

	return releases, nil
}