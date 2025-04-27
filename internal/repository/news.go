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

type NewsRepository struct {
	MongoDB *mongo.Client
}

func newNewsRepository(mongoDB *mongo.Client) *NewsRepository {
	return &NewsRepository{
		MongoDB: mongoDB,
	}
}

func (n *NewsRepository) getCollection(collectionName string) *mongo.Collection {
	database := n.MongoDB.Database("saedori")
	collection := database.Collection(collectionName)
	return collection
}

func (n *NewsRepository) GetNewsDetails() (*model.News, error) {
	database := n.MongoDB.Database("saedori")
	collection := database.Collection("News")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// created_at 인덱스를 활용하여 최신 데이터 조회
	opts := options.FindOne().SetSort(bson.D{{Key: "created_at", Value: -1}})

	var news model.News
	err := collection.FindOne(ctx, bson.M{}, opts).Decode(&news)
	if err != nil {
		log.Fatalf("Error getting latest news:", err)
		return nil, err
	}

	return &news, nil
}

// GetNewsByDateRange returns news within the specified date range
func (n *NewsRepository) GetNewsByDateRange(startDate, endDate int64) ([]*model.News, error) {
	collection := n.getCollection("News")
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
		log.Fatalf("error getting news: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []*model.News
	if err = cursor.All(ctx, &results); err != nil {
		log.Fatalf("error decoding news: %v", err)
		return nil, err
	}

	return results, nil
}
