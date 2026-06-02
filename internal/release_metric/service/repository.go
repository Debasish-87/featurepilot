package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/Debasish-87/featurepilot/internal/release_metric/domain"
)

type Repository interface {
	Create(
		ctx context.Context,
		metric *domain.ReleaseMetric,
	) error

	GetLatestByRelease(
		ctx context.Context,
		releaseID uuid.UUID,
	) (*domain.ReleaseMetric, error)
}