package model

import (
	"github.com/bestkkii/saedori-api-server/pkg"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RealtimeSearch struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Country    string             `bson:"country" json:"country" binding:"required"`
	SearchWord string             `bson:"search_word" json:"search_word" binding:"required"`
	Rank       int64              `bson:"rank" json:"rank" binding:"required"`
	CreatedAt  int                `bson:"created_at" json:"created_at" binding:"required"`
}

type RealtimeSearchDetail struct {
	KrSearchWords []string `json:"kr"`
	UsSearchWords []string `json:"us"`
}

type RealtimeSearchDetailWrapper struct {
	RealtimeSearchDetail RealtimeSearchDetail `json:"realtime_search"`
}

type RealtimeSearchDetailResponse struct {
	*pkg.ApiResponse
	RealtimeSearchDetailWrapper RealtimeSearchDetailWrapper `json:"result"`
}

type RealtimeSearchDownload struct {
	CreatedAt      int                  `bson:"created_at" json:"created_at"`
	RealtimeSearch RealtimeSearchDetail `json:"realtime_search"`
}

type CrawledRealtimeSearch struct {
	RealtimeSearchDetail RealtimeSearchDetail `bson:"realtime_search" json:"realtime_search"`
	CreatedAt            int64                `bson:"created_at" json:"created_at"`
}
