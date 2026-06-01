package service

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/Debasish-87/featurepilot/internal/project/domain"
)

func (s *Service) Create(
	ctx context.Context,
	organizationID uuid.UUID,
	name string,
) (*domain.Project, error) {

	now := time.Now().UTC()

	project := &domain.Project{
		ID:             uuid.New(),
		OrganizationID: organizationID,
		Name:           name,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	if err := s.repo.Create(
		ctx,
		project,
	); err != nil {
		return nil, err
	}

	return project, nil
}
