package repository

import (
	"context"

	"github.com/Debasish-87/featurepilot/internal/environment/domain"
)

func (r *PostgresRepository) List(
	ctx context.Context,
) ([]domain.Environment, error) {

	query := `
		SELECT
			id,
			project_id,
			name,
			created_at,
			updated_at
		FROM environments
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

	var environments []domain.Environment

	for rows.Next() {

		var environment domain.Environment

		if err := rows.Scan(
			&environment.ID,
			&environment.ProjectID,
			&environment.Name,
			&environment.CreatedAt,
			&environment.UpdatedAt,
		); err != nil {
			return nil, err
		}

		environments = append(
			environments,
			environment,
		)
	}

	return environments, nil
}
