package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) okResponse(c *gin.Context, result interface{}) {
	c.JSON(http.StatusOK, result)
}

func (h *Handler) failedResponse(c *gin.Context, result interface{}) {
	c.JSON(http.StatusBadRequest, result)
}
