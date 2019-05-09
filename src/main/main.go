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

	// start web-server in goroutine
	go func() {
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
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.StaticFile("/html/header.html", "./src/html/header.html")
	router.Static("/css", "./src/css")
	router.Static("/js", "./src/js")
	router.Static("/assets", "./src/assets")

	router.LoadHTMLGlob(RootRepoFolder + "/src/html/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	PostRoutes()
	AuthRoutes()
	AdminRoutes()
}
