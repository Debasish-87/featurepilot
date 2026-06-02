package service

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/Debasish-87/featurepilot/internal/release_metric/domain"
)

func (s *Service) Create(
	ctx context.Context,
	releaseID uuid.UUID,
	errorRate float64,
	latencyMS float64,
) error {

	metric := &domain.ReleaseMetric{
		ID:        uuid.New(),
		ReleaseID: releaseID,
		ErrorRate: errorRate,
		LatencyMS: latencyMS,
		CreatedAt: time.Now().UTC(),
	}

	if err := s.repo.Create(
		ctx,
		metric,
	); err != nil {
		return err
	}

	return s.releaseSvc.RollbackIfNeeded(
		ctx,
		releaseID,
		errorRate,
	)
}
