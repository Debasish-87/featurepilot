package repository

import (
	"context"

	"github.com/Debasish-87/featurepilot/internal/audit/domain"
)

func (r *PostgresRepository) Create(
	ctx context.Context,
	log *domain.AuditLog,
) error {

	query := `
		INSERT INTO audit_logs (
			id,
			action,
			entity_type,
			entity_id,
			metadata,
			created_at
		)
		VALUES ($1,$2,$3,$4,$5,$6)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		log.ID,
		log.Action,
		log.EntityType,
		log.EntityID,
		log.Metadata,
		log.CreatedAt,
	)

	return err
}
