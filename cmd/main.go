package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"user-service/config"
	"user-service/config/migrations"
	"user-service/controllers"
	"user-service/repositories"
	"user-service/routes"
	"user-service/services"
)

func main() {
	app := fiber.New()
	config.DatabaseInit()
	migrations.Migrate()
	app.Use(cors.New(
		cors.Config{
			AllowOrigins: "*",
			AllowHeaders: "Origin, Content-Type, Accept",
		}))
	userRepo := repositories.NewUserRepository(config.DB)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	routes.RouteInit(app, userController)

	app.Listen(":3000")
}
