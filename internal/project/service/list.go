package service

import (
	"context"

	"github.com/Debasish-87/featurepilot/internal/project/domain"
)

func (s *Service) List(
	ctx context.Context,
) ([]domain.Project, error) {

	return s.repo.List(ctx)
}
