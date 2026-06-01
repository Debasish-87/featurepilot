package service

import (
	"context"

	"github.com/Debasish-87/featurepilot/internal/organization/domain"
	"github.com/google/uuid"
)

func (s *Service) GetByID(
	ctx context.Context,
	id uuid.UUID,
) (*domain.Organization, error) {

	return s.repo.GetByID(
		ctx,
		id,
	)
}
