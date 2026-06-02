package domain

import (
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusDraft      Status = "DRAFT"
	StatusActive     Status = "ACTIVE"
	StatusRollingOut Status = "ROLLING_OUT"
	StatusCompleted  Status = "COMPLETED"
	StatusFailed     Status = "FAILED"
	StatusRolledBack Status = "ROLLED_BACK"
)

type Release struct {
	ID uuid.UUID

	ProjectID uuid.UUID

	Version string

	Status Status

	CreatedAt time.Time
}

func (r *Release) IsTerminal() bool {
	return r.Status == StatusCompleted ||
		r.Status == StatusFailed ||
		r.Status == StatusRolledBack
}

func (r *Release) CanStart() bool {
	return r.Status == StatusDraft
}

func (r *Release) CanRollback() bool {
	return r.Status == StatusActive ||
		r.Status == StatusRollingOut
}

func (r *Release) CanComplete() bool {
	return r.Status == StatusActive ||
		r.Status == StatusRollingOut
}