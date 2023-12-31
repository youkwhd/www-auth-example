package middlewares

import (
	"www-auth-example/cookie"
	"www-auth-example/db"
	"github.com/gofiber/fiber/v2"
)

func AuthRedirect(c *fiber.Ctx) error {
    clientCookie := c.Cookies(cookie.COOKIE_AUTH, cookie.COOKIE_AUTH_NONE)

    if clientCookie != cookie.COOKIE_AUTH_NONE {
        // if the client has a cookie but it's invalid
        // then just continue off
        _, err := db.Data.Sessions.Get(clientCookie)
        if err != nil {
            return c.Next()
        }

        return c.JSON(map[string]any{
            // TODO:
            // Implement command to redirect
            "success": true,
        })
    }

    return c.Next()
}

func AuthRequired(c *fiber.Ctx) error {
    clientCookie := c.Cookies(cookie.COOKIE_AUTH, cookie.COOKIE_AUTH_NONE)

    if clientCookie == cookie.COOKIE_AUTH_NONE {
        return c.JSON(map[string]any{
            "success": false,
        })
    }

    _, err := db.Data.Sessions.Get(clientCookie)

    if err != nil {
        return c.JSON(map[string]any{
            "success": false,
            // "message": err,
        })
    }

    return c.Next()
}
