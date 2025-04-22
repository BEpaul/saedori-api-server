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

	r.GET("/", func(c *gin.Context) {
		handler.NewHandler(c, service.Dashboard)
	})

	return r
}

func (r *Router) ServerStart(port string) error {
	return r.engine.Run(port)
}
