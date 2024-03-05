package orchestrators

import (
	"docker_sphere/systems/mongodb"
	"docker_sphere/systems/user"
	"go.mongodb.org/mongo-driver/bson"
	"log/slog"
)

// CreateUser - tries to create a new user and persist it to the database
func CreateUser(username string, password string) error {

	/**
	 * Ask the database service if there is already a user
	 * with the same username
	 */
	type x struct {
		name string `bson:"some_string,omitempty"`
	}

	res := mongodb.Search(
		"users",
		bson.M{"username": username},
	)

	// print raw result
	println(res.Raw())

	//y := x{}
	//
	//println(res.Raw())
	//
	//res.Decode(y)
	//
	//println(y.name)

	// Create a new user
	user := user.User{
		Username: username,
		Password: password,
	}

	// Persist the user to the database
	err := mongodb.Persist("users", user)
	if err != nil {
		slog.Error("Failed to persist the user to the database")
		return err
	}

	return nil
}
