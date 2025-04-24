package router

import (
	"github.com/bestkkii/saedori-api-server/internal/handler"
	"github.com/bestkkii/saedori-api-server/internal/service"
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
}

func NewRouter(service *service.Service) *Router {
	r := &Router{
		engine: gin.Default(),
	}

	h := handler.NewHandler(service.Dashboard)
	apiV1 := r.engine.Group("/api/v1")

	apiV1.GET("/keywords", h.GetKeywordList)
	apiV1.GET("/interest/detail", h.GetInterestDetail)

	return r
}

func (r *Router) ServerStart(port string) error {
	return r.engine.Run(port)
}
