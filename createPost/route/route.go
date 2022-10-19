package route

import (
	"forPractice/projectGolang/BlogWebsite/authorization/middleware"
	"forPractice/projectGolang/BlogWebsite/createPost/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Use(middleware.IsAuthenicate)
	app.Post("/api/blogweb/create-post", controller.CreatePost)
	app.Get("/api/blogweb/posts", controller.AllPost)
}
