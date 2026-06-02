package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/Debasish-87/featurepilot/internal/release/domain"
)

func (s *Service) Activate(
	ctx context.Context,
	releaseID uuid.UUID,
) error {

	return s.repo.UpdateStatus(
		ctx,
		releaseID,
		domain.StatusActive,
	)
}