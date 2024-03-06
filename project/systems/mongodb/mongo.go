package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"
)

var database *mongo.Database

func Initialize() {
	slog.Info("Initializing MongoDB")

	clientOptions := options.Client().ApplyURI("mongodb://root:root@localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	database = client.Database("docker-sphere")

	slog.Info("MongoDB initialized successfully")
}

func GetClient() *mongo.Client {
	return database.Client()
}
