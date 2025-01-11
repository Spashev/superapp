package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type UserClaims struct {
	ID        int64  `json:"id"`
	FirstName string `json:"frist_name"`
	Email     string `json:"email"`
	IsActive  bool   `json:"is_active"`
	jwt.RegisteredClaims
}

func NewUserClaims(id int64, email string, firstName string, isActive bool, duration time.Duration) (*UserClaims, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("error generating token ID: %w", err)
	}

	return &UserClaims{
		Email:     email,
		ID:        id,
		FirstName: firstName,
		IsActive:  isActive,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        tokenID.String(),
			Subject:   email,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}, nil
}
