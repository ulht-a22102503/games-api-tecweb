package db

import (
	"fmt"
	"log"
	"os"

	"github.com/martim-lusofona/games-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	_DB_CONNECTION *gorm.DB
)

func Connect() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return
	}
	_DB_CONNECTION = db
}

func Migrate() {
	if _DB_CONNECTION == nil {
		log.Println("Migrate NOK")
		return
	}
	err := _DB_CONNECTION.AutoMigrate(
		&models.Game{},
	)

	if err != nil {
		log.Println(err)
	} else {
		log.Println("Migrate OK")
	}
}

func Populate() {
	if _DB_CONNECTION == nil {
		log.Println("Populate NOK")
		return
	}
	var game models.Game
	err := _DB_CONNECTION.Where(models.Game{
		Title: "",
	}).FirstOrCreate(&game).Error

	if err != nil {
		log.Println(err)
	} else {
		log.Println("Populate OK")
	}
}

func GetDbConnection() *gorm.DB {
	if _DB_CONNECTION == nil {
		Connect()
	}

	return _DB_CONNECTION
}
