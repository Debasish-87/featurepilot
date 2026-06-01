package repository

import (
	"context"

	"github.com/google/uuid"
)

func (r *PostgresRepository) Enable(
	ctx context.Context,
	id uuid.UUID,
) error {

	query := `
		UPDATE features
		SET enabled = true,
		    updated_at = NOW()
		WHERE id = $1
	`

	_, err := r.db.Exec(
		ctx,
		query,
		id,
	)

	return err
}
