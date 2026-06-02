package service

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/Debasish-87/featurepilot/internal/release/domain"
)

func (s *Service) Create(
	ctx context.Context,
	projectID uuid.UUID,
	version string,
) (*domain.Release, error) {

	release := &domain.Release{
		ID:        uuid.New(),
		ProjectID: projectID,
		Version:   version,
		Status:    domain.StatusRollingOut,
		CreatedAt: time.Now().UTC(),
	}

	if err := s.repo.Create(
		ctx,
		release,
	); err != nil {
		return nil, err
	}

	return release, nil
}
