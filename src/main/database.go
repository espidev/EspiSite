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
	Posts []IPost
	Users []IUser
}

type IPost struct {
	Name string
	UUID string
	UserID string
	Categories []string

	TimeCreated int64
	TimeUpdated int64

	Icon string
	Content string
}

type IUser struct {
	Name string
	UUID string

	TimeRegistered int64
	Posts []string

	Icon string
	Description string
}

func LoadDB() {

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