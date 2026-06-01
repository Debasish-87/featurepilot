package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	environmentDTO "github.com/Debasish-87/featurepilot/internal/environment/dto"
)

func (h *Handler) List(c *gin.Context) {

	environments, err := h.service.List(
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
		[]environmentDTO.EnvironmentResponse,
		0,
		len(environments),
	)

	for _, environment := range environments {

		response = append(
			response,
			environmentDTO.EnvironmentResponse{
				ID:        environment.ID.String(),
				ProjectID: environment.ProjectID.String(),
				Name:      environment.Name,
			},
		)
	}

	c.JSON(
		http.StatusOK,
		response,
	)
}
