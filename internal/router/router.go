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
	apiV1 := r.engine.Group("/api/v1") // GCP 접두어 추가 필요

	// endpoints
	apiV1.GET("/keywords", h.GetKeywordsList) // 오늘의 단어 조회 요청
	// apiV1.GET("/download", h.다운로드하는 함수) // 컨텐츠 다운로드 요청
	apiV1.GET("/interest/detail", h.GetInterestDetail) // 관심사 분야 상세 데이터 요청. ?category= 쿼리 파라미터 필요
	return r
}

func (r *Router) ServerStart(port string) error {
	return r.engine.Run(port)
}
