package middleware

import (
	"forPractice/projectGolang/BlogWebsite/authorization/util"

	"github.com/gofiber/fiber/v2"
)

func IsAuthenicate(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	if _, err := util.Parsejwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})

	}
	return c.Next() //pull next routes
}
