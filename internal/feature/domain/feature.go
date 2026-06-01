package domain

import (
	"time"

	"github.com/google/uuid"
)

type Feature struct {
	ID uuid.UUID

	EnvironmentID uuid.UUID

	Key string

	Name string

	Description string

	Enabled bool

	CreatedAt time.Time
	UpdatedAt time.Time
}
