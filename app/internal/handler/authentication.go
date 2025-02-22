package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	"superapp/internal/repository"
	schema "superapp/internal/schema/auth"
	"superapp/internal/service"
	"superapp/internal/util/token"
)

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
