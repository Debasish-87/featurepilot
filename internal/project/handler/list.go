package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	projectDTO "github.com/Debasish-87/featurepilot/internal/project/dto"
)

func (h *Handler) List(c *gin.Context) {

	projects, err := h.service.List(
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
		[]projectDTO.ListResponse,
		0,
		len(projects),
	)

	for _, project := range projects {

		response = append(
			response,
			projectDTO.ListResponse{
				ID:             project.ID.String(),
				OrganizationID: project.OrganizationID.String(),
				Name:           project.Name,
			},
		)
	}

	c.JSON(
		http.StatusOK,
		response,
	)
}
