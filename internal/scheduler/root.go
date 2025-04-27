package scheduler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/bestkkii/saedori-api-server/internal/model"
	"github.com/bestkkii/saedori-api-server/internal/repository"
)

type Dashboard struct {
	DashboardRepository *repository.DashboardRepository
}

func (d *Dashboard) StartScheduler() {
	go func() {
		for {
			now := time.Now()
			next := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 10, 0, 0, now.Location())
			if now.After(next) {
				next = next.Add(1 * time.Hour)
			}
			duration := next.Sub(now)
			time.AfterFunc(duration, func() {
				d.fetchData()
			})
			time.Sleep(duration + 1*time.Second)
		}
	}()
}

func (d *Dashboard) fetchData() {
	resp, err := http.Get("http://localhost:8000/api/v1/crawl")
	if err != nil {
		log.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return
	}

	var jsonData interface{}
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		log.Println("JSON 파싱 에러:", err)
		return
	}

	d.processData(jsonData)
}

func (d *Dashboard) processData(jsonData interface{}) error {
	data, ok := jsonData.(map[string]interface{})
	if !ok {
		return nil
	}

	processCrawlData := func(crawlData map[string]interface{}, dataType string, saveFunc func(interface{}) error) {
		if crawlingStatus, exists := crawlData["crawling"].(string); exists && crawlingStatus == "Success" {
			resultData := crawlData["result"].(map[string]interface{})[dataType]
			err := saveFunc(resultData)
			if err != nil {
				log.Printf("%s 데이터 저장 실패: %v", dataType, err)
			} else {
				log.Printf("%s 데이터 저장 성공: %v", dataType, resultData)
			}
		}
	}

	if musicCrawl, exists := data["music_crawl"].(map[string]interface{}); exists {
		processCrawlData(musicCrawl, "music", func(resultData interface{}) error {
			musicData := resultData.(map[string]interface{})
			crawledMusic := &model.CrawledMusic{
				Music: model.MusicRegion{
					Domestic: parseMusicDetails(musicData["domestic"].([]interface{})),
					Global:   parseMusicDetails(musicData["global"].([]interface{})),
				},
				CreatedAt: int64(data["created_at"].(float64)),
			}
			return d.DashboardRepository.ScheduleRepository.SaveMusic(crawledMusic)
		})
	}

	if newsCrawl, exists := data["news_crawl"].(map[string]interface{}); exists {
		processCrawlData(newsCrawl, "news", func(resultData interface{}) error {
			newsData := resultData.([]interface{})
			crawledNews := &model.CrawledNews{
				NewsItems: parseNewsItems(newsData),
				CreatedAt: int64(data["created_at"].(float64)),
			}
			return d.DashboardRepository.ScheduleRepository.SaveNews(crawledNews)
		})
	}

	if realtimeSearchWordsCrawl, exists := data["realtime_search_words_crawl"].(map[string]interface{}); exists {
		processCrawlData(realtimeSearchWordsCrawl, "realtime_search_words", func(resultData interface{}) error {
			realtimeSearchWordsData := resultData.(map[string]interface{})
			krData := realtimeSearchWordsData["kr"].([]interface{})
			usData := realtimeSearchWordsData["us"].([]interface{})
			createdAt := int64(data["created_at"].(float64))

			saveRealtimeSearch := func(country string, data []interface{}) {
				for _, item := range data {
					word := item.(map[string]interface{})
					rankStr := word["rank"].(string)
					rank, err := strconv.Atoi(rankStr)
					if err != nil {
						log.Println("Rank 변환 실패:", err)
						rank = 0
					}
					realtimeSearch := model.RealtimeSearch{
						Country:    country,
						SearchWord: word["search_word"].(string),
						Rank:       int64(rank),
						CreatedAt:  int(createdAt),
					}

					err = d.DashboardRepository.ScheduleRepository.SaveRealtimeSearch(&realtimeSearch)
					if err != nil {
						log.Println("실시간 검색어 데이터 저장 실패:", err)
					} else {
						log.Println("실시간 검색어 데이터 저장 성공:", realtimeSearch)
					}
				}
			}

			saveRealtimeSearch("kr", krData)
			saveRealtimeSearch("us", usData)
			return nil
		})
	}

	return nil
}

func parseMusicDetails(data []interface{}) []model.MusicDetail {
	var details []model.MusicDetail
	for _, m := range data {
		musicInfo := m.(map[string]interface{})
		title, _ := musicInfo["title"].(string)
		singer, _ := musicInfo["singer"].(string)
		url, _ := musicInfo["url"].(string)

		details = append(details, model.MusicDetail{
			Title:  title,
			Singer: singer,
			URL:    url,
		})
	}
	return details
}

func parseNewsItems(data []interface{}) []model.NewsItem {
	var items []model.NewsItem
	for _, item := range data {
		newsItem := item.(map[string]interface{})
		company, _ := newsItem["company"].(string)
		title, _ := newsItem["title"].(string)
		url, _ := newsItem["url"].(string)
		lead, _ := newsItem["lead"].(string)

		items = append(items, model.NewsItem{
			Company: company,
			Title:   title,
			URL:     url,
			Lead:    lead,
		})
	}
	return items
}
