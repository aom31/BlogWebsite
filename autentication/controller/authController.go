package controller

import (
	"blogwebsite/autentication/dao"
	"fmt"
	"log"
	"regexp"
	"strings"

	"blogwebsite/common/models"

	"github.com/gofiber/fiber/v2"
)

//check format of email should be
func validateEmail(email string) bool {
	//format xxx@mail.com
	Re := regexp.MustCompile(`[a-z0-9.+-_&]+@[a-z]+\.[a-z]`)
	return Re.MatchString(email)
}

func Register(c *fiber.Ctx) error {
	// use map[data sent from page]type data in model all /use interface because have string and byte
	var data map[string]interface{}
	var userData models.User
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}

	//check if  password less than 8 characters
	//convert pass to string first because it is byte
	if len(data["password"].(string)) <= 8 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password must be  greater than 8 charactor",
		})
	}
	//check field email not be empty should have
	if !validateEmail(strings.TrimSpace(data["email"].(string))) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid Email Address",
		})
	}
	//check email already exist db
	dao.DB.Where("email= ?", strings.TrimSpace(data["email"].(string))).First(&userData)
	if userData.ID != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Email already exist",
		})
	}
	//record write to db
	user := models.User{
		FirstName: data["first_name"].(string),
		LastName:  data["last_name"].(string),
		Phone:     data["phone"].(string),
		Email:     strings.TrimSpace(data["email"].(string)),
	}
	user.SetPassword(data["password"].(string))
	err := dao.DB.Create(&user)
	if err != nil {
		log.Println(err)
	}
	c.Status(200)
	return c.JSON(fiber.Map{
		"user":    user,
		"message": "Account created successfullys",
	})
}
