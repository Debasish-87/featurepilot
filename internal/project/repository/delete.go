package repository

import (
	"context"

	"github.com/google/uuid"
)

func (r *PostgresRepository) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {

	query := `
		DELETE
		FROM projects
		WHERE id = $1
	`

	_, err := r.db.Exec(
		ctx,
		query,
		id,
	)

	return err
}
