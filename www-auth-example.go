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

    cors := cors.New(cors.Config{
        AllowOrigins: "http://localhost:8000, http://127.0.0.1:8000, http://0.0.0.0:8000",
        AllowCredentials: true,
        AllowHeaders: "Origin, Content-Type, Accept",
    })

    app.Use(cors)

    app.Get("/user", middlewares.AuthRequired, func(c *fiber.Ctx) error {
        fmt.Println("GET /user")

        clientCookie := c.Cookies(cookie.COOKIE_AUTH, cookie.COOKIE_AUTH_NONE)
        session := db.Data.Sessions[clientCookie]

        return c.JSON(map[string]any{
            "success": true,
            "username": session.User.Username,
        })
    })

    app.Post("/login", func(c *fiber.Ctx) error {
        fmt.Println("POST /login")

        user := user.User{}

        err := c.BodyParser(&user)
        if err != nil {
            fmt.Println(err)

            // should have message
            // on why it failed
            return c.JSON(map[string]any{
                "success": false,
            })
        }

        foundUser, found := db.Data.Users[user.Username]
        if !found {
            fmt.Println(found)
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

        cookie := cookie.NewAuthCookie()
        tmp := db.Data.Users[user.Username]
        db.Data.Sessions.Add(cookie.Value, &tmp, cookie.Expires)
        c.Cookie(&cookie)

        return c.JSON(map[string]any{
            "success": true,
        })
    })

    app.Post("/register", func(c *fiber.Ctx) error {
        fmt.Println("POST /register")
        fromclient := c.Cookies(cookie.COOKIE_AUTH, cookie.COOKIE_AUTH_NONE)
        fmt.Println(fromclient)

        if fromclient != cookie.COOKIE_AUTH_NONE {
            // already logged in
            return c.JSON(map[string]any{
                "success": true,
            })
        }

        user := user.User{}

        err := c.BodyParser(&user)
        if err != nil {
            fmt.Println(err)

            // should have message
            // on why it failed
            return c.JSON(map[string]any{
                "success": false,
            })
        }

        _, found := db.Data.Users[user.Username]
        if found {
            // user already exist
            return c.JSON(map[string]any{
                "success": false,
            })
        }

        db.Data.Users.Add(user.Username, user.Password)

        cookie := cookie.NewAuthCookie()
        tmp := db.Data.Users[user.Username]
        db.Data.Sessions.Add(cookie.Value, &tmp, cookie.Expires)
        c.Cookie(&cookie)

        return c.JSON(map[string]any{
            "success": true,
        })
    })

    app.Listen(":3000")
}
