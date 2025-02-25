package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	_ "github.com/spashev/superapp/internal/models"
	"github.com/spashev/superapp/internal/repository"
	"github.com/spashev/superapp/internal/service"
)

// GetCategories retrieves a list of all categories
// @Summary Get all categories
// @Description Fetches all available categories
// @Tags categories
// @Produce json
// @Success 200 {array} models.Category "List of categories"
// @Failure 500 {object} fiber.Map "Failed to fetch categories"
// @Router /categories [get]
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
