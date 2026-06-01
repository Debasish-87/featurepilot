package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/Debasish-87/featurepilot/internal/environment/domain"
)

type Repository interface {
	Create(
		ctx context.Context,
		environment *domain.Environment,
	) error

	GetByID(
		ctx context.Context,
		id uuid.UUID,
	) (*domain.Environment, error)

	List(
		ctx context.Context,
	) ([]domain.Environment, error)

	Delete(
		ctx context.Context,
		id uuid.UUID,
	) error
}
