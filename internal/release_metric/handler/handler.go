package handler

import releaseMetricService "github.com/Debasish-87/featurepilot/internal/release_metric/service"

type Handler struct {
	service *releaseMetricService.Service
}

func New(
	service *releaseMetricService.Service,
) *Handler {

	return &Handler{
		service: service,
	}
}