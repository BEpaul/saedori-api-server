package repository

import (
	"context"
	"time"

	"github.com/bestkkii/saedori-api-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type KeywordRepository struct {
	mongoDB *mongo.Client
}

func newKeywordRepository(mongoDB *mongo.Client) *KeywordRepository {
	return &KeywordRepository{
		mongoDB: mongoDB,
	}
}

func (k *KeywordRepository) GetKeywords() ([]*model.Keywords, error) {
	database := k.mongoDB.Database("saedori")
	collection := database.Collection("Keyword")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	categories := []string{"music", "search_word", "news", "coin"}
	var keywords []*model.Keywords

	for _, category := range categories {
		opts := options.FindOne().SetSort(bson.D{{Key: "created_at", Value: -1}})
		var keyword model.Keywords
		err := collection.FindOne(ctx, bson.M{"category": category}, opts).Decode(&keyword)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				continue
			}
			return nil, err
		}
		keywords = append(keywords, &keyword)
	}

	return keywords, nil
} 