package handler

import (
	evaluationService "github.com/Debasish-87/featurepilot/internal/evaluation/service"
)

type Handler struct {
	service *evaluationService.Service
}

func New(
	service *evaluationService.Service,
) *Handler {
	return &Handler{
		service: service,
	}
}
