package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	releaseDTO "github.com/Debasish-87/featurepilot/internal/release/dto"
)

func (h *Handler) Create(c *gin.Context) {

	var req releaseDTO.CreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	projectID, err := uuid.Parse(
		req.ProjectID,
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "invalid project id"},
		)
		return
	}

	release, err := h.service.Create(
		c.Request.Context(),
		projectID,
		req.Version,
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		releaseDTO.ReleaseResponse{
			ID: release.ID.String(),
			ProjectID: release.ProjectID.String(),
			Version: release.Version,
			Status: string(release.Status),
		},
	)
}