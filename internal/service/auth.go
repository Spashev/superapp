package service

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"

	"superapp/internal/repository"
	schema "superapp/internal/schema/auth"
	"superapp/internal/util/token"
)

type AuthenticationService struct {
	repo       *repository.AuthenticationRepository
	tokenMaker *token.JWTMaker
}

func NewAuthenticationService(db *sqlx.DB, tokenMaker *token.JWTMaker) *AuthenticationService {
	repo := repository.NewAuthenticationRepository(db)
	return &AuthenticationService{repo: repo, tokenMaker: tokenMaker}
}

func (s *AuthenticationService) Login(r *http.Request) (string, *token.UserClaims, error) {
	var loginReq schema.AuthLoginReq
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		return "", nil, errors.New("invalid request body")
	}

	user, err := s.repo.ValidateUserCredentials(loginReq.Email, loginReq.Password)
	if err != nil {
		log.Println("Invalid login credentials", err)
		return "", nil, errors.New("invalid login credentials")
	}

	duration := time.Hour * 24
	tokenString, claims, err := s.tokenMaker.CreateToken(user.Id, user.Email, user.First_name, user.Is_active, duration)
	if err != nil {
		log.Println("Failed to create JWT token:", err)
		return "", nil, errors.New("failed to generate token")
	}

	return tokenString, claims, nil
}
