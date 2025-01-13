package middleware

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Paginate(c *fiber.Ctx) error {
	pageStr := c.Query("page")
	limitStr := c.Query("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 20
	}

	c.Locals("page", page)
	c.Locals("limit", limit)

	return c.Next()
}
