package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/Debasish-87/featurepilot/internal/feature/domain"
)

func (s *Service) GetByID(
	ctx context.Context,
	id uuid.UUID,
) (*domain.Feature, error) {

	return s.repo.GetByID(
		ctx,
		id,
	)
}
