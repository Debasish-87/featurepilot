package service

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/Debasish-87/featurepilot/internal/feature/domain"
)

func (s *Service) Create(
	ctx context.Context,
	environmentID uuid.UUID,
	key string,
	name string,
	description string,
	rolloutPercentage int,
) (*domain.Feature, error) {

	if rolloutPercentage < 0 ||
		rolloutPercentage > 100 {

		return nil, errors.New(
			"rollout percentage must be between 0 and 100",
		)
	}

	// default rollout
	if rolloutPercentage == 0 {
		rolloutPercentage = 100
	}

	now := time.Now().UTC()

	feature := &domain.Feature{
		ID:                uuid.New(),
		EnvironmentID:     environmentID,
		Key:               key,
		Name:              name,
		Description:       description,
		Enabled:           false,
		RolloutPercentage: rolloutPercentage,
		CreatedAt:         now,
		UpdatedAt:         now,
	}

	if err := s.repo.Create(
		ctx,
		feature,
	); err != nil {
		return nil, err
	}

	// Audit Event
	if s.audit != nil {
		s.audit.Log(
			ctx,
			"FEATURE_CREATED",
			"feature",
			feature.ID,
			key,
		)
	}

	return feature, nil
}
