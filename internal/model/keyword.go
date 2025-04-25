package model

import (
	"github.com/bestkkii/saedori-api-server/pkg"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MongoDB Document 구조체
type Keyword struct {
	// ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CreatedAt int64    `bson:"created_at" json:"created_at"`
	Keyword   []string `bson:"keyword" json:"keyword"`
	Category  string   `bson:"category" json:"category"`
}

// API 응답 구조체
type KeywordResponse struct {
	Category string   `json:"category"`
	Keyword  []string `json:"keyword"`
}

type GetKeywordListResponse struct {
	*pkg.ApiResponse
	Keywords []KeywordResponse `json:"top3_keywords"`
}
