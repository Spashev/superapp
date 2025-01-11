package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"

	"superapp/internal/handler"
	md "superapp/internal/middleware"
)

func RegisterRoutes(db *sqlx.DB) http.Handler {
	r := chi.NewRouter()

	r.Use(md.CorsHandler())
	r.Use(md.GeneralMiddleware())

	r.Route("/api/v1", func(r chi.Router) {
		r.With(md.Paginate).Get("/products/get", handler.GetProductList(db))
		r.Get("/products/{slug}", handler.GetProductBySlug(db))
		r.Get("/categories", handler.GetCategories(db))
	})

	return r
}
