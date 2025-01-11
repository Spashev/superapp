package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"

	"superapp/internal/service"
	"superapp/internal/util/token"
)

func Login(db *sqlx.DB, tokenMaker *token.JWTMaker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authService := service.NewAuthenticationService(db, tokenMaker)

		jwtToken, claims, err := authService.Login(r)
		if err != nil {
			http.Error(w, "Failed to login", http.StatusUnauthorized)
			log.Println(err)
			return
		}

		fmt.Println(jwtToken, claims)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(jwtToken); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			log.Println(err)
		}
	}
}
