package scheduler

import (
	"log"
	"time"

	"github.com/bestkkii/saedori-api-server/internal/model"
	"github.com/bestkkii/saedori-api-server/internal/repository"
	"github.com/bestkkii/saedori-api-server/internal/service"
)

type KeywordScheduler struct {
	dashboardService        *service.Dashboard
	musicScheduler          *MusicScheduler
	newsScheduler           *NewsScheduler
	realtimeSearchScheduler *RealtimeSearchScheduler
	keywordRepository       *repository.KeywordRepository
}

func KeywordSchedulerService(dashboardService *service.Dashboard, keywordRepository *repository.KeywordRepository) *KeywordScheduler {
	return &KeywordScheduler{
		dashboardService:  dashboardService,
		keywordRepository: keywordRepository,
	}
}

func (k *KeywordScheduler) StartKeywordScheduler() {
	go func() {
		for {
			now := time.Now()
			next := time.Date(now.Year(), now.Month(), now.Day(), 7, 10, 0, 0, now.Location())
			if now.After(next) {
				next = next.Add(24 * time.Hour)
			}
			duration := next.Sub(now)
			time.AfterFunc(duration, func() {
				k.putKeywords()
			})
			time.Sleep(duration + 1*time.Second)
		}
	}()
}

func (k *KeywordScheduler) putKeywords() {
	k.musicScheduler = MusicService(k.dashboardService)
	k.newsScheduler = NewsService(k.dashboardService)
	k.realtimeSearchScheduler = RealtimeSearchService(k.dashboardService)

	k.processKeywords("music", k.musicScheduler.GetKeywordsFromMusics)
	k.processKeywords("news", k.newsScheduler.GetKeywordsFromNewsData)
	k.processKeywords("realtime_search", k.realtimeSearchScheduler.GetKeywordsFromRealtimeSearchData)
}

func (k *KeywordScheduler) processKeywords(category string, getKeywordsFunc func() ([]string, error)) {
	keywords, err := getKeywordsFunc()
	if err != nil {
		log.Printf("%s keywords 조회 실패: %v", category, err)
		return
	}

	keywordModel := []*model.Keywords{
		{
			Category:  category,
			Keywords:  keywords,
			CreatedAt: time.Now().Unix(),
		},
	}
	err = k.keywordRepository.SaveKeywords(keywordModel)
	if err != nil {
		log.Printf("%s keywords 저장 실패: %v", category, err)
		return
	}

	log.Printf("Save %s keywords: %v", category, keywords)
}
