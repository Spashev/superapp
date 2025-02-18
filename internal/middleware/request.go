package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Logger(c *fiber.Ctx) error {
	start := time.Now()

	err := c.Next()

	duration := time.Since(start)
	status := c.Response().StatusCode()
	method := c.Method()
	path := c.Path()
	reqiestID := c.Request().Header.Peek("REQUEST-ID")

	fmt.Printf("[INFO] %s - %s %s %d %s %s\033[0m\n", method, path, time.Now().Format("2006-01-02 15:04:05"), status, reqiestID, duration)

	return err
}
