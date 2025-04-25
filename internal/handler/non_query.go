package handler

import (
	"github.com/bestkkii/saedori-api-server/internal/model"
	"github.com/bestkkii/saedori-api-server/pkg"
	"github.com/gin-gonic/gin"
)

/**
* 쿼리 스트링 없는 API 리스트
**/

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

// Music 순위 조회
func (h *Handler) GetMusicList(c *gin.Context) {
	musics, err := h.dashboardService.GetMusicList()

	if err != nil {
		h.failedResponse(c, pkg.NewApiResponse("FAILED"))
		return
	}

	h.okResponse(c, model.MusicResponse{
		ApiResponse: pkg.NewApiResponse("SUCCESS"),
		Musics:      musics,
	})
}

func (h *Handler) GetRealtimeSearchDetail(c *gin.Context) {
	realtimeSearchDetail, err := h.dashboardService.GetRealtimeSearchDetailList()

	if err != nil {
		h.failedResponse(c, pkg.NewApiResponse("FAILED"))
		return
	}

	h.okResponse(c, model.RealtimeSearchDetailResponse{
		ApiResponse:                 pkg.NewApiResponse("SUCCESS"),
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
