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

func (repo *AuthenticationRepository) GetUserByEmail(email string) (*models.CreateUser, error) {
	var user models.CreateUser
	query := `
		SELECT 
			id, 
			email,
			first_name, 
			last_name,
			middle_name,
			date_of_birth,
			phone_number,
			avatar,
			iin,
			role,
			is_active,
			password
		FROM users 
		WHERE email = $1
	`
	err := repo.db.Get(&user, query, email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (repo *AuthenticationRepository) GetUserByID(id int) (*models.User, error) {
	var user models.User
	query := `
		SELECT 
			id, 
			email,
			first_name,
			last_name, 
			middle_name, 
			phone_number, 
			date_of_birth,
			avatar, 
			iin,
			is_active,
			date_joined 
		FROM 
			users
		 WHERE id = $1
	`

	err := repo.db.Get(&user, query, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *AuthenticationRepository) CreateUser(user *models.CreateUser) error {
	query := `
		INSERT INTO users (
			email, 
			first_name, 
			last_name, 
			middle_name, 
			date_of_birth, 
			phone_number, 
			avatar, 
			iin, 
			role, 
			is_active, 
			password, 
			date_joined,
			created_at,
			updated_at,
			is_superuser,
			is_admin,
			is_staff
		) VALUES (
			:email, 
			:first_name, 
			:last_name, 
			:middle_name, 
			:date_of_birth, 
			:phone_number, 
			:avatar, 
			:iin, 
			:role, 
			:is_active, 
			:password, 
			:date_joined,
			NOW(),
			NOW(),
			false,
			false,
			false
		)
	`

	_, err := repo.db.NamedExec(query, user)
	if err != nil {
		return errors.New("failed to create user: " + err.Error())
	}

	return nil
}
