package controller

import (
	"math/rand"

	"github.com/gofiber/fiber/v2"
)

var letter = []rune("abcdefghijklmnopgrstuvwxyz")

func randletter(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func Image(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	files := form.File["image"]
	filename := ""

	for _, file := range files {
		filename = randletter(5) + "-" + file.Filename
		if err := c.SaveFile(file, "./images/"+filename); err != nil {
			return err
		}
	}

	return c.JSON(fiber.Map{"url": "http://localhost:3000/api/images/" + filename})
}
