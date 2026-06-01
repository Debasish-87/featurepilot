package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/Debasish-87/featurepilot/internal/project/domain"
)

type Repository interface {
	Create(
		ctx context.Context,
		project *domain.Project,
	) error

	GetByID(
		ctx context.Context,
		id uuid.UUID,
	) (*domain.Project, error)

	List(
		ctx context.Context,
	) ([]domain.Project, error)

	Delete(
		ctx context.Context,
		id uuid.UUID,
	) error
}
