package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	_ "github.com/spashev/superapp/internal/models"
	"github.com/spashev/superapp/internal/repository"
	schema "github.com/spashev/superapp/internal/schema/auth"
	"github.com/spashev/superapp/internal/service"
	"github.com/spashev/superapp/internal/util/token"
)

// Login handles user authentication
// @Summary User login
// @Description Authenticate user and return access tokens
// @Tags auth
// @Accept json
// @Produce json
// @Param request body schema.AuthLoginReq true "User credentials"
// @Success 200 {object} schema.AuthLoginRes "Successfully authenticated"
// @Failure 400 {object} fiber.Map "Invalid request body or credentials"
// @Router /auth/login [post]
func Login(db *sqlx.DB, tokenMaker *token.JWTMaker) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var loginReq schema.AuthLoginReq

		if err := c.BodyParser(&loginReq); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid request body",
			})
		}

		repo := repository.NewAuthenticationRepository(db)
		authService := service.NewAuthenticationService(repo, tokenMaker)

		tokens, err := authService.Authenticate(&loginReq)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid credentials",
			})
		}

		return c.Status(fiber.StatusOK).JSON(tokens)
	}
}

// UserMe retrieves authenticated user information
// @Summary Get user details
// @Description Returns user details for the authenticated user
// @Tags auth
// @Produce json
// @Success 200 {object} models.User "User details"
// @Failure 401 {object} fiber.Map "Unauthorized"
// @Security BearerAuth
// @Router /auth/me [get]
func UserMe(db *sqlx.DB, tokenMaker *token.JWTMaker) fiber.Handler {
	return func(c *fiber.Ctx) error {
		repo := repository.NewAuthenticationRepository(db)
		authService := service.NewAuthenticationService(repo, tokenMaker)

		user, err := authService.UserMe(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(user)
	}
}

// Register handles user registration
// @Summary User registration
// @Description Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param request body schema.RegisterReq true "User registration data"
// @Success 201 {string} string "Registration successful"
// @Failure 400 {object} fiber.Map "Invalid request body or registration error"
// @Router /auth/register [post]
func Register(db *sqlx.DB, tokenMaker *token.JWTMaker) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var registerReq schema.RegisterReq
		if err := c.BodyParser(&registerReq); err != nil {
			fmt.Println("Error decoding JSON:", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		repo := repository.NewAuthenticationRepository(db)
		authService := service.NewAuthenticationService(repo, tokenMaker)

		if err := authService.Register(&registerReq); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).SendString("Registration successful")
	}
}
