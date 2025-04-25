package handler

import (
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

// Top3 Keyword 목록 조회
func (h *Handler) GetKeywordsList(c *gin.Context) {
	keywords, err := h.dashboardService.GetKeywordsList()
	if err != nil {
		h.failedResponse(c, pkg.NewApiResponse("FAILED"))
		return
	}

	h.okResponse(c, model.GetKeywordsListResponse{
		ApiResponse: pkg.NewApiResponse("SUCCESS"),
		Keywords:    keywords,
	})
}

// 관심사 분야 쿼리 파라미터 매칭
func (h *Handler) GetInterestDetail(c *gin.Context) {
	category := c.DefaultQuery("category", "default_category")

	if category == "music" {
		h.GetMusicList(c)
		return
	} else if category == "realtime-search" {
		h.GetRealtimeSearchDetail(c)
		return
	} else if category == "news" {
		h.GetNewsDetails(c)
		return
	}
}

// Music 순위 조회
func (h *Handler) GetMusicList(c *gin.Context) {
	musics, err := h.dashboardService.GetMusicList()

	if err != nil {
		h.failedResponse(c, pkg.NewApiResponse("FAILED"))
		return
	}

	h.okResponse(c, model.MusicResponse{
		ApiResponse: pkg.NewApiResponse("SUCCESS"),
		Musics:   musics,
	})
}

func (h *Handler) GetRealtimeSearchDetail(c *gin.Context) {
	realtimeSearchDetail, err := h.dashboardService.GetRealtimeSearchDetailList()

	if err != nil {
		h.failedResponse(c, pkg.NewApiResponse("FAILED"))
		return
	}

	h.okResponse(c, model.RealtimeSearchDetailResponse{
		ApiResponse: pkg.NewApiResponse("SUCCESS"),
		RealtimeSearchDetailWrapper: realtimeSearchDetail.RealtimeSearchDetailWrapper,
	})
}

// News 상세 목록 조회
func (h *Handler) GetNewsDetails(c *gin.Context) {
	news, err := h.dashboardService.GetNewsDetails()
	if err != nil {
		h.failedResponse(c, pkg.NewApiResponse("FAILED"))
		return
	}

	h.okResponse(c, model.NewsResponse{
		ApiResponse: pkg.NewApiResponse("SUCCESS"),
		News:        news,
	})
}

// News 요약 조회
func (h *Handler) GetNewsSummary(c *gin.Context) {
	summaries, err := h.dashboardService.GetNewsSummary()
	if err != nil {
		h.failedResponse(c, pkg.NewApiResponse("FAILED"))
		return
	}

	h.okResponse(c, model.NewsSummaryResponse{
		ApiResponse: pkg.NewApiResponse("SUCCESS"),
		NewsSummary: summaries,
	})
}