package service

import (
	"context"
	"time"

	"github.com/Debasish-87/featurepilot/internal/organization/domain"
	"github.com/google/uuid"
)

func (s *Service) Create(
	ctx context.Context,
	name string,
) (*domain.Organization, error) {

	now := time.Now().UTC()

	organization := &domain.Organization{
		ID:        uuid.New(),
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := s.repo.Create(
		ctx,
		organization,
	); err != nil {
		return nil, err
	}

	return organization, nil
}
