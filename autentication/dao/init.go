package dao

import (
	"blogwebsite/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

//connect to db mysql pass: 3143
func ConnectDb() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error lod .env file")
	}
	dsn := os.Getenv("DSN")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could  connect to the database")
	} else {
		log.Println("connect db successfully")
	}
	DB = database
	//auto create table in db from model
	database.AutoMigrate(
		&models.User{},
	)
}
