package handler

import releaseService "github.com/Debasish-87/featurepilot/internal/release/service"

type Handler struct {
	service *releaseService.Service
}

func New(
	service *releaseService.Service,
) *Handler {
	return &Handler{
		service: service,
	}
}