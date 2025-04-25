package scheduler

import (
	"log"
	"regexp"
	"strings"

	"github.com/bestkkii/saedori-api-server/internal/service"
)

type NewsScheduler struct {
	dashboardService *service.Dashboard
}

func NewsService(dashboardService *service.Dashboard) *NewsScheduler {
	return &NewsScheduler{
		dashboardService: dashboardService,
	}
}

// 뉴스 데이터로부터 오늘의 단어 3개 뽑기
func (n *NewsScheduler) GetKeywordsFromNewsData() ([]string, error) {
	// 뉴스 데이터 조회
	newsData, err := n.dashboardService.GetNewsDetails()
	if err != nil {
		log.Printf("뉴스 데이터 조회 실패: %v", err)
		return nil, err
	}

	// 결과를 저장할 배열
	keywords := make([]string, 0, 3)

	// 가장 최신 뉴스 데이터에서 title 3개 가져오기
	if len(newsData) > 0 && len(newsData[0].NewsItems) > 0 {
		for i, newsItem := range newsData[0].NewsItems {
			// 제목이 비어있는지 확인
			if newsItem.Title == "" {
				continue
			}

			// 제목 가공
			title := processNewsTitle(newsItem.Title)
			if title != "" {
				keywords = append(keywords, title)
				if len(keywords) >= 3 {
					break
				}
			}
		}
	}

	return keywords, nil
}

// 뉴스 제목 가공 함수
func processNewsTitle(title string) string {
	// 1. [ ]로 감싸진 문자열 제거
	re := regexp.MustCompile(`\[.*?\]`)
	title = re.ReplaceAllString(title, "")

	// 2. 모든 특수문자 제거 (한글, 영문, 숫자, 공백만 남김)
	re = regexp.MustCompile(`[^가-힣a-zA-Z0-9\s]`)
	title = re.ReplaceAllString(title, "")

	// 3. 앞뒤 공백 제거
	title = strings.TrimSpace(title)

	// 4. 연속된 공백을 하나로 치환
	re = regexp.MustCompile(`\s+`)
	title = re.ReplaceAllString(title, " ")

	// 5. 3어절까지만 남기기
	words := strings.Fields(title)
	if len(words) > 3 {
		words = words[:3]
	}

	return strings.Join(words, " ")
} 