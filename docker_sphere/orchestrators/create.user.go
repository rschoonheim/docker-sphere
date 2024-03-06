package orchestrators

import (
	"context"
	systems_mongodb "docker_sphere/systems/mongodb"
	"docker_sphere/systems/security"
	"docker_sphere/systems/user"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"log/slog"
)

// CreateUser - tries to create a new user and persist it to the database
func CreateUser(username string, password string) (*user.User, error) {

	mongodb := systems_mongodb.GetClient()

	// Check if the user already exists in the database.
	userExists := mongodb.Database("docker-sphere").Collection("users").FindOne(context.TODO(), bson.D{{"username", username}})
	if userExists.Err() == nil {
		slog.Warn("Tried to create a user with an existing username")
		return nil, errors.New("User already exists in the database")
	}

	userDocument := user.User{
		Username: username,
		Password: security.HashPassword(password),
	}

	// Insert the user into the database
	_, err := mongodb.Database("docker-sphere").Collection("users").InsertOne(context.TODO(), userDocument)
	if err != nil {
		slog.Error("Error inserting user into the database: ", err)
		return nil, errors.New("Error inserting user into the database")
	}

	return &userDocument, nil
}
