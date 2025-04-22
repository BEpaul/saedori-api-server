package model

import "time"

type Keyword struct {
	Keyword   string    `json:"keyword"`
	CreatedAt time.Time `json:"created_at"`
}
