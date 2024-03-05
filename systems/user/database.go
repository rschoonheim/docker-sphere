package systems_user

import (
	"docker_sphere/structures/structure_database"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
)

type User struct {
	Username string
	Password string
}

type UserDatabase interface {
	ByUsername(username string) (*User, error)
}

type InMemoryUserDatabase struct {
	Database *structure_database.Database
}

func (i *InMemoryUserDatabase) Persist() error {
	return i.Database.Save("userdb.json")
}

func (i *InMemoryUserDatabase) Load() error {
	return i.Database.Load("userdb.json")
}

func (i *InMemoryUserDatabase) ByUsername(username string) (*User, error) {
	data, ok := i.Database.Data[username]
	if !ok {
		return nil, nil
	}

	userData, ok := data.(string)
	if !ok {
		return nil, nil
	}

	var user User
	err := json.Unmarshal([]byte(userData), &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (i *InMemoryUserDatabase) CreateUser(username, password string) error {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("Error generating password hash", "error", err)
	}

	user := &User{
		Username: username,
		Password: string(hash),
	}

	userData, err := json.Marshal(user)
	if err != nil {
		return err
	}

	i.Database.Data[username] = string(userData)

	return nil
}

// RemoveUser - removes a user from the database
func (i *InMemoryUserDatabase) RemoveUser(username string) error {
	delete(i.Database.Data, username)
	i.Persist()
	return nil
}

// Authenticate - authenticates a user
func (i *InMemoryUserDatabase) Authenticate(username, password string) (bool, error) {
	user, err := i.ByUsername(username)
	if err != nil {
		return false, err
	}

	if user == nil {
		return false, nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, nil
	}

	return true, nil
}
