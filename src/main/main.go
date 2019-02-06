package main

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
)

const (
	RootRepoFolder = "."
)

func main() {
	log.Printf("Starting EspiSite...\n")
	setupConfig()

	router = gin.Default()
	setupRoutes()

	srv := &http.Server {
		Addr: ":" + strconv.Itoa(config.Port),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make (chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<- quit
	log.Println("Shutting down EspiSite...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}
	log.Println("EspiSite has stopped.")
}

func setupRoutes() {
	router.LoadHTMLFiles(RootRepoFolder + "/src/html/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	router.GET("/test", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{})
	})
}
