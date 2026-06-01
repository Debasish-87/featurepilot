package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) Disable(c *gin.Context) {

	id, err := uuid.Parse(
		c.Param("id"),
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid id",
			},
		)
		return
	}

	if err := h.service.Disable(
		c.Request.Context(),
		id,
	); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.Status(http.StatusNoContent)
}
