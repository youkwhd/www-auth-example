package main

import (
	"fmt"

	"www-auth-example/cookie"
	"www-auth-example/middlewares/auth"

	"www-auth-example/db"
	"www-auth-example/db/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
    app := fiber.New()

    config := InitConfig()
    config.AddUsersToDatabase()

    cors := cors.New(cors.Config{
        AllowOrigins: config.GenerateAllowedOrigins(),
        AllowCredentials: true,
        AllowHeaders: "Origin, Content-Type, Accept",
    })

    app.Use(cors)

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

    app.Get("/user", middlewares.AuthRequired, func(c *fiber.Ctx) error {
        fmt.Println("GET /user")

        clientCookie := c.Cookies(cookie.COOKIE_AUTH, cookie.COOKIE_AUTH_NONE)
        session := db.Data.Sessions[clientCookie]

        return c.JSON(map[string]any{
            "success": true,
            "username": session.User.Username,
        })
    })

    app.Post("/login", middlewares.AuthRedirect, func(c *fiber.Ctx) error {
        fmt.Println("POST /login")
        user := user.User{}

        err := c.BodyParser(&user)
        if err != nil {
            return c.JSON(map[string]any{
                "success": false,
            })
        }

        foundUser, found := db.Data.Users[user.Username]
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
        tmp := db.Data.Users[user.Username]
        db.Data.Sessions.Add(cookie.Value, &tmp, cookie.Expires)
        c.Cookie(&cookie)

        return c.JSON(map[string]any{
            "success": true,
        })
    })

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

        _, found := db.Data.Users[newUser.Username]
        if found {
            // user already exist
            return c.JSON(map[string]any{
                "success": false,
            })
        }

        db.Data.Users.Add(newUser.Username, newUser.Password)
        databaseUser := db.Data.Users[newUser.Username]

        cookie := cookie.NewAuthCookie(config.Cookie.ExpiredAfter)
        db.Data.Sessions.Add(cookie.Value, &databaseUser, cookie.Expires)

        c.Cookie(&cookie)

        return c.JSON(map[string]any{
            "success": true,
        })
    })

    app.Listen(":3000")
}
