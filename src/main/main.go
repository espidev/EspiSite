package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

var (
	router *gin.Engine
	config Config
)

const (
	RootRepoFolder = "."
)

func main() {
	log.Printf("Starting EspiSite...\n")
	setupConfig()

	router = gin.Default()
	setupRoutes()

	err := router.Run(":" + strconv.Itoa(config.Port))
	if err != nil {
		log.Fatal(err)
	}
}

func setupRoutes() {
	router.LoadHTMLFiles("/src/html/*")
	router.Static("/", RootRepoFolder + "/src/html/")
}
g