package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func RequestID(c *fiber.Ctx) error {
	uid := uuid.NewString()
	c.Request().Header.Add("REQUEST-ID", uid)

	return c.Next()
}
