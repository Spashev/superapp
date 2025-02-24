package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	"github.com/spashev/superapp/internal/repository"
	"github.com/spashev/superapp/internal/service"
)

func GetCategories(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		repo := repository.NewCategoryRepository(db)
		categoryService := service.NewCategoryService(repo)

		categories, err := categoryService.GetAllCategories()
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to fetch categories",
			})
		}

		return c.Status(fiber.StatusOK).JSON(categories)
	}
}
