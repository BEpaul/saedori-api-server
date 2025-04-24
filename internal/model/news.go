package model

import (
	"regexp"
	"strings"

	"github.com/bestkkii/saedori-api-server/pkg"
	"go.mongodb.org/mongo-driver/bson/primitive" // 오늘의 단어
)

// FastAPI 응답을 위한 모델
type FastAPIResponse struct {
	Message string `json:"message"`
	CreatedAt int64 `json:"created_at"`
	NewsCrawl NewsCrawlResponse `json:"news_crawl"`
}

type NewsCrawlResponse struct {
	Crawling string `json:"crawling"`
	Result   NewsResult `json:"result"`
}

type NewsResult struct {
	News []NewsItem `json:"news"`
}

// MongoDB 저장을 위한 모델
type News struct {
	CreatedAt int              `bson:"created_at" json:"created_at" binding:"required"`
	NewsItems []NewsItem        `bson:"news" json:"news" binding:"required"`
}

type NewsItem struct {
	Company string `bson:"company" json:"company" binding:"required"`
	Title   string `bson:"title" json:"title" binding:"required"`
	Lead    string `bson:"lead" json:"lead" binding:"required"`
	URL     string `bson:"url" json:"url" binding:"required"`
}

// API 응답을 위한 모델
type NewsResponse struct {
	*pkg.ApiResponse
	News []*News `json:"result" binding:"required"`
}

// 키워드 처리를 위한 모델
type NewsKeywords struct {
	CreatedAt int               `bson:"created_at" json:"created_at"`
	Keywords  []string          `bson:"keyword" json:"keyword"`
	Category  string            `bson:"category" json:"category"`
}

// ProcessNewsKeywords는 뉴스 제목에서 키워드를 추출하고 가공합니다.
func ProcessNewsKeywords(newsItems []NewsItem) []string {
	keywords := make([]string, 0, 3)
	
	// 최대 3개의 뉴스 제목에서 키워드 추출
	for i := 0; i < 3 && i < len(newsItems); i++ {
		title := newsItems[i].Title
		
		// [ ]로 감싸진 문자열 제거
		re := regexp.MustCompile(`\[.*?\]`)
		title = re.ReplaceAllString(title, "")
		
		// 앞쪽 공백 제거
		title = strings.TrimSpace(title)
		
		// 모든 특수문자 제거 (한글, 영문, 숫자, 공백만 남김)
		re = regexp.MustCompile(`[^가-힣a-zA-Z0-9\s]`)
		title = re.ReplaceAllString(title, "")
		
		// 3어절까지만 추출
		words := strings.Fields(title)
		if len(words) > 3 {
			words = words[:3]
		}
		
		keywords = append(keywords, strings.Join(words, " "))
	}
	
	return keywords
}
