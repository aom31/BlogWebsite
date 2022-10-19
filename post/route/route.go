package route

import (
	"forPractice/projectGolang/BlogWebsite/authorization/middleware"
	"forPractice/projectGolang/BlogWebsite/post/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Use(middleware.IsAuthenicate)

	app.Post("/api/blogweb/create-post", controller.CreatePost)
	app.Get("/api/blogweb/posts", controller.AllPost)
	app.Get("/api/blogweb/post/:id", controller.DetailPostId)
	app.Put("/api/blogweb/post/:id", controller.UpdatePostId)
	app.Get("/api/blogweb/post", controller.UniquePost)
	app.Delete("/api/blogweb/post/:id", controller.DeletePostId)
}
