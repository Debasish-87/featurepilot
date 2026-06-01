package service

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/Debasish-87/featurepilot/internal/environment/domain"
)

func (s *Service) Create(
	ctx context.Context,
	projectID uuid.UUID,
	name string,
) (*domain.Environment, error) {

	now := time.Now().UTC()

	environment := &domain.Environment{
		ID:        uuid.New(),
		ProjectID: projectID,
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := s.repo.Create(
		ctx,
		environment,
	); err != nil {
		return nil, err
	}

	return environment, nil
}
