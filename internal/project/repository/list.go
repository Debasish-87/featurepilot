package repository

import (
	"context"

	"github.com/Debasish-87/featurepilot/internal/project/domain"
)

func (r *PostgresRepository) List(
	ctx context.Context,
) ([]domain.Project, error) {

	query := `
		SELECT
			id,
			organization_id,
			name,
			created_at,
			updated_at
		FROM projects
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

	var projects []domain.Project

	for rows.Next() {

		var project domain.Project

		if err := rows.Scan(
			&project.ID,
			&project.OrganizationID,
			&project.Name,
			&project.CreatedAt,
			&project.UpdatedAt,
		); err != nil {
			return nil, err
		}

		projects = append(
			projects,
			project,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}
