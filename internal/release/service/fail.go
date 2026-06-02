package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/Debasish-87/featurepilot/internal/release/domain"
)

func (s *Service) Fail(
	ctx context.Context,
	releaseID uuid.UUID,
) error {

	return s.repo.UpdateStatus(
		ctx,
		releaseID,
		domain.StatusFailed,
	)
}