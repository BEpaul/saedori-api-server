package service

import (
	"github.com/bestkkii/saedori-api-server/internal/model"
	"github.com/bestkkii/saedori-api-server/internal/repository"
)

type Dashboard struct {
	dashboardRepository *repository.DashboardRepository
}

func newDashboardService(dashboardRepository *repository.DashboardRepository) *Dashboard {
	return &Dashboard{
		dashboardRepository: dashboardRepository,
	}
}

func (d *Dashboard) GetKeywordsList() ([]*model.Keywords, error) {
	keywordsList, err := d.dashboardRepository.KeywordRepository.GetKeywords()
	if err != nil {
		return nil, err
	}
	return keywordsList, nil
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
	newsList, err := d.dashboardRepository.NewsRepository.GetNewsDetails()
	if err != nil {
		return nil, err
	}
	return newsList, nil
}
