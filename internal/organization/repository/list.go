package repository

import (
	"context"

	"github.com/Debasish-87/featurepilot/internal/organization/domain"
)

func (r *PostgresRepository) List(
	ctx context.Context,
) ([]domain.Organization, error) {

	query := `
		SELECT
			id,
			name,
			created_at,
			updated_at
		FROM organizations
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

	var organizations []domain.Organization

	for rows.Next() {

		var org domain.Organization

		if err := rows.Scan(
			&org.ID,
			&org.Name,
			&org.CreatedAt,
			&org.UpdatedAt,
		); err != nil {
			return nil, err
		}

		organizations = append(
			organizations,
			org,
		)
	}

	return organizations, nil
}
