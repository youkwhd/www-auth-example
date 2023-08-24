package cookie

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

const (
    COOKIE_AUTH = "AUTHSESSIONID" 
    COOKIE_AUTH_NONE = "NONE"
)

func NewAuthCookie(duration int) fiber.Cookie {
    cookie := fiber.Cookie{}
    cookie.Name = COOKIE_AUTH
    cookie.Value = uuid.NewString() 
    cookie.Path = "/"
    
    cookie.HTTPOnly = true

    // see: <https://stackoverflow.com/questions/46288437/set-cookies-for-cross-origin-requests>
    cookie.Secure = true
    cookie.SameSite = "None"

    expiredAfter := time.Duration(duration) * time.Minute
    cookie.Expires = time.Now().Add(expiredAfter)
    return cookie
}
