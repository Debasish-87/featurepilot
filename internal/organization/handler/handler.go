package handler

import "github.com/Debasish-87/featurepilot/internal/organization/service"

type Handler struct {
	service *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}
