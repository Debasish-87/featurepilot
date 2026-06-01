package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	featureDTO "github.com/Debasish-87/featurepilot/internal/feature/dto"
)

func (h *Handler) List(c *gin.Context) {

	features, err := h.service.List(
		c.Request.Context(),
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	response := make(
		[]featureDTO.FeatureResponse,
		0,
		len(features),
	)

	for _, feature := range features {

		response = append(
			response,
			featureDTO.FeatureResponse{
				ID:            feature.ID.String(),
				EnvironmentID: feature.EnvironmentID.String(),
				Key:           feature.Key,
				Name:          feature.Name,
				Description:   feature.Description,
				Enabled:       feature.Enabled,
			},
		)
	}

	c.JSON(
		http.StatusOK,
		response,
	)
}
