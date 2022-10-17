package route

import (
	"forPractice/projectGolang/BlogWebsite/authorization/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	//app.Use(middleware.IsAuthenicate)
	app.Post("/api/blogweb/login", controller.Login)
}
