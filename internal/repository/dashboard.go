package repository

import (
	"context"
	"log"
	"time"

	"github.com/bestkkii/saedori-api-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DashboardRepository struct {
	MongoDB                  *mongo.Client
	NewsRepository           *NewsRepository
	KeywordRepository        *KeywordRepository
	RealtimeSearchRepository *RealtimeSearchRepository
	MusicRepository          *MusicRepository
	ScheduleRepository       *ScheduleRepository
}

func newDashboardRepository() *DashboardRepository {
	client, err := ConnectMongoDB()
	if err != nil {
		log.Fatal(err)
	}

	return &DashboardRepository{
		MongoDB:                  client,
		NewsRepository:           newNewsRepository(client),
		KeywordRepository:        newKeywordRepository(client),
		RealtimeSearchRepository: newRealtimeSearchRepository(client),
		MusicRepository:          newMusicRepository(client),
		ScheduleRepository:       newScheduleRepository(client),
	}
}

func (d *DashboardRepository) getCollection(collectionName string) *mongo.Collection {
	database := d.MongoDB.Database("saedori")
	collection := database.Collection(collectionName)
	return collection
}

func (d *DashboardRepository) GetMusics() ([]*model.Music, error) {
	collection := d.getCollection("Music")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	opts := options.FindOne().SetSort(bson.D{{Key: "created_at", Value: -1}})
	var music model.Music
	err := collection.FindOne(ctx, bson.M{}, opts).Decode(&music)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		log.Fatalf("Error getting music: %v", err)
		return nil, err
	}

	return []*model.Music{&music}, nil
}

func (d *DashboardRepository) GetRealtimeSearches() ([]*model.RealtimeSearch, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	realtimeSearchList, err := d.GetRealtimeSearchesByCountry(ctx, "kr", 5)

	if err != nil {
		return nil, err
	}

	return realtimeSearchList, nil
}

func (d *DashboardRepository) GetRealtimeSearchDetails() ([]*model.RealtimeSearch, []*model.RealtimeSearch, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	krList, err := d.GetRealtimeSearchesByCountry(ctx, "kr", 10)
	if err != nil {
		return nil, nil, err
	}

	usList, err := d.GetRealtimeSearchesByCountry(ctx, "us", 10)
	if err != nil {
		return nil, nil, err
	}

	return krList, usList, nil
}

// 국가별로 데이터를 조회하는 함수
func (d *DashboardRepository) GetRealtimeSearchesByCountry(ctx context.Context, country string, count int64) ([]*model.RealtimeSearch, error) {
	collection := d.getCollection("RealtimeSearch")

	cursor, err := collection.Find(ctx,
		bson.M{"country": country},
		options.Find().
			SetSort(bson.D{
				{"created_at", -1},
				{"rank", 1},
			}).
			SetLimit(count),
	)
	if err != nil {
		log.Fatalf("error getting %s realtime search list: %v", country, err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []*model.RealtimeSearch

	for cursor.Next(ctx) {
		var realtimeSearch model.RealtimeSearch
		if err := cursor.Decode(&realtimeSearch); err != nil {
			log.Fatalf("error decoding %s realtime search: %v", country, err)
			return nil, err
		}
		results = append(results, &realtimeSearch)
	}

	if err := cursor.Err(); err != nil {
		log.Fatalf("%s cursor error: %v", country, err)
		return nil, err
	}

	return results, nil
}

// GetKeywordsByDateRange returns keywords within the specified date range
func (d *DashboardRepository) GetKeywordsByDateRange(startDate, endDate int64, categories []string) ([]*model.Keyword, error) {
	collection := d.getCollection("Keyword")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 기본 필터: created_at 범위
	filter := bson.M{
		"created_at": bson.M{
			"$gte": startDate,
			"$lte": endDate,
		},
	}

	// 카테고리 매핑
	categoryMap := map[string]string{
		"news":            "news",
		"realtime-search": "search_word",
		"music":           "music",
		"coin":            "coin",
	}

	// 카테고리 필터 추가
	if len(categories) > 0 {
		var categoryFilters []bson.M
		for _, category := range categories {
			if mongoCategory, exists := categoryMap[category]; exists {
				categoryFilters = append(categoryFilters, bson.M{"category": mongoCategory})
			}
		}
		if len(categoryFilters) > 0 {
			filter["$or"] = categoryFilters
		}
	}

	cursor, err := collection.Find(ctx, filter, options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}}))
	if err != nil {
		log.Fatalf("error getting keywords: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var keywords []*model.Keyword
	if err = cursor.All(ctx, &keywords); err != nil {
		log.Fatalf("error decoding keywords: %v", err)
		return nil, err
	}

	return keywords, nil
}
