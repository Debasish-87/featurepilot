package handler

import (
	projectService "github.com/Debasish-87/featurepilot/internal/project/service"
)

type Handler struct {
	service *projectService.Service
}

func New(
	service *projectService.Service,
) *Handler {
	return &Handler{
		service: service,
	}
}
