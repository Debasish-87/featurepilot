package service

import (
	"context"

	"github.com/Debasish-87/featurepilot/internal/release/domain"
)

func (s *Service) List(
	ctx context.Context,
) ([]domain.Release, error) {

	return s.repo.List(ctx)
}