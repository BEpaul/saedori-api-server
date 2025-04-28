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
	MongoDB *mongo.Client
}

func newKeywordRepository(mongoDB *mongo.Client) *KeywordRepository {
	return &KeywordRepository{
		MongoDB: mongoDB,
	}
}

func (k *KeywordRepository) GetKeywords() ([]*model.Keywords, error) {
	database := k.MongoDB.Database("saedori")
	collection := database.Collection("Keyword")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	categories := []string{"music", "realtime_search", "news", "coin"}
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

func (k *KeywordRepository) SaveKeywords(keywords []*model.Keywords) error {
	database := k.MongoDB.Database("saedori")
	collection := database.Collection("Keyword")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	for _, keyword := range keywords {
		_, err := collection.InsertOne(ctx, keyword)
		if err != nil {
			return err
		}
	}
	return nil
}
