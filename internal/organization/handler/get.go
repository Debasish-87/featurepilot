package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/Debasish-87/featurepilot/internal/organization/dto"
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

	org, err := h.service.GetByID(
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
		dto.OrganizationResponse{
			ID:   org.ID.String(),
			Name: org.Name,
		},
	)
}
