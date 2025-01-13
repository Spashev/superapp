package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	"superapp/internal/repository"
	"superapp/internal/service"
)

func GetProductList(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		page := c.QueryInt("page", 1)
		limit := c.QueryInt("limit", 10)

		repo := repository.NewProductRepository(db)
		productService := service.NewProductService(repo)

		products, err := productService.GetAllProducts(page, limit)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to fetch products",
			})
		}

		return c.Status(fiber.StatusOK).JSON(products)
	}
}

func GetProductBySlug(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		slug := c.Params("slug")
		if slug == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Slug is required",
			})
		}

		repo := repository.NewProductRepository(db)
		productService := service.NewProductService(repo)

		product, err := productService.GetProductBySlug(slug)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to fetch product",
			})
		}

		return c.Status(fiber.StatusOK).JSON(product)
	}
}
