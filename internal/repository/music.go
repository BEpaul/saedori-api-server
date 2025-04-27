package repository

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/bestkkii/saedori-api-server/internal/model"
)

type MusicRepository struct {
	MongoDB *mongo.Client
}

func newMusicRepository(mongoDB *mongo.Client) *MusicRepository {
	return &MusicRepository{
		MongoDB: mongoDB,
	}
}

func (m *MusicRepository) getCollection(collectionName string) *mongo.Collection {
	database := m.MongoDB.Database("saedori")
	collection := database.Collection(collectionName)
	return collection
}

// GetMusicByDateRange returns music data within the specified date range
func (m *MusicRepository) GetMusicByDateRange(startDate, endDate int64) ([]*model.MusicDownload, error) {
	collection := m.getCollection("Music")
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
		log.Fatalf("error getting music: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []*model.MusicDownload
	if err = cursor.All(ctx, &results); err != nil {
		log.Fatalf("error decoding music: %v", err)
		return nil, err
	}

	return results, nil
}
