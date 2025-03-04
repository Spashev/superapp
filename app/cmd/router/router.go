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
	app := fiber.New(initializers.NewFiberConfig())

	app.Use(middleware.CorsHandler)
	app.Use(requestid.New())
	app.Use(initializers.NewLogger())
	app.Use(middleware.AuthMiddleware(db, tokenMaker))
	app.Use(initializers.NewSwagger())

	apiV1 := app.Group("/api/v1")
	{
		apiV1.Post("/users/create", handler.Register(db, tokenMaker))
		apiV1.Post("/users/token", handler.Login(db, tokenMaker))
		apiV1.Get("/users/me", handler.UserMe(db, tokenMaker))

		apiV1.Get("/user/favorite/products", handler.GetUserFavoriteProducts(db))

		apiV1.Get("/products", middleware.Paginate, handler.GetProductList(db))
		apiV1.Post("/products/:id/like", handler.LikeProductById(db))
		apiV1.Delete("/products/:id/like", handler.DislikeProductById(db))
		apiV1.Get("/products/:slug", handler.GetProductBySlug(db))

		apiV1.Get("/categories", handler.GetCategories(db))
		apiV1.Get("/conveniences", handler.GetConveniences(db))
		apiV1.Get("/types", handler.GetTypes(db))

		apiV1.Get("/countries", handler.GetCountry(db))
		apiV1.Get("/cities", handler.GetCity(db))
		apiV1.Get("/country/:id/cities", handler.GetCityByCountryId(db))
	}

	return app
}
