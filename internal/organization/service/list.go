package service

import (
	"context"

	"github.com/Debasish-87/featurepilot/internal/organization/domain"
)

func (s *Service) List(
	ctx context.Context,
) ([]domain.Organization, error) {

	return s.repo.List(ctx)
}
