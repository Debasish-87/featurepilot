package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	releaseMetricDTO "github.com/Debasish-87/featurepilot/internal/release_metric/dto"
)

func (h *Handler) Create(
	c *gin.Context,
) {

	var req releaseMetricDTO.CreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	releaseID, err := uuid.Parse(
		req.ReleaseID,
	)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid release id",
			},
		)

		return
	}

	err = h.service.Create(
		c.Request.Context(),
		releaseID,
		req.ErrorRate,
		req.LatencyMS,
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
		gin.H{
			"message": "release metric created",
		},
	)
}