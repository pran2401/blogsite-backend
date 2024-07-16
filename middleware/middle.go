package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func Middle(c *fiber.Ctx) error {
	cookie := c.Cookies("auth")
	return c.JSON(fiber.Map{"cookie": cookie})
}
