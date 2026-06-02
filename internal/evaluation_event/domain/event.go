package domain

import (
	"time"

	"github.com/google/uuid"
)

type EvaluationEvent struct {
	ID uuid.UUID

	Environment string

	FeatureKey string

	UserID string

	Enabled bool

	CreatedAt time.Time
}