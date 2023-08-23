package cookie

import (
	"time"
	"github.com/google/uuid"
	"github.com/gofiber/fiber/v2"
)

const (
    COOKIE_AUTH = "AUTHSESSIONID" 
    COOKIE_AUTH_NONE = "NONE"
)

func NewAuthCookie() fiber.Cookie {
    cookie := fiber.Cookie{}
    cookie.Name = COOKIE_AUTH
    cookie.Value = uuid.NewString() 
    cookie.Path = "/"
    
    cookie.HTTPOnly = true

    // see: <https://stackoverflow.com/questions/46288437/set-cookies-for-cross-origin-requests>
    cookie.Secure = true
    cookie.SameSite = "None"

    cookie.Expires = time.Now().Add(10 * time.Second)
    return cookie
}
