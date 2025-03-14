package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/martim-lusofona/games-api/db"
	"github.com/martim-lusofona/games-api/rest"
)

func main() {
	log.Println("Starting API...")
	db.Connect()
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})

	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ping": "pong",
			"db":   rest.PingDB(),
		})

	})

	router.Run(":50007")
}
