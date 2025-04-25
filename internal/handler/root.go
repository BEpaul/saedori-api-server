package handler

import (
	"sync"

	"github.com/bestkkii/saedori-api-server/internal/service"
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
