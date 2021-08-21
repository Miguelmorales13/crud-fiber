package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
)

func LoggerMiddleware() fiber.Handler {
	return logger.New(logger.Config{
		Format:     "${red}[${time}]-${blue}[${locals:requestid}] [${ip}|${latency}] [${method}|${url}] [${status}|${body}]  \n",
		TimeFormat: "2006-01-02 15:04:05.00000",
		Output:     os.Stderr,
	})
}
