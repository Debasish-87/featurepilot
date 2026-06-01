package service

import (
	"context"

	"github.com/google/uuid"

	environmentDomain "github.com/Debasish-87/featurepilot/internal/environment/domain"
)

type EnvironmentRepository interface {
	GetByID(
		ctx context.Context,
		id uuid.UUID,
	) (*environmentDomain.Environment, error)
}
