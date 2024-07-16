package controller

import (
	"fmt"
	"strconv"

	//"log"
	"strings"

	//"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/pran2401/blog/database"
	"github.com/pran2401/blog/models"
)

func Register(c *fiber.Ctx) error {
	var data map[string]interface{}
	var userdata models.User
	if err := c.BodyParser(&data); err != nil {
		fmt.Println(("unable to parse"))
	}
	if len(data["password"].(string)) <= 5 {
		c.Status(400)
		return c.JSON(fiber.Map{"message": "password must be more than 5 letters"})
	}

	database.DB.Where("email =?", strings.TrimSpace(data["email"].(string))).First(&userdata)
	if userdata.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{"message": "already exits"})
	}
	user := models.User{
		Firstname: data["first_name"].(string),
		Lastname:  data["last_name"].(string),
		Email:     data["email"].(string),
		Password:  data["password"].(string),
	}

	database.DB.Create(&user)

	c.Status(200)
	return c.JSON(fiber.Map{"user": user, "message": "acc created"})

}

func Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		fmt.Println(("unable to parse"))
	}
	var user models.User
	var user1 models.User

	cookie := fiber.Cookie{
		Name:  "auth",
		Value: "0",
	}
	database.DB.Where("email =?", data["email"]).First(&user)

	if user == user1 {
		c.Status(404)
		c.Cookie(&cookie)
		return c.JSON(fiber.Map{"message": "user not found", "user": user})
	}

	if user.Password != data["password"] {
		c.Status(400)
		c.Cookie(&cookie)
		return c.JSON(fiber.Map{"message": "incorrect password"})
	}
	if user.Password == data["password"] {
		cookie.Value = strconv.FormatUint(uint64(user.Id), 10)
	}

	c.Cookie(&cookie)
	c.Status(200)
	return c.JSON(fiber.Map{"message": "welcome" + user.Firstname, "user": user})

}
