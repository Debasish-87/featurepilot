package service

import (
	"context"

	eventDomain "github.com/Debasish-87/featurepilot/internal/evaluation_event/domain"
)

type EventRepository interface {
	Create(
		ctx context.Context,
		event *eventDomain.EvaluationEvent,
	) error
}
