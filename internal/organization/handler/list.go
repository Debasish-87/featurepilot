package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Debasish-87/featurepilot/internal/organization/dto"
)

func (h *Handler) List(c *gin.Context) {

	orgs, err := h.service.List(
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
		[]dto.OrganizationResponse,
		0,
		len(orgs),
	)

	for _, org := range orgs {

		response = append(
			response,
			dto.OrganizationResponse{
				ID:   org.ID.String(),
				Name: org.Name,
			},
		)
	}

	c.JSON(
		http.StatusOK,
		response,
	)
}
