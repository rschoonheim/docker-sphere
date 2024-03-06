package orchestrators

import (
	"context"
	"docker_sphere/project/systems/mongodb"
	"docker_sphere/project/systems/security"
	"docker_sphere/project/systems/user"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
)

type AuthenticationResult struct {
	Username string
	Token    string
}

// Login - tries to authenticate a user with a username and password when authentication succeeds, it returns
// a AuthenticationResult containing the username and a token otherwise it returns an error
func Login(username string, password string) (AuthenticationResult, error) {

	database := mongodb.GetClient()

	// Search user
	userExists := database.Database("docker-sphere").Collection("users").FindOne(context.TODO(), bson.D{{"username", username}})
	if userExists.Err() != nil {
		return AuthenticationResult{}, errors.New("User does not exist")
	}

	// Decode user
	var userExistsDecoded user.User
	err := userExists.Decode(&userExistsDecoded)
	if err != nil {
		return AuthenticationResult{}, errors.New("Error decoding user")
	}

	// Check password
	if !security.VerifyPassword(password, userExistsDecoded.Password) {
		return AuthenticationResult{}, errors.New("Password is incorrect")
	}

	// Return token
	return AuthenticationResult{
		Username: username,
		Token:    security.JwtCreate(username),
	}, nil
}
