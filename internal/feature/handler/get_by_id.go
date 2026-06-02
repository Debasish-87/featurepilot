package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	featureDTO "github.com/Debasish-87/featurepilot/internal/feature/dto"
)

func (h *Handler) GetByID(c *gin.Context) {

	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid id",
			},
		)
		return
	}

	feature, err := h.service.GetByID(
		c.Request.Context(),
		id,
	)

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		featureDTO.FeatureResponse{
			ID:                feature.ID.String(),
			EnvironmentID:     feature.EnvironmentID.String(),
			Key:               feature.Key,
			Name:              feature.Name,
			Description:       feature.Description,
			Enabled:           feature.Enabled,
			RolloutPercentage: feature.RolloutPercentage,
		},
	)
}
