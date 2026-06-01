package service

import (
	"context"

	"github.com/Debasish-87/featurepilot/internal/feature/domain"
)

func (s *Service) List(
	ctx context.Context,
) ([]domain.Feature, error) {

	return s.repo.List(ctx)
}
