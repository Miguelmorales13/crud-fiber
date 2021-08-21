package middlewares

import (
	"crud/utils/jwt"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func JwtMiddleware(c *fiber.Ctx) error {
	h := c.Get("Authorization")
	if h == "" {
		return fiber.ErrUnauthorized
	}
	chunks := strings.Split(h, " ")
	if len(chunks) < 2 {
		return fiber.ErrUnauthorized
	}
	user, err := jwt.Verify(chunks[1])
	if err != nil {
		return fiber.ErrUnauthorized
	}
	c.Locals("USER", user.ID)
	return c.Next()
}
