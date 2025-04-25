package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/bestkkii/saedori-api-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DashboardRepository struct {
	MongoDB *mongo.Client
	NewsRepository *NewsRepository
	KeywordRepository *KeywordRepository
}

func newDashboardRepository() *DashboardRepository {
	client, err := ConnectMongoDB()
	if err != nil {
		fmt.Println(err)
	}

	return &DashboardRepository{
		MongoDB: client,
		NewsRepository: newNewsRepository(client),
		KeywordRepository: newKeywordRepository(client),
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

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("Error getting music list:", err)
		return nil, err
	}

	var musics []*model.Music
	for cursor.Next(ctx) {
		var music model.Music
		if err := cursor.Decode(&music); err != nil {
			fmt.Println("Error decoding music:", err)
			return nil, err
		}
		musics = append(musics, &music)
	}

	if err := cursor.Err(); err != nil {
		fmt.Println("Cursor error:", err)
		return nil, err
	}

	return musics, nil
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
			return nil, nil,err
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
        return nil, fmt.Errorf("error getting %s realtime search list: %v", country, err)
    }
    defer cursor.Close(ctx)

    var results []*model.RealtimeSearch
		
    for cursor.Next(ctx) {
        var realtimeSearch model.RealtimeSearch
        if err := cursor.Decode(&realtimeSearch); err != nil {
            return nil, fmt.Errorf("error decoding %s realtime search: %v", country, err)
        }
        results = append(results, &realtimeSearch)
    }

    if err := cursor.Err(); err != nil {
        return nil, fmt.Errorf("%s cursor error: %v", country, err)
    }
	
    return results, nil
}
