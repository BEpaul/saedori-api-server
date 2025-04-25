package model

import "github.com/bestkkii/saedori-api-server/pkg"

type AllCategoriesResponse struct {
	*pkg.ApiResponse
	Musics                      []*Music                    `json:"music_summary"`
	RealtimeSearchDetailWrapper RealtimeSearchDetailWrapper `json:"realtime_search_summary"`
	News                        []*News                     `json:"news_summary"`
}
