package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pran2401/blog/controller"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	app.Post("/api/create", controller.Create)
	app.Get("/api/view", controller.Getpost)
	app.Get("/api/view/:id", controller.Post)
	app.Put("api/update/:id", controller.Update)
	app.Get("/api/userposts", controller.Userposts)
	app.Delete("/api/delete/:id", controller.Delete)
	app.Post("/api/upload-image", controller.Image)
	app.Static("api/images", "./images")
}
