package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
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
	})

	app.Use(middleware.CorsHandler)
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "[${pid} ${locals:requestid}] ${ip}:${port} ${status} - ${method} ${path}\n",
	}))
	app.Use(middleware.AuthMiddleware(db, tokenMaker))

	apiV1 := app.Group("/api/v1")
	{
		apiV1.Post("/register", handler.Register(db, tokenMaker))
		apiV1.Post("/login", handler.Login(db, tokenMaker))
		apiV1.Get("/users/me", handler.UserMe(db, tokenMaker))

		apiV1.Get("/products", middleware.Paginate, handler.GetProductList(db))
		apiV1.Get("/products/:slug", handler.GetProductBySlug(db))
		apiV1.Get("/products/:slug/like", handler.LikeProductBySlug(db))
		apiV1.Get("/categories", handler.GetCategories(db))
	}

	return app
}
