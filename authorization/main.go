package main

import (
	"forPractice/projectGolang/BlogWebsite/authorization/dao"
	"forPractice/projectGolang/BlogWebsite/authorization/route"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	dao.ConnectDb()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env files")
	} else {
		log.Println("connect successfully")
	}
	port := os.Getenv("PORT")
	app := fiber.New()
	route.Setup(app)
	app.Listen(":" + port)
}
