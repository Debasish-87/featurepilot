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
