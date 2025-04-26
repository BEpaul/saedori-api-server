package repository

import (
	"context"
	"fmt"
	"sort"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/bestkkii/saedori-api-server/internal/model"
)

type RealtimeSearchRepository struct {
	MongoDB *mongo.Client
}

func newRealtimeSearchRepository(mongoDB *mongo.Client) *RealtimeSearchRepository {
	return &RealtimeSearchRepository{
		MongoDB: mongoDB,
	}
}

func (r *RealtimeSearchRepository) getCollection(collectionName string) *mongo.Collection {
	database := r.MongoDB.Database("saedori")
	collection := database.Collection(collectionName)
	return collection
}

// GetRealtimeSearchByDateRange returns realtime search data within the specified date range
func (r *RealtimeSearchRepository) GetRealtimeSearchByDateRange(startDate, endDate int64) ([]*model.RealtimeSearchDownload, error) {
	collection := r.getCollection("RealtimeSearch")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{
		"created_at": bson.M{
			"$gte": startDate,
			"$lte": endDate,
		},
	}

	cursor, err := collection.Find(ctx, filter, options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}}))
	if err != nil {
		return nil, fmt.Errorf("error getting realtime search: %v", err)
	}
	defer cursor.Close(ctx)

	// created_at별로 그룹화된 데이터를 저장할 맵
	groupedData := make(map[int]*model.RealtimeSearchDownload)

	for cursor.Next(ctx) {
		var search model.RealtimeSearch
		if err := cursor.Decode(&search); err != nil {
			return nil, fmt.Errorf("error decoding realtime search: %v", err)
		}

		// 해당 created_at의 데이터가 없으면 새로 생성
		if _, exists := groupedData[search.CreatedAt]; !exists {
			groupedData[search.CreatedAt] = &model.RealtimeSearchDownload{
				CreatedAt: search.CreatedAt,
				RealtimeSearch: model.RealtimeSearchDetailWrapper{
					RealtimeSearchDetail: model.RealtimeSearchDetail{
						KrSearchWords: []string{},
						UsSearchWords: []string{},
					},
				},
			}
		}

		// 국가별로 검색어 추가
		if search.Country == "kr" {
			groupedData[search.CreatedAt].RealtimeSearch.RealtimeSearchDetail.KrSearchWords = append(
				groupedData[search.CreatedAt].RealtimeSearch.RealtimeSearchDetail.KrSearchWords,
				search.SearchWord,
			)
		} else if search.Country == "us" {
			groupedData[search.CreatedAt].RealtimeSearch.RealtimeSearchDetail.UsSearchWords = append(
				groupedData[search.CreatedAt].RealtimeSearch.RealtimeSearchDetail.UsSearchWords,
				search.SearchWord,
			)
		}
	}

	// 맵을 슬라이스로 변환
	var results []*model.RealtimeSearchDownload
	for _, data := range groupedData {
		results = append(results, data)
	}

	// created_at 기준으로 내림차순 정렬
	sort.Slice(results, func(i, j int) bool {
		return results[i].CreatedAt > results[j].CreatedAt
	})

	return results, nil
} 