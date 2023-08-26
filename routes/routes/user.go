package routes

import (
	"fmt"

	"www-auth-example/config"
	"www-auth-example/cookie"
	"www-auth-example/db"
	"www-auth-example/middlewares"

	"github.com/gofiber/fiber/v2"
)

func User(app *fiber.App, config config.Config) {
    app.Get("/user", middlewares.AuthRequired, func(c *fiber.Ctx) error {
        fmt.Println("GET /user")

        clientCookie := c.Cookies(cookie.COOKIE_AUTH, cookie.COOKIE_AUTH_NONE)
        session, err := db.Data.Sessions.Get(clientCookie)

        if err != nil {
            return c.JSON(map[string]any{
                "success": false,
            })
        }

        return c.JSON(map[string]any{
            "success": true,
            "username": session.User.Username,
        })
    })
}
