package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/Debasish-87/featurepilot/internal/release/domain"
)

const ErrorRateThreshold = 10.0

func (s *Service) RollbackIfNeeded(
	ctx context.Context,
	releaseID uuid.UUID,
	errorRate float64,
) error {

	if errorRate < ErrorRateThreshold {
		return nil
	}

	release, err := s.repo.GetByID(
		ctx,
		releaseID,
	)
	if err != nil {
		return err
	}

	previousRelease, err := s.repo.GetPreviousActive(
		ctx,
		release.ProjectID,
	)
	if err != nil {
		return err
	}

	if err := s.repo.UpdateStatus(
		ctx,
		releaseID,
		domain.StatusFailed,
	); err != nil {
		return err
	}

	return s.repo.UpdateStatus(
		ctx,
		previousRelease.ID,
		domain.StatusActive,
	)
}
