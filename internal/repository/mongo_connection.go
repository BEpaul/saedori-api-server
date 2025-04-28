package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongoUri = "mongodb+srv://saedori:toehfldhkrhdwnsla@saedori-org.7alxffm.mongodb.net/?retryWrites=true&w=majority&appName=saedori-org"

func ConnectMongoDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(mongoUri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Error connecting to MongoDB:", err)
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("Error pinging MongoDB:", err)
		return nil, err
	}

	log.Println("Connected to MongoDB!")
	return client, nil
}
