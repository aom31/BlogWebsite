package controller

import (
	"fmt"
	"forPractice/projectGolang/BlogWebsite/common/models"
	"forPractice/projectGolang/BlogWebsite/createPost/dao"

	"github.com/gofiber/fiber/v2"
)

func CreatePost(c *fiber.Ctx) error {
	var blogpost models.Blog
	if err := c.BodyParser(&blogpost); err != nil {
		fmt.Println("Unable to parse body")
	}
	if err := dao.DB.Create(&blogpost).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid payload",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Congratulation!, Your post created",
	})
}
