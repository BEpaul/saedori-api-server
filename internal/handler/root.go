package handler

import (
	"net/http"
	"sync"

	"github.com/bestkkii/saedori-api-server/internal/service"
	"github.com/gin-gonic/gin"
)

var (
	handlerInit     sync.Once
	handlerInstance *Handler
)

type Handler struct {
	dashboardService *service.Dashboard
}

func NewHandler(c *gin.Context, dashboardService *service.Dashboard) *Handler {
	handlerInit.Do(func() {
		handlerInstance = &Handler{
			dashboardService: dashboardService,
		}

		handlerInstance.TestHandler(c)
	})

	return handlerInstance
}

func (h *Handler) TestHandler(c *gin.Context) {
	result, err := h.dashboardService.GetKeywordList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
