package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) List(
	c *gin.Context,
) {

	releases, err := h.service.List(
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

	c.JSON(
		http.StatusOK,
		releases,
	)
}