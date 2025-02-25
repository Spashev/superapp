package initializers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewLogger() fiber.Handler {
	return logger.New(logger.Config{
		Format: "[${pid} ${locals:requestid}] ${ip}:${port} ${status} - ${method} ${path}\n",
	})
}
