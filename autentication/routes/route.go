package routes

import (
	"forPractice/projectGolang/BlogWebsite/autentication/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/blogweb/register", controller.Register)
}
