package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/Debasish-87/featurepilot/internal/feature/domain"
)

type Repository interface {
	Create(
		ctx context.Context,
		feature *domain.Feature,
	) error

	GetByID(
		ctx context.Context,
		id uuid.UUID,
	) (*domain.Feature, error)

	List(
		ctx context.Context,
	) ([]domain.Feature, error)

	Delete(
		ctx context.Context,
		id uuid.UUID,
	) error

	Enable(
		ctx context.Context,
		id uuid.UUID,
	) error

	Disable(
		ctx context.Context,
		id uuid.UUID,
	) error
}
