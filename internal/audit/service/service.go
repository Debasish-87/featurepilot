package service

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"

	"github.com/Debasish-87/featurepilot/internal/audit/domain"
)

type Repository interface {
	Create(
		ctx context.Context,
		logEntry *domain.AuditLog,
	) error
}

type Service struct {
	repo Repository
}

func New(
	repo Repository,
) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Log(
	ctx context.Context,
	action string,
	entityType string,
	entityID uuid.UUID,
	metadata string,
) {

	log.Println(
		"AUDIT EVENT:",
		action,
		entityType,
		entityID,
	)

	err := s.repo.Create(
		ctx,
		&domain.AuditLog{
			ID:         uuid.New(),
			Action:     action,
			EntityType: entityType,
			EntityID:   entityID,
			Metadata:   metadata,
			CreatedAt:  time.Now().UTC(),
		},
	)

	if err != nil {
		log.Println("AUDIT INSERT FAILED:", err)
		return
	}

	log.Println("AUDIT INSERT SUCCESS")
}
