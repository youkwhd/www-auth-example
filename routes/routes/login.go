package routes

import (
	"fmt"

	"www-auth-example/config"
	"www-auth-example/cookie"
	"www-auth-example/db"
	"www-auth-example/db/user"
	"www-auth-example/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Login(app *fiber.App, config config.Config) {
    app.Post("/login", middlewares.AuthRedirect, func(c *fiber.Ctx) error {
        fmt.Println("POST /login")
        user := user.User{}

        err := c.BodyParser(&user)
        if err != nil {
            return c.JSON(map[string]any{
                "success": false,
            })
        }

        foundUser, found := db.Data.Users.Get(user.Username)
        if !found {
            // user does not exist
            return c.JSON(map[string]any{
                "success": false,
            })
        }

        if foundUser.Password != user.Password {
            // password is not the same
            return c.JSON(map[string]any{
                "success": false,
            })
        }

        cookie := cookie.NewAuthCookie(config.Cookie.ExpiredAfter)
        db.Data.Sessions.Add(cookie.Value, &foundUser, cookie.Expires)

        c.Cookie(&cookie)

        return c.JSON(map[string]any{
            "success": true,
        })
    })
}
