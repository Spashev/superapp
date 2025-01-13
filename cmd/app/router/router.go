package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	"superapp/internal/handler"
	"superapp/internal/middleware"
	"superapp/internal/util/token"
)

func RegisterRoutes(db *sqlx.DB, tokenMaker *token.JWTMaker) *fiber.App {
	app := fiber.New()

	app.Use(middleware.CorsHandler)
	app.Use(middleware.Logger)

	apiV1 := app.Group("/api/v1")
	{
		apiV1.Post("/register", middleware.Paginate, handler.Register(db, tokenMaker))
		apiV1.Post("/login", handler.Login(db, tokenMaker))
		apiV1.Get("/users/me", handler.UserMe(db, tokenMaker))

		apiV1.Get("/products/get", handler.GetProductList(db))
		apiV1.Get("/products/:slug", handler.GetProductBySlug(db))
		apiV1.Get("/categories", handler.GetCategories(db))
	}

	return app
}
