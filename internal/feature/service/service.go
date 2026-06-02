package service

import (
	"context"

	"github.com/google/uuid"
)

type AuditService interface {
	Log(
		ctx context.Context,
		action string,
		entityType string,
		entityID uuid.UUID,
		metadata string,
	)
}

type Service struct {
	repo      Repository
	envRepo   EnvironmentRepository
	evalCache EvaluationCache
	audit     AuditService
}

func New(
	repo Repository,
	envRepo EnvironmentRepository,
	evalCache EvaluationCache,
	audit AuditService,
) *Service {

	return &Service{
		repo:      repo,
		envRepo:   envRepo,
		evalCache: evalCache,
		audit:     audit,
	}
}
