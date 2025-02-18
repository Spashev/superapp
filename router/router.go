package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	"superapp/internal/handler"
	"superapp/internal/middleware"
	"superapp/internal/util/token"
)

func RegisterRoutes(db *sqlx.DB, tokenMaker *token.JWTMaker) *fiber.App {
	app := fiber.New(fiber.Config{
		ServerHeader:  "Bookit",
		AppName:       "Bookit App v0.1-beta",
		CaseSensitive: true,
		Prefork:       true, //TODO: need check it on server
	})

	app.Use(middleware.CorsHandler)
	app.Use(middleware.RequestID)
	app.Use(middleware.Logger)
	app.Use(middleware.AuthMiddleware(db, tokenMaker))

	apiV1 := app.Group("/api/v1")
	{
		apiV1.Post("/register", handler.Register(db, tokenMaker))
		apiV1.Post("/login", handler.Login(db, tokenMaker))
		apiV1.Get("/users/me", handler.UserMe(db, tokenMaker))

		apiV1.Get("/products", middleware.Paginate, handler.GetProductList(db))
		apiV1.Get("/products/:slug", handler.GetProductBySlug(db))
		apiV1.Get("/categories", handler.GetCategories(db))
	}

	return app
}
