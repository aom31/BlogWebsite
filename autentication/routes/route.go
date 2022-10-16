package routes

import (
	"blogwebsite/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/blogweb/register", controller.Register)
}
