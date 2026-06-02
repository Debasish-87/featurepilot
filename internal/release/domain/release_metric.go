package domain

import (
	"time"

	"github.com/google/uuid"
)

type ReleaseMetric struct {
	ID        uuid.UUID
	ReleaseID uuid.UUID

	ErrorRate float64
	LatencyMS float64

	CreatedAt time.Time
}