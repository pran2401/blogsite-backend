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
		fmt.Println("unable to parse")
		return c.Status(400).JSON(fiber.Map{"message": "Unable to parse request body"})
	}

	password, ok := data["password"].(string)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid password format"})
	}

	if len(password) <= 5 {
		c.Status(400)
		return c.JSON(fiber.Map{"message": "password must be more than 5 letters"})
	}

	email, ok := data["email"].(string)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid email format"})
	}

	database.DB.Where("email =?", strings.TrimSpace(email)).First(&userdata)
	if userdata.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{"message": "already exits"})
	}

	firstName, ok := data["firstname"].(string)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid first name format"})
	}

	lastName, ok := data["lastname"].(string)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid last name format"})
	}

	user := models.User{
		Firstname: firstName,
		Lastname:  lastName,
		Email:     email,
		Password:  password,
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
