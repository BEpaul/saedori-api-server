package scheduler

import (
	"log"

	"github.com/bestkkii/saedori-api-server/internal/service"
)

type RealtimeSearchScheduler struct {
	dashboardService *service.Dashboard
}

func RealtimeSearchService(dashboardService *service.Dashboard) *RealtimeSearchScheduler {
	return &RealtimeSearchScheduler{
		dashboardService: dashboardService,
	}
}

// 실시간 검색어로부터 오늘의 단어 3개 뽑기
func(r *RealtimeSearchScheduler) GetKeywordsFromRealtimeSearchData() ([]string, error) {
	realtimeSearchData, err := r.dashboardService.GetRealtimeSearchList()

	if err != nil {
		log.Println("실시간 검색어 데이터 조회 실패:", err)
		return nil, err
	}

	return realtimeSearchData[:3], nil
}
