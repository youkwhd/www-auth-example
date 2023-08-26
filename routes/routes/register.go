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

func Register(app *fiber.App, config config.Config) {
    app.Post("/register", middlewares.AuthRedirect, func(c *fiber.Ctx) error {
        fmt.Println("POST /register")
        newUser := user.User{}

        err := c.BodyParser(&newUser)
        if err != nil {
            // should have message
            // on why it failed
            return c.JSON(map[string]any{
                "success": false,
            })
        }

        _, found := db.Data.Users.Get(newUser.Username)
        if found {
            // user already exist
            return c.JSON(map[string]any{
                "success": false,
            })
        }

        db.Data.Users.Add(newUser.Username, newUser.Password)
        databaseUser, _ := db.Data.Users.Get(newUser.Username)

        cookie := cookie.NewAuthCookie(config.Cookie.ExpiredAfter)
        db.Data.Sessions.Add(cookie.Value, &databaseUser, cookie.Expires)

        c.Cookie(&cookie)

        return c.JSON(map[string]any{
            "success": true,
        })
    })
}
