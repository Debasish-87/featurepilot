package repository

import (
	"context"

	evaluationService "github.com/Debasish-87/featurepilot/internal/evaluation/service"
)

func (r *PostgresRepository) Evaluate(
	ctx context.Context,
	environment string,
	featureKey string,
) (*evaluationService.EvaluationResult, error) {

	query := `
		SELECT
			f.enabled,
			f.rollout_percentage
		FROM features f
		JOIN environments e
			ON e.id = f.environment_id
		WHERE e.name = $1
		AND f.key = $2
	`

	var result evaluationService.EvaluationResult

	err := r.db.QueryRow(
		ctx,
		query,
		environment,
		featureKey,
	).Scan(
		&result.Enabled,
		&result.RolloutPercentage,
	)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
