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
		FROM organizations
		WHERE id = $1
	`

	_, err := r.db.Exec(
		ctx,
		query,
		id,
	)

	return err
}
