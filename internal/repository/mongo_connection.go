package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongoUri = "mongodb+srv://saedori:toehfldhkrhdwnsla@saedori-org.7alxffm.mongodb.net/?retryWrites=true&w=majority&appName=saedori-org"

func ConnectMongoDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(mongoUri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to MongoDB: %v", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, fmt.Errorf("Error pinging MongoDB: %v", err)
	}

	fmt.Println("Connected to MongoDB!")
	return client, nil
}
