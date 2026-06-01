package service

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/Debasish-87/featurepilot/internal/feature/domain"
)

func (s *Service) Create(
	ctx context.Context,
	environmentID uuid.UUID,
	key string,
	name string,
	description string,
) (*domain.Feature, error) {

	now := time.Now().UTC()

	feature := &domain.Feature{
		ID:            uuid.New(),
		EnvironmentID: environmentID,
		Key:           key,
		Name:          name,
		Description:   description,
		Enabled:       false,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	if err := s.repo.Create(
		ctx,
		feature,
	); err != nil {
		return nil, err
	}

	return feature, nil
}
