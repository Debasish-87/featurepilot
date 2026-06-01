package handler

import (
	"net/http"

	"github.com/Debasish-87/featurepilot/internal/organization/dto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(c *gin.Context) {
	var req dto.CreateOrganizationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	org, err := h.service.Create(
		c.Request.Context(),
		req.Name,
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
		dto.CreateOrganizationResponse{
			ID:   org.ID.String(),
			Name: org.Name,
		},
	)

}
