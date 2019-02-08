package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

const (
	DBLocation = "./db.json"
)

type IDatabase struct {
	Posts []IPost `json:"posts"`
	Users []IUser `json:"users"`
}

type IPost struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
	UserID string `json:"userid"`
	Categories []string `json:"categories"`

	TimeCreated int64 `json:"timecreated"`
	TimeUpdated int64 `json:"timeupdated"`

	Icon string `json:"icon"`
	Content string `json:"content"`
}

type IUser struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`

	TimeRegistered int64 `json:"timeregistered"`
	Posts []string `json:"posts"`

	Icon string `json:"icon"`
	Description string `json:"description"`
}

func LoadDB() {
	bV, err := ioutil.ReadFile(DBLocation)
	if err != nil {
		log.Fatalf("Cannot load database: %s\n", err)
	}
	err = json.Unmarshal(bV, &db)
	if err != nil {
		log.Fatalf("Error unmarshalling db from json: %s\n", err)
	}
}

func StoreDB() {
	err := os.Rename(DBLocation, DBLocation + ".backup")
	if err != nil {
		log.Fatalf("Cannot create backup: %s\n", err)
	}
	b, err := json.Marshal(db)
	if err != nil {
		log.Printf("Cannot marshal db to JSON: %s\n", err)
		return
	}
	err = ioutil.WriteFile(DBLocation, []byte(b), 0644)
	if err != nil {
		log.Fatalf("Cannot write DB to file %s\n", err)
	}
	err = os.Remove(DBLocation + ".backup")
	if err != nil {
		log.Printf("Cannot delete backup: %s\n", err)
	}
}