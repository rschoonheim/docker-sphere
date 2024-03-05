package structure_database

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Database struct {
	Data map[string]interface{}
}

func NewDatabase() *Database {
	return &Database{
		Data: make(map[string]interface{}),
	}
}

func (db *Database) Load(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &db.Data)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) Save(filename string) error {
	data, err := json.Marshal(db.Data)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, data, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
