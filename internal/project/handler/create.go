package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/Debasish-87/featurepilot/internal/project/dto"
)

func (h *Handler) Create(c *gin.Context) {

	var request dto.CreateProjectRequest

	if err := c.ShouldBindJSON(
		&request,
	); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	organizationID, err := uuid.Parse(
		request.OrganizationID,
	)
	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid organization id",
			},
		)
		return
	}

	project, err := h.service.Create(
		c.Request.Context(),
		organizationID,
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
		dto.ProjectResponse{
			ID:             project.ID.String(),
			OrganizationID: project.OrganizationID.String(),
			Name:           project.Name,
		},
	)
}
