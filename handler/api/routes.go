package handler

import (
	"net/http"

	"github.com/go-chi/chi"
)

var r *chi.Mux

func RegisterRoutes(handler *Handler) *chi.Mux {
	r = chi.NewRouter()

	return r
}

func Start(addr string) error {
	return http.ListenAndServe(addr, r)
}
