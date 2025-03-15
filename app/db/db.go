package db

import (
	"fmt"

	"github.com/martim-lusofona/games-api/models"
)

func Ping() bool {
	dbConnection := GetDbConnection()
	if dbConnection == nil {
		return false
	}

	db, err := dbConnection.DB()
	if err != nil {
		return false
	}

	err = db.Ping()
	return err == nil
}

func GetGames() ([]models.Game, error) {
	dbConnection := GetDbConnection()
	if dbConnection == nil {
		return []models.Game{}, fmt.Errorf("connection error")
	}

	var games []models.Game
	err := dbConnection.Find(&games).Error
	return games, err
}
