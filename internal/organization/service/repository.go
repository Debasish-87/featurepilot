package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/Debasish-87/featurepilot/internal/organization/domain"
)

type Repository interface {
	Create(
		ctx context.Context,
		organization *domain.Organization,
	) error

	GetByID(
		ctx context.Context,
		id uuid.UUID,
	) (*domain.Organization, error)

	List(
		ctx context.Context,
	) ([]domain.Organization, error)

	Delete(
		ctx context.Context,
		id uuid.UUID,
	) error
}
