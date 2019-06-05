package main

/*
   EspiSite - a quick and dirty CMS
   Copyright (C) 2019 EspiDev

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

const (
	DBLocation = "./db.json"
)

type IDatabase struct {
	Posts []*IPost `json:"posts"`
	Users []*IUser `json:"users"`
}

type PostID struct {
	IDYear  string `json:"id_year"`
	IDMonth string `json:"id_month"`
	IDDay   string `json:"id_day"`
	IDNum   string `json:"id_num"`
}

type IPost struct {
	Name       string   `json:"name"`
	UserID     string   `json:"userid"`
	Categories []string `json:"categories"`

	ID PostID `json:"id"`

	TimeCreated int64 `json:"timecreated"`
	TimeUpdated int64 `json:"timeupdated"`

	Icon    string `json:"icon"`
	Content string `json:"content"`

	Visibility bool `json:"visibility"`
}

type IUser struct {
	DisplayName string `json:"displayname"`
	UserName    string `json:"username"`
	Password    string `json:"password"`

	TimeRegistered int64    `json:"timeregistered"`
	Posts          []string `json:"posts"`

	Icon        string `json:"icon"`
	Description string `json:"description"`
	IsAdmin     bool   `json:"is_admin"`
}

func (ip *IPost) isID(id PostID) bool {
	return ip.ID.IDDay == id.IDDay && ip.ID.IDMonth == id.IDMonth && ip.ID.IDYear == id.IDYear && ip.ID.IDNum == id.IDNum
}

func GetPost(id PostID) (*IPost, error) {
	for _, post := range db.Posts {
		if post.isID(id) {
			return post, nil
		}
	}
	return nil, errors.New("not found")
}

func LoadDB() {

	if _, err := os.Stat(DBLocation); os.IsNotExist(err) {
		StoreDB()
	}

	bV, err := ioutil.ReadFile(DBLocation)
	if err != nil {
		log.Fatalf("Cannot load database: %s\n", err)
	}
	err = json.Unmarshal(bV, &db)
	if err != nil {
		log.Fatalf("Error unmarshalling db from json: %s\n", err)
	}

	for _, post := range db.Posts {
		n, err := strconv.ParseInt(post.ID.IDNum, 10, 64)
		if err != nil {
			continue
		}
		if postNum <= n {
			postNum = n + 1
		}
	}
}

func StoreDB() {
	if _, err := os.Stat(DBLocation); !os.IsNotExist(err) {
		err := os.Rename(DBLocation, DBLocation+".backup")
		if err != nil {
			log.Fatalf("Cannot create backup: %s\n", err)
		}
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
