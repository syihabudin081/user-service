package routes

import (
	"github.com/gofiber/fiber/v2"
	"user-service/controllers"
)

func RouteInit(app *fiber.App, userController *controllers.UserController) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/api/users", userController.CreateUser)
	app.Get("/api/users", userController.GetAllUsers)
	app.Get("/api/users/:id", userController.GetUserById)
	app.Put("/api/users/:id", userController.UpdateUser)
	app.Delete("/api/users/:id", userController.DeleteUser)
}
