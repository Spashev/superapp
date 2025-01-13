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

	var colorStatus string
	if status >= 400 {
		colorStatus = "\033[31m"
	} else {
		colorStatus = "\033[32m"
	}

	fmt.Printf("%s[INFO] %s - %s %s %d %s\033[0m\n", colorStatus, method, path, time.Now().Format("2006-01-02 15:04:05"), status, duration)

	return err
}
