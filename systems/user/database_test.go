package systems_user_test

import (
	"docker_sphere/structures/structure_database"
	"docker_sphere/systems/user"
	"golang.org/x/crypto/bcrypt"
	"os"
	"testing"
)

func TestCreateUserAndByUsername(t *testing.T) {
	db := &systems_user.InMemoryUserDatabase{
		Database: structure_database.NewDatabase(),
	}

	username := "testuser"
	password := "testpassword"

	err := db.CreateUser(username, password)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = db.Persist()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	db.Database.Data = make(map[string]interface{})

	err = db.Load()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	user, err := db.ByUsername(username)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if user == nil {
		t.Errorf("Expected user, got nil")
	}

	if user.Username != username {
		t.Errorf("Expected username %s, got %s", username, user.Username)
	}

	// Compare the password with the password hash
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		t.Errorf("Could not validate hashed password: %v", err)
	}

	// Clean up the test file
	os.Remove("userdb.json")
}

// Test delete user
func TestDeleteUser(t *testing.T) {
	db := &systems_user.InMemoryUserDatabase{
		Database: structure_database.NewDatabase(),
	}

	username := "testuser"
	password := "testpassword"

	err := db.CreateUser(username, password)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = db.Persist()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	db.Database.Data = make(map[string]interface{})

	err = db.Load()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	user, err := db.ByUsername(username)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if user == nil {
		t.Errorf("Expected user, got nil")
	}

	err = db.RemoveUser(username)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	user, err = db.ByUsername(username)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if user != nil {
		t.Errorf("Expected nil, got %v", user)
	}

	os.Remove("userdb.json")
}

func TestAuthenticateUser(t *testing.T) {
	db := &systems_user.InMemoryUserDatabase{
		Database: structure_database.NewDatabase(),
	}

	username := "testuser"
	password := "testpassword"

	err := db.CreateUser(username, password)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = db.Persist()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	db.Database.Data = make(map[string]interface{})

	err = db.Load()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	user, err := db.Authenticate(username, password)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if user == false {
		t.Errorf("Expected true, got false")
	}

	os.Remove("userdb.json")
}
