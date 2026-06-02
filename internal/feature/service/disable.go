package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *Service) Disable(
	ctx context.Context,
	id uuid.UUID,
) error {

	feature, err := s.repo.GetByID(
		ctx,
		id,
	)
	if err != nil {
		return err
	}

	if err := s.repo.Disable(
		ctx,
		id,
	); err != nil {
		return err
	}

	environment, err := s.envRepo.GetByID(
		ctx,
		feature.EnvironmentID,
	)
	if err != nil {
		return err
	}

	// Audit Event
	if s.audit != nil {
		s.audit.Log(
			ctx,
			"FEATURE_DISABLED",
			"feature",
			id,
			fmt.Sprintf(
				`{"key":"%s","environment":"%s"}`,
				feature.Key,
				environment.Name,
			),
		)
	}

	// TODO:
	// rollout cache uses:
	// feature:<env>:<feature>:<user>
	// current invalidation won't clear all user caches

	cacheKey := fmt.Sprintf(
		"feature:%s:%s",
		environment.Name,
		feature.Key,
	)

	_ = s.evalCache.Delete(
		ctx,
		cacheKey,
	)

	return nil
}
