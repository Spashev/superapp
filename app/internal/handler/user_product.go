package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	_ "github.com/spashev/superapp/internal/models"
	"github.com/spashev/superapp/internal/repository"
	"github.com/spashev/superapp/internal/service"
	"github.com/spashev/superapp/internal/util/token"
)

// GetUserFavoriteProducts retrieves a list of user favorite products
// @Summary Get a list of user favorite products
// @Description Returns a paginated list of user favorite products
// @Tags UserProducts
// @Accept json
// @Produce json
// @Param limit query int false "Limit" default(20)
// @Param offset query int false "Offset" default(0)
// @Success 200 {array} models.Product
// @Failure 401 {object} fiber.Map "Unauthorized"
// @Failure 500 {object} map[string]string "Server error"
// @Security BearerAuth
// @Router /user/favorite/products [get]
func GetUserFavoriteProducts(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		limit := c.QueryInt("limit", 20)
		offset := c.QueryInt("offset", 0)

		var user token.UserClaims
		if u, ok := c.Locals("user").(*token.UserClaims); ok && u != nil {
			user = *u
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		repo := repository.NewUserProductRepository(db)
		productService := service.NewUserProductService(repo)

		products, err := productService.GetUserFavoriteProducts(user.UserID, limit, offset)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to fetch user favorite products",
			})
		}
		return c.Status(fiber.StatusOK).JSON(products)
	}
}
