package repository

import (
	"time"

	"github.com/bestkkii/saedori-api-server/internal/model"
)

type DashboardRepository struct {
	// Access DB (mongoDB)

	// for test
	keywordList []*model.Keyword
}

func newDashboardRepository() *DashboardRepository {
	return &DashboardRepository{
		keywordList: []*model.Keyword{
			{
				Keyword:   "test",
				CreatedAt: time.Now(),
			},
			{
				Keyword:   "test2",
				CreatedAt: time.Now(),
			},
		},
	}
}

// GetKeywordList 키워드 목록 조회
func (d *DashboardRepository) GetKeywordList() ([]*model.Keyword, error) {
	return d.keywordList, nil
}
