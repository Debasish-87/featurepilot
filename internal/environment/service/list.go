package service

import (
	"context"

	environmentDomain "github.com/Debasish-87/featurepilot/internal/environment/domain"
)

func (s *Service) List(
	ctx context.Context,
) ([]environmentDomain.Environment, error) {

	return s.repo.List(ctx)
}
