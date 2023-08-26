package routes

import (
	"www-auth-example/config"
	"www-auth-example/routes/routes"

	"github.com/gofiber/fiber/v2"
)

func RegisterAllRoutes(app *fiber.App, config config.Config) {
	routes.Register(app, config)
	routes.Login(app, config)
	routes.Logout(app, config)
	routes.User(app, config)
}
