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

	r.GET("/test", h.TestHandler)
	r.GET("/keyword", h.GetKeywordList)

	return r
}

func (r *Router) ServerStart(port string) error {
	return r.engine.Run(port)
}
