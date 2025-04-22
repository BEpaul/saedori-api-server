package router

import (
	"github.com/gin-gonic/gin"
)

func (n *Router) GET(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.GET(path, handler...)
}
