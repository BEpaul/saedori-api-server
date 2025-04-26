package model

import (
	"github.com/bestkkii/saedori-api-server/pkg"
)

type DownloadData struct {
	Keywords        []*Keyword               `json:"dn_keywords"`
	News           []*News                  `json:"dn_news"`
	RealtimeSearch []*RealtimeSearchDownload `json:"dn_realtime_search"`
	Music          []*MusicDownload         `json:"dn_music"`
}

type DownloadDataResponse struct {
	*pkg.ApiResponse
	Result *DownloadData `json:"result"`
} 