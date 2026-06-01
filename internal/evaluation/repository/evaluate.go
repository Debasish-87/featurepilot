package repository

import (
	"context"
)

func (r *PostgresRepository) Evaluate(
	ctx context.Context,
	environment string,
	featureKey string,
) (bool, error) {

	query := `
		SELECT f.enabled
		FROM features f
		JOIN environments e
		ON e.id = f.environment_id
		WHERE e.name = $1
		AND f.key = $2
	`

	var enabled bool

	err := r.db.QueryRow(
		ctx,
		query,
		environment,
		featureKey,
	).Scan(
		&enabled,
	)

	return enabled, err
}
