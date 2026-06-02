package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/Debasish-87/featurepilot/internal/release/domain"
)

type Repository interface {
	Create(
		ctx context.Context,
		release *domain.Release,
	) error

	List(
		ctx context.Context,
	) ([]domain.Release, error)

	GetByID(
		ctx context.Context,
		id uuid.UUID,
	) (*domain.Release, error)

	GetActive(
		ctx context.Context,
		projectID uuid.UUID,
	) (*domain.Release, error)

	GetPreviousActive(
		ctx context.Context,
		projectID uuid.UUID,
	) (*domain.Release, error)

	UpdateStatus(
		ctx context.Context,
		releaseID uuid.UUID,
		status domain.Status,
	) error
}