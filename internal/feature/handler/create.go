package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	featureDTO "github.com/Debasish-87/featurepilot/internal/feature/dto"
)

func (h *Handler) Create(c *gin.Context) {

	var request featureDTO.CreateRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	environmentID, err := uuid.Parse(
		request.EnvironmentID,
	)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid environment_id",
			},
		)
		return
	}

	feature, err := h.service.Create(
		c.Request.Context(),
		environmentID,
		request.Key,
		request.Name,
		request.Description,
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

	c.JSON(
		http.StatusCreated,
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
