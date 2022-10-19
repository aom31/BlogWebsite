package controller

import (
	"errors"
	"fmt"

	"forPractice/projectGolang/BlogWebsite/authorization/util"
	"forPractice/projectGolang/BlogWebsite/common/models"
	"forPractice/projectGolang/BlogWebsite/post/dao"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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

func AllPost(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 5
	offset := (page - 1) * limit
	var total int64
	var getblog []models.Blog
	dao.DB.Preload("User").Offset(offset).Limit(limit).Find(&getblog)
	dao.DB.Model(&models.Blog{}).Count(&total)
	return c.JSON(fiber.Map{
		"data": getblog,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})
}

func DetailPostId(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var blogpost models.Blog
	dao.DB.Where("id=?", id).Preload("User").First(&blogpost)
	return c.JSON(fiber.Map{
		"data": blogpost,
	})
}

func UpdatePostId(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	blog := models.Blog{
		ID: uint(id),
	}
	if err := c.BodyParser(&blog); err != nil {
		fmt.Println("Unable to parse body")
	}
	dao.DB.Model(&blog).Updates(blog)
	return c.JSON(fiber.Map{
		"message": "Updated post successfully",
	})
}

func UniquePost(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	id, _ := util.Parsejwt(cookie)
	var blog []models.Blog
	dao.DB.Model(&blog).Where("user_id = ?", id).Preload("User").Find(&blog)

	return c.JSON(blog)
}

func DeletePostId(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	blog := models.Blog{
		ID: uint(id),
	}
	deleteQuery := dao.DB.Delete(&blog)
	if errors.Is(deleteQuery.Error, gorm.ErrRecordNotFound) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Opps! ,record not found",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Deleted post  successfully!",
	})

}
