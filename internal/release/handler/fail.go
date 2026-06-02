package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) Fail(c *gin.Context) {

	releaseID, err := uuid.Parse(
		c.Param("id"),
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid release id",
		})
		return
	}

	err = h.service.Fail(
		c.Request.Context(),
		releaseID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "release failed",
	})
}
