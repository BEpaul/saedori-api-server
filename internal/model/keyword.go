package model

import (
	"time"

	"github.com/bestkkii/saedori-api-server/pkg"
)

type Keyword struct {
	Keyword   string    `json:"keyword" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"required"`
}

type GetKeywordListResponse struct {
	*pkg.ApiResponse
	Keywords []*Keyword `json:"result"`
}
