package model

import (
	"github.com/bestkkii/saedori-api-server/pkg"
)

type MusicResponse struct {
    *pkg.ApiResponse
    Musics []*Music `json:"result" binding:"required"`
}

type Music struct {
	MusicData MusicRegion         `bson:"music" json:"music"`
}

type MusicRegion struct {
	Domestic []MusicDetail `bson:"domestic" json:"domestic"`
	Global   []MusicDetail `bson:"global" json:"global"`
}

type MusicDetail struct {
    Singer string `json:"singer" binding:"required"`
    Title  string `json:"title" binding:"required"`
    URL    string `json:"url" binding:"required"`
}

type MusicDownload struct {
	MusicData MusicRegion `bson:"music" json:"music"`
	CreatedAt int         `bson:"created_at" json:"created_at"`
}
