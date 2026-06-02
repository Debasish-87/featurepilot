package domain

import (
	"time"

	"github.com/google/uuid"
)

type AuditLog struct {
	ID uuid.UUID

	Action string

	EntityType string

	EntityID uuid.UUID

	Metadata string

	CreatedAt time.Time
}
