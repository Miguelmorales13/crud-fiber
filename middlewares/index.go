package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func Config(e *fiber.App) {
	e.Use(LoggerMiddleware())
}
