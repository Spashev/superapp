package router

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"superapp/handler"
)

func RegisterRoutes(db *sql.DB) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/products", handler.GetProductList(db))
	r.Get("/products/{slug}", handler.GetProductBySlug(db))

	return r
}
