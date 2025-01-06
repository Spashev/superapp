package middleware

import (
	"context"
	"net/http"
	"strconv"
)

func Paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pageStr := r.URL.Query().Get("page")
		limitStr := r.URL.Query().Get("limit")

		page, err := strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			page = 1
		}

		limit, err := strconv.Atoi(limitStr)
		if err != nil || limit < 1 {
			limit = 20
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "page", page)
		ctx = context.WithValue(ctx, "limit", limit)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
