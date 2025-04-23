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

func (d *Dashboard) GetKeywordList() ([]*model.Keyword, error) {
	keywordList, err := d.dashboardRepository.GetKeywords()
	if err != nil {
		return nil, err
	}
	return keywordList, nil
}
