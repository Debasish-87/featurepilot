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

	// Percentage of users eligible for this feature.
	// Range: 0-100
	RolloutPercentage int

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (f *Feature) IsRolloutValid() bool {
	return f.RolloutPercentage >= 0 &&
		f.RolloutPercentage <= 100
}
