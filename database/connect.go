package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pran2401/blog/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env file")
	}

	dsn := os.Getenv("DSN")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("could not connect to db")
	} else {
		log.Println("connected")
	}
	DB = database

	database.AutoMigrate(&models.User{}, &models.Blog{})

}
