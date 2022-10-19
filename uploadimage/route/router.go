package route

import (
	"forPractice/projectGolang/BlogWebsite/uploadimage/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Post("/api/upload-image", controller.Upload)
	//app.Static("/api/uploadimgs", "./uploadimgs")

}
