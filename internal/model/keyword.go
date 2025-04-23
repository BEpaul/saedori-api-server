package model

import (
	"github.com/bestkkii/saedori-api-server/pkg"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Keyword struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Top3Keywords []string           `bson:"top3_keywords" json:"top3_keywords" binding:"required"`
	CreatedAt    int64              `bson:"created_at" json:"created_at" binding:"required"`
}

type GetKeywordListResponse struct {
	*pkg.ApiResponse
	Keywords []*Keyword `json:"result"`
}
