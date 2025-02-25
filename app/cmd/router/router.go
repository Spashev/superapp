package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/jmoiron/sqlx"

	"github.com/spashev/superapp/initializers"
	"github.com/spashev/superapp/internal/handler"
	"github.com/spashev/superapp/internal/middleware"
	"github.com/spashev/superapp/internal/util/token"
)

func RegisterRoutes(db *sqlx.DB, tokenMaker *token.JWTMaker) *fiber.App {
	app := fiber.New(fiber.Config{
		ServerHeader:  "Bookit",
		AppName:       "Bookit App v0.1-beta",
		CaseSensitive: true,
	})

	app.Use(middleware.CorsHandler)
	app.Use(requestid.New())
	app.Use(initializers.NewLogger())
	app.Use(middleware.AuthMiddleware(db, tokenMaker))
	app.Use(initializers.NewSwagger())

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
