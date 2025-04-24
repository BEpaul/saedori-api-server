package service

import (
	"github.com/bestkkii/saedori-api-server/internal/model"
	"github.com/bestkkii/saedori-api-server/internal/repository"
	"regexp"
	"strings"
)

type Dashboard struct {
	dashboardRepository *repository.DashboardRepository
}

func newDashboardService(dashboardRepository *repository.DashboardRepository) *Dashboard {
	return &Dashboard{
		dashboardRepository: dashboardRepository,
	}
}

func (d *Dashboard) GetKeywordList() ([]*model.Keyword, error) {
	keywordList, err := d.dashboardRepository.GetKeywords()
	if err != nil {
		return nil, err
	}
	return keywordList, nil
}

func (d *Dashboard) GetMusicList() ([]*model.Music, error) {
	musicList, err := d.dashboardRepository.GetMusics()
	if err != nil {
		return nil, err
	}
	return musicList, nil
}

// 실시간 검색어 목록 조회 - summary
func (d *Dashboard) GetRealtimeSearchList() ([]string, error) {
	realtimeSearchList, err := d.dashboardRepository.GetRealtimeSearches()
	searchWordList := make([]string, 0, len(realtimeSearchList))
	if err != nil {
		return nil, err
	}

	for _, item := range realtimeSearchList {
		searchWordList = append(searchWordList, item.SearchWord)
	}
	return searchWordList, nil
}

// 실시간 검색어 목록 조회 - detail
func (d *Dashboard) GetRealtimeSearchDetailList() (*model.RealtimeSearchDetailResponse, error) {
	krList, usList, err := d.dashboardRepository.GetRealtimeSearchDetails()
	if err != nil {
		return nil, err
	}

	krWordList := make([]string, 0, len(krList))
	usWordList := make([]string, 0, len(usList))

	for _, item := range krList {
		krWordList = append(krWordList, item.SearchWord)
	}	

	for _, item := range usList {
		usWordList = append(usWordList, item.SearchWord)
	}

	realtimeSearchDetail := model.RealtimeSearchDetail{
		KrSearchWords: krWordList,
		UsSearchWords: usWordList,
	}

	realtimeSearchDetailWrapper := model.RealtimeSearchDetailWrapper{
		RealtimeSearchDetail: realtimeSearchDetail,
	}

	realtimeSearchDetailResponse := model.RealtimeSearchDetailResponse{
		RealtimeSearchDetailWrapper: realtimeSearchDetailWrapper,
	}

	return &realtimeSearchDetailResponse, nil
}

func (d *Dashboard) GetNewsDetails() ([]*model.News, error) {
	newsList, err := d.dashboardRepository.GetNewsDetails()
	if err != nil {
		return nil, err
	}
	return newsList, nil
}

func (d *Dashboard) GetNewsSummary() ([]model.NewsSummary, error) {
	news, err := d.dashboardRepository.GetNewsSummary()
	if err != nil {
		return nil, err
	}

	// 최대 5개의 뉴스만 처리
	summaryCount := 5
	if len(news.NewsItems) < summaryCount {
		summaryCount = len(news.NewsItems)
	}

	summaries := make([]model.NewsSummary, 0, summaryCount)
	for i := 0; i < summaryCount; i++ {
		item := news.NewsItems[i]
		
		// [ ]로 감싸진 문자열 제거
		re := regexp.MustCompile(`\[.*?\]`)
		title := re.ReplaceAllString(item.Title, "")
		
		// 앞쪽 공백 제거
		title = strings.TrimSpace(title)
		
		// 모든 특수문자 제거 (한글, 영문, 숫자, 공백만 남김)
		re = regexp.MustCompile(`[^가-힣a-zA-Z0-9\s]`)
		title = re.ReplaceAllString(title, "")

		summaries = append(summaries, model.NewsSummary{
			Company: item.Company,
			Title:   title,
		})
	}

	return summaries, nil
}