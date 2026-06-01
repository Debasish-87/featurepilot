package handler

import featureService "github.com/Debasish-87/featurepilot/internal/feature/service"

type Handler struct {
	service *featureService.Service
}

func New(
	service *featureService.Service,
) *Handler {
	return &Handler{
		service: service,
	}
}
