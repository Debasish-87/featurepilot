package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) Activate(c *gin.Context) {

	id, err := uuid.Parse(
		c.Param("id"),
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid release id",
		})
		return
	}

	err = h.service.Activate(
		c.Request.Context(),
		id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "release activated",
	})
}
