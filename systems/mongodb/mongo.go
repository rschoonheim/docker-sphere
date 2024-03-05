package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database *mongo.Database

func Initialize() {
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

}

// Persist - persist a document in the database
func Persist(collection string, document interface{}) error {
	_, err := database.Collection(collection).InsertOne(context.Background(), document)
	return err
}

// Search - search for a document in the database
func Search(collection string, filter interface{}) *mongo.SingleResult {
	return database.Collection(collection).FindOne(context.Background(), filter)
}
