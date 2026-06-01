package handler

import environmentService "github.com/Debasish-87/featurepilot/internal/environment/service"

type Handler struct {
	service *environmentService.Service
}

func New(
	service *environmentService.Service,
) *Handler {
	return &Handler{
		service: service,
	}
}
