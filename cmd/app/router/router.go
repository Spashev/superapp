package router

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"superapp/internal/handler"
	md "superapp/internal/middleware"
)

func RegisterRoutes(db *sql.DB) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.StripSlashes)

	r.With(md.Paginate).Get("/api/v1/products/get", handler.GetProductList(db))
	r.Get("/api/v1/products/{slug}", handler.GetProductBySlug(db))

	return r
}
