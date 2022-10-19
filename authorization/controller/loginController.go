package controller

import (
	"fmt"
	"forPractice/projectGolang/BlogWebsite/common/models"
	"strconv"
	"time"

	"forPractice/projectGolang/BlogWebsite/authorization/dao"
	"forPractice/projectGolang/BlogWebsite/authorization/util"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type Claims struct {
	jwt.StandardClaims
}

func Login(c *fiber.Ctx) error {
	var data map[string]string
	var user models.User
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}
	//check email login have exist in db if not tell user to register first
	dao.DB.Where("email= ?", data["email"]).First(&user) //find in db exist data first
	//email not register before
	if user.ID == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"meassage": "Email address doesn't exist, kindly register an account",
		})
	}
	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}
	token, err := util.GenerateJwt(strconv.Itoa(int(user.ID)))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "You have successfully login",
		"user":    user,
	})

}
