package structure_database_test

import (
	"docker_sphere/structures/structure_database"
	"os"
	"testing"
)

func TestLoadAndSave(t *testing.T) {
	db := structure_database.NewDatabase()

	key := "testkey"
	value := "testvalue"

	db.Data[key] = value

	err := db.Save("testfile.json")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	db.Data = make(map[string]interface{})

	err = db.Load("testfile.json")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if db.Data[key] != value {
		t.Errorf("Expected value %s, got %s", value, db.Data[key])
	}

	// Clean up the test file
	os.Remove("testfile.json")
}
