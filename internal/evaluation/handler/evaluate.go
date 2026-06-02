package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Debasish-87/featurepilot/internal/evaluation/dto"
)

func (h *Handler) Evaluate(c *gin.Context) {

	var request dto.EvaluateRequest

	if err := c.ShouldBindJSON(&request); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	enabled, err := h.service.Evaluate(
		c.Request.Context(),
		request.Environment,
		request.FeatureKey,
		request.UserID,
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
		dto.EvaluateResponse{
			Enabled: enabled,
		},
	)
}