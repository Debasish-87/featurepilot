package repository

import (
	"context"

	"github.com/google/uuid"

	"github.com/Debasish-87/featurepilot/internal/organization/domain"
)

func (r *PostgresRepository) GetByID(
	ctx context.Context,
	id uuid.UUID,
) (*domain.Organization, error) {

	query := `
		SELECT
			id,
			name,
			created_at,
			updated_at
		FROM organizations
		WHERE id = $1
	`

	var org domain.Organization

	err := r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&org.ID,
		&org.Name,
		&org.CreatedAt,
		&org.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &org, nil
}
