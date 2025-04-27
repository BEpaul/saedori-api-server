package repository

import (
	"context"
	"time"

	"github.com/bestkkii/saedori-api-server/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type ScheduleRepository struct {
	MongoDB *mongo.Client
}

func newScheduleRepository(mongoDB *mongo.Client) *ScheduleRepository {
	return &ScheduleRepository{
		MongoDB: mongoDB,
	}
}

func (r *ScheduleRepository) getCollection(collectionName string) *mongo.Collection {
	database := r.MongoDB.Database("saedori")
	collection := database.Collection(collectionName)
	return collection
}

func (r *ScheduleRepository) SaveMusic(schedule *model.CrawledMusic) error {
	collection := r.getCollection("Music")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, schedule)
	return err
}

func (r *ScheduleRepository) SaveNews(schedule *model.CrawledNews) error {
	collection := r.getCollection("News")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, schedule)
	return err
}

func (r *ScheduleRepository) SaveRealtimeSearch(schedule *model.RealtimeSearch) error {
	collection := r.getCollection("RealtimeSearch")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, schedule)
	return err
}
