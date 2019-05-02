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
	"context"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

var (
	router *gin.Engine
	config Config
	db     IDatabase
	postNum int64
)

const (
	RootRepoFolder = "."
)

func main() {
	log.Printf("Starting EspiSite...\n")

	// License disclaimer
	log.Println(`This program comes with ABSOLUTELY NO WARRANTY;
This is free software, and you are welcome to redistribute it under certain conditions.`)

	// Load config and data
	setupConfig()
	LoadDB()

	// Init web-server
	router = gin.Default()
	setupRoutes()

	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(config.Port),
		Handler: router,
	}

	go func() { // start web-server in goroutine
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// listen for sigint to shutdown gracefully
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down EspiSite...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown: ", err)
	}
	log.Println("EspiSite has stopped.")
}

func setupRoutes() {
	router.LoadHTMLGlob(RootRepoFolder + "/src/html/*")

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H {})
	})
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	router.GET("/test", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{})
	})
	router.GET(config.AdminRoute, func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin.html", gin.H{})
	})

	router.GET(config.AdminRoute+"/new-post", func(c *gin.Context) {
		c.HTML(http.StatusOK, "new-post.html", gin.H{
			"title":   "post title",
			"content": template.HTML("<p>this is some good quality content right here!</p>"),
		})
	})

	router.POST(config.AdminRoute+"/new-post", func(c *gin.Context) {

		t := time.Now()

		id := PostID{
			IDYear:  strconv.Itoa(t.Year()),
			IDMonth: t.Month().String(),
			IDDay:   strconv.Itoa(t.Day()),
			IDNum:   strconv.FormatInt(postNum, 10),
		}

		postNum++

		post := IPost{
			Name:        c.PostForm("title"),
			UserID:      "",
			Categories:  []string{},
			ID:          id,
			TimeCreated: t.Unix(),
			TimeUpdated: t.Unix(),
			Icon:        "",
			Content:     c.PostForm("content"),
		}

		db.Posts = append(db.Posts, &post)

		c.Redirect(301, "/posts/"+id.IDYear+"/"+id.IDMonth+"/"+id.IDDay+"/"+id.IDNum)

		go func() {
			StoreDB()
		}()
	})

	router.GET("/posts", func(c *gin.Context) {
		c.HTML(http.StatusOK, "blog.html", db.Posts)
	})

	router.GET("/posts/:year/:month/:day/:num", func(c *gin.Context) {
		id := PostID{IDYear: c.Params.ByName("year"),
			IDDay:   c.Params.ByName("day"),
			IDMonth: c.Params.ByName("month"),
			IDNum:   c.Params.ByName("num")}

		post, err := GetPost(id)
		if err != nil {
			c.HTML(http.StatusNotFound, "404.html", gin.H{})
			return
		}

		c.HTML(http.StatusOK, "post.html", gin.H{
			"postName":    post.Name,
			"timeUpdated": time.Unix(post.TimeUpdated, 0),
			"content":     template.HTML(post.Content),
		})

	})
}
