package repository

import (
	"context"

	"github.com/Debasish-87/featurepilot/internal/feature/domain"
)

func (r *PostgresRepository) List(
	ctx context.Context,
) ([]domain.Feature, error) {

	query := `
		SELECT
			id,
			environment_id,
			key,
			name,
			description,
			enabled,
			rollout_percentage,
			created_at,
			updated_at
		FROM features
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

	var features []domain.Feature

	for rows.Next() {

		var feature domain.Feature

		if err := rows.Scan(
			&feature.ID,
			&feature.EnvironmentID,
			&feature.Key,
			&feature.Name,
			&feature.Description,
			&feature.Enabled,
			&feature.RolloutPercentage,
			&feature.CreatedAt,
			&feature.UpdatedAt,
		); err != nil {
			return nil, err
		}

		features = append(
			features,
			feature,
		)
	}

	return features, nil
}
