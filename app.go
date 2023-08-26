package main

import (
	"www-auth-example/config"
	"www-auth-example/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
    app := fiber.New()
    config := config.InitConfig()

    config.AddUsersToDatabase()

    cors := cors.New(cors.Config{
        AllowOrigins: config.GenerateAllowedOrigins(),
        AllowCredentials: true,
        AllowHeaders: "Origin, Content-Type, Accept",
    })

    app.Use(cors)

	routes.RegisterAllRoutes(app, config)

    app.Listen(":3000")
}
