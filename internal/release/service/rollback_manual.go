package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/Debasish-87/featurepilot/internal/release/domain"
)

func (s *Service) Rollback(
	ctx context.Context,
	releaseID uuid.UUID,
) error {

	// current release fetch
	release, err := s.repo.GetByID(
		ctx,
		releaseID,
	)
	if err != nil {
		return err
	}

	// previous active release fetch
	previousRelease, err := s.repo.GetPreviousActive(
		ctx,
		release.ProjectID,
	)
	if err != nil {
		return err
	}

	if previousRelease == nil {
		return fmt.Errorf(
			"no previous active release found",
		)
	}

	// mark current release as rolled back
	if err := s.repo.UpdateStatus(
		ctx,
		releaseID,
		domain.StatusRolledBack,
	); err != nil {
		return err
	}

	// reactivate previous stable release
	return s.repo.UpdateStatus(
		ctx,
		previousRelease.ID,
		domain.StatusActive,
	)
}
