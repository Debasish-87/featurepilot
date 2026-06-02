package repository

import (
	"context"

	"github.com/google/uuid"

	releaseDomain "github.com/Debasish-87/featurepilot/internal/release/domain"
)

func (r *PostgresRepository) UpdateStatus(
	ctx context.Context,
	releaseID uuid.UUID,
	status releaseDomain.Status,
) error {

	query := `
		UPDATE releases
		SET status = $2
		WHERE id = $1
	`

	_, err := r.db.Exec(
		ctx,
		query,
		releaseID,
		status,
	)

	return err
}
