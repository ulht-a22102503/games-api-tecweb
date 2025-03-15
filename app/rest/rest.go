package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/martim-lusofona/games-api/db"
)

func PingDB() bool {
	return db.Ping()
}

func GetGames(c *gin.Context) {
	games, err := db.GetGames()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "server error",
		})
	}

	c.JSON(http.StatusOK, games)
}
