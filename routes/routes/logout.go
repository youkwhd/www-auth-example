package routes

import (
	"fmt"

	"www-auth-example/config"
	"www-auth-example/cookie"
	"www-auth-example/db"
	"www-auth-example/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Logout(app *fiber.App, config config.Config) {
    app.Get("/logout", middlewares.AuthRequired, func(c *fiber.Ctx) error {
        fmt.Println("GET /logout")

        clientCookie := c.Cookies(cookie.COOKIE_AUTH, cookie.COOKIE_AUTH_NONE)
        delete(db.Data.Sessions, clientCookie)

        cookie := cookie.NewAuthCookie(0)
        c.Cookie(&cookie)

        return c.JSON(map[string]any{
            "success": true,
        })
    })
}
