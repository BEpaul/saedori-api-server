package handler

import (
	"net/http"
	"sync"

	"github.com/bestkkii/saedori-api-server/internal/model"
	"github.com/bestkkii/saedori-api-server/internal/service"
	"github.com/bestkkii/saedori-api-server/pkg"
	"github.com/gin-gonic/gin"
)

var (
	handlerInit     sync.Once
	handlerInstance *Handler
)

type Handler struct {
	dashboardService *service.Dashboard
}

func NewHandler(dashboardService *service.Dashboard) *Handler {
	handlerInit.Do(func() {
		handlerInstance = &Handler{
			dashboardService: dashboardService,
		}
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

func (h *Handler) GetKeywordList(c *gin.Context) {
	keywords, err := h.dashboardService.GetKeywordList()
	if err != nil {
		h.failedResponse(c, pkg.NewApiResponse("FAILED", err))
		return
	}

	h.okResponse(c, model.GetKeywordListResponse{
		ApiResponse: pkg.NewApiResponse("SUCCESS", nil),
		Keywords:    keywords,
	})
}
