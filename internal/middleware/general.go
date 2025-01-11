package middleware

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

func GeneralMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		next = middleware.Logger(next)
		next = middleware.Recoverer(next)
		next = middleware.StripSlashes(next)
		return next
	}
}
