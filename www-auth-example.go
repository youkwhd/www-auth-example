package main

import (
	"fmt"
	"time"

	"www-auth-example/db"
	"www-auth-example/db/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/google/uuid"
)

const (
    COOKIE_AUTH = "AUTHSESSIONID" 
    COOKIE_AUTH_NONE = "NONE"
)

func NewAuthCookie() *fiber.Cookie {
    cookie := fiber.Cookie{}
    cookie.Name = COOKIE_AUTH
    cookie.Value = uuid.NewString() 
    cookie.Path = "/"
    
    cookie.HTTPOnly = true

    // see: <https://stackoverflow.com/questions/46288437/set-cookies-for-cross-origin-requests>
    cookie.Secure = true
    cookie.SameSite = "None"

    cookie.Expires = time.Now().Add(30 * time.Second)
    return &cookie
}

func main() {
    app := fiber.New()
    database := db.Init()

    cors := cors.New(cors.Config{
        AllowOrigins: "http://localhost:8000, http://127.0.0.1:8000, http://0.0.0.0:8000",
        AllowCredentials: true,
        AllowHeaders: "Origin, Content-Type, Accept",
    })

    app.Use(cors)

    app.Get("/user", func(c *fiber.Ctx) error {
        fmt.Println("GET /user")
        fromclient := c.Cookies(COOKIE_AUTH, COOKIE_AUTH_NONE)

        // get db session
        if fromclient == COOKIE_AUTH_NONE {
            return c.JSON(map[string]any{
                "success": false,
            })
        }

        session, found := database.Sessions[fromclient]
        fmt.Println(session)

        if !found {
            return c.JSON(map[string]any{
                "success": false,
            })
        }

        return c.JSON(map[string]any{
            "success": true,
            "username": session.User.Username,
        })
    })

    app.Post("/login", func(c *fiber.Ctx) error {
        fmt.Println("POST /login")
        fromclient := c.Cookies(COOKIE_AUTH, COOKIE_AUTH_NONE)

        if fromclient != COOKIE_AUTH_NONE {
            // already logged in
            return c.JSON(map[string]any{
                "success": false,
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

        foundUser, found := database.Users[user.Username]
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

        cookie := NewAuthCookie()
        tmp := database.Users[user.Username]
        database.Sessions.Add(cookie.Value, &tmp)
        c.Cookie(cookie)

        return c.JSON(map[string]any{
            "success": true,
        })
    })

    app.Post("/register", func(c *fiber.Ctx) error {
        fmt.Println("POST /register")
        fromclient := c.Cookies(COOKIE_AUTH, COOKIE_AUTH_NONE)
        fmt.Println(fromclient)

        if fromclient != COOKIE_AUTH_NONE {
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

        _, found := database.Users[user.Username]
        if found {
            // user already exist
            return c.JSON(map[string]any{
                "success": false,
            })
        }

        database.Users.Add(user.Username, user.Password)

        cookie := NewAuthCookie()
        tmp := database.Users[user.Username]
        database.Sessions.Add(cookie.Value, &tmp)
        c.Cookie(cookie)

        return c.JSON(map[string]any{
            "success": true,
        })
    })

    app.Listen(":3000")
}
