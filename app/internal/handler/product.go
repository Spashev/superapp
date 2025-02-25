package handler

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	_ "github.com/spashev/superapp/internal/models"
	"github.com/spashev/superapp/internal/repository"
	"github.com/spashev/superapp/internal/service"
	"github.com/spashev/superapp/internal/util/token"
)

// GetProductList retrieves a list of products
// @Summary Get a list of products
// @Description Returns a paginated list of products
// @Tags Products
// @Accept json
// @Produce json
// @Param limit query int false "Limit" default(20)
// @Param offset query int false "Offset" default(0)
// @Success 200 {array} models.Product
// @Failure 500 {object} map[string]string "Server error"
// @Router /products [get]
func GetProductList(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		limit := c.QueryInt("limit", 20)
		offset := c.QueryInt("offset", 0)

		var user token.UserClaims
		if u, ok := c.Locals("user").(*token.UserClaims); ok && u != nil {
			user = *u
		}

		repo := repository.NewProductRepository(db)
		productService := service.NewProductService(repo)

		products, err := productService.GetAllProducts(user.UserID, limit, offset)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to fetch products",
			})
		}

		return c.Status(fiber.StatusOK).JSON(products)
	}
}

// GetProductBySlug retrieves product information by slug
// @Summary Get a product by slug
// @Description Returns product details
// @Tags Products
// @Accept json
// @Produce json
// @Param slug path string true "Product slug"
// @Success 200 {object} models.Product
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Server error"
// @Router /products/{slug} [get]
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

// LikeProductById increases the like count of a product
// @Summary Like a product by slug
// @Description Increments the product's like count
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]string "Product likeed successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Server error"
// @Router /products/{id}/like [post]
func LikeProductById(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		params := c.Params("id")
		if params == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Id is required",
			})
		}

		id, err := strconv.Atoi(params)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Id must be an integer",
			})
		}

		var user token.UserClaims
		if u, ok := c.Locals("user").(*token.UserClaims); ok && u != nil {
			user = *u
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		repo := repository.NewProductRepository(db)
		productService := service.NewProductService(repo)

		if err := productService.LikeProductById(user.UserID, id); err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to like product",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Product liked successfully",
		})
	}
}

// DislikeProductById decreases the like count of a product
// @Summary Dislike a product by id
// @Description Increments the product's like count
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]string "Product disliked successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Server error"
// @Router /products/{id}/like [delete]
func DislikeProductById(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		params := c.Params("id")
		if params == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Id is required",
			})
		}

		id, err := strconv.Atoi(params)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Id must be an integer",
			})
		}

		var user token.UserClaims
		if u, ok := c.Locals("user").(*token.UserClaims); ok && u != nil {
			user = *u
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		repo := repository.NewProductRepository(db)
		productService := service.NewProductService(repo)

		if err := productService.DislikeProductById(user.UserID, id); err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to dislike product",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Product disliked successfully",
		})
	}
}
