package scheduler

import (
	"log"
	"github.com/bestkkii/saedori-api-server/internal/service"
)

type MusicScheduler struct {
	dashboardService *service.Dashboard
}

func MusicService(dashboardService *service.Dashboard) *MusicScheduler {
	return &MusicScheduler{
		dashboardService: dashboardService,
	}
}

func (d *Dashboard) GetKeywordsFromMusics() ([]string, error) {
	musics, err := d.dashboardRepository.GetMusics()
	if err != nil {
		return nil, err
	}

	keywords := make([]string, 0)
	for _, music := range musics {
		if len(music.MusicData.Domestic) > 0 {
			keywords = append(keywords, music.MusicData.Domestic[0].Title)		//국내 1위
		}
		if len(music.MusicData.Global) > 0 {
			keywords = append(keywords, music.MusicData.Global[0].Title)		//해외 1위
		}
		if len(music.MusicData.Domestic) > 0 {
			keywords = append(keywords, music.MusicData.Domestic[1].Title)		//국내 2위
		}
	}

	if len(keywords) > 3 {
		keywords = keywords[:3]
	}

	return keywords, nil
}