package controller

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/pran2401/blog/database"
	"github.com/pran2401/blog/models"
	"gorm.io/gorm"
)

func Create(c *fiber.Ctx) error {
	var blog models.Blog
	if err := c.BodyParser(&blog); err != nil {
		fmt.Println(("unable to parse"))
	}

	if err := database.DB.Create(&blog).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{"message": "in valid blog"})
	}
	return c.JSON(fiber.Map{"message": "created", "blog": blog})
}

func Getpost(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		page = 1
	}
	limit := 10
	offset := (page - 1) * limit
	var total int64
	var getblog []models.Blog
	database.DB.Preload("user").Offset(offset).Limit(limit).Find(&getblog)
	database.DB.Model(&models.Blog{}).Count(&total)
	fmt.Println("in get post api")
	return c.JSON(fiber.Map{
		"data": getblog,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(total) / float64(limit)),
		},
	})
}

func Post(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		id = 0
	}

	var blogpost models.Blog
	database.DB.Where("id=?", id).Preload("user").First(&blogpost)
	return c.JSON(fiber.Map{"blog": blogpost})
}

func Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		id = 0
	}

	blog := models.Blog{
		Id: uint(id),
	}

	if err := c.BodyParser(&blog); err != nil {
		fmt.Println(("unable to parse"))
	}
	database.DB.Model(&blog).Updates(blog)
	return c.JSON(blog)
}

func Userposts(c *fiber.Ctx) error {
	cookie := c.Cookies(("auth"))
	id, err := strconv.Atoi(cookie)
	if err != nil {
		id = 0
	}
	var blog []models.Blog
	database.DB.Model(&blog).Where("userid=?", id).Preload("user").Find(&blog)
	return c.JSON(fiber.Map{"blogs": blog})
}

func Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		id = 0
	}

	blog := models.Blog{
		Id: uint(id),
	}

	query := database.DB.Delete(&blog)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		c.Status(400)
		return c.JSON(fiber.Map{"message": "post not found"})
	}

	return c.JSON(fiber.Map{"message": "deleted"})
}
