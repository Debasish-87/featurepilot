package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	environmentDTO "github.com/Debasish-87/featurepilot/internal/environment/dto"
)

func (h *Handler) Create(c *gin.Context) {

	var request environmentDTO.CreateRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	projectID, err := uuid.Parse(
		request.ProjectID,
	)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid project_id",
			},
		)
		return
	}

	environment, err := h.service.Create(
		c.Request.Context(),
		projectID,
		request.Name,
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
		environmentDTO.EnvironmentResponse{
			ID:        environment.ID.String(),
			ProjectID: environment.ProjectID.String(),
			Name:      environment.Name,
		},
	)
}
