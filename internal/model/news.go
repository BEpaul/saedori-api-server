package model

import (
	"regexp"
	"strings"

	"github.com/bestkkii/saedori-api-server/pkg"
)

// FastAPI 응답을 위한 모델
type FastAPIResponse struct {
	*pkg.ApiResponse // Message string `json:"message"`
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
	CreatedAt int        `bson:"created_at" json:"created_at" binding:"required"`
	NewsItems []NewsItem `bson:"news" json:"news" binding:"required"`
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
