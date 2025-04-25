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

type NewsRepository struct {
	MongoDB *mongo.Client
}

func newNewsRepository(mongoDB *mongo.Client) *NewsRepository {
	return &NewsRepository{
		MongoDB: mongoDB,
	}
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
		fmt.Println("Error getting latest news:", err)
		return nil, err
	}

	return &news, nil
} 