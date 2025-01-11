package repository

import (
	"errors"

	"github.com/jmoiron/sqlx"

	"superapp/internal/models"
)

type AuthenticationRepository struct {
	db *sqlx.DB
}

func NewAuthenticationRepository(db *sqlx.DB) *AuthenticationRepository {
	return &AuthenticationRepository{db: db}
}

func (repo *AuthenticationRepository) ValidateUserCredentials(email, password string) (*models.User, error) {
	var user models.User
	query := `
		SELECT 
			id, 
			email,
			first_name, 
			password, 
			is_active 
		FROM users 
		WHERE email = $1
	`
	err := repo.db.Get(&user, query, email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if user.Password != password {
		return nil, errors.New("incorrect password")
	}

	return &user, nil
}
