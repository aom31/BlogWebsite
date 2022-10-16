package routes

import (
	"blogwebsite/autentication/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/blogweb/register", controller.Register)
}
