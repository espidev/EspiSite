package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

var (
	router *gin.Engine
	config Config
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
	router.GET("/", func (c *gin.Context) {
		c.String(http.StatusOK, "HI")
	})
}
