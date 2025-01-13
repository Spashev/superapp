package service

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"

	"superapp/internal/models"
	"superapp/internal/repository"
	schema "superapp/internal/schema/auth"
	"superapp/internal/util"
	"superapp/internal/util/token"
)

type AuthenticationService struct {
	repo       *repository.AuthenticationRepository
	tokenMaker *token.JWTMaker
}

func NewAuthenticationService(repo *repository.AuthenticationRepository, tokenMaker *token.JWTMaker) *AuthenticationService {
	return &AuthenticationService{repo: repo, tokenMaker: tokenMaker}
}

func (s *AuthenticationService) Authenticate(req *schema.AuthLoginReq) (*schema.AuthLoginRes, error) {
	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		log.Println("Invalid email, user not found", err)
		return nil, errors.New("invalid email, user not found")
	}

	if err := util.CheckPassword(req.Password, user.Password); err != nil {
		log.Println("Invalid password", err)
		return nil, errors.New("invalid password")
	}

	duration := time.Hour * 24
	accessToken, err := s.tokenMaker.CreateToken(user.Id, duration)
	if err != nil {
		log.Println("Failed to create JWT token:", err)
		return nil, errors.New("failed to generate access token")
	}

	refreshTokenDuration := time.Hour * 24 * 7
	refreshToken, err := s.tokenMaker.CreateToken(user.Id, refreshTokenDuration)
	if err != nil {
		log.Println("Failed to create refresh token:", err)
		return nil, errors.New("failed to generate refresh token")
	}

	return &schema.AuthLoginRes{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthenticationService) Register(req *schema.RegisterReq) error {
	existingUser, _ := s.repo.GetUserByEmail(req.Email)
	if existingUser != nil {
		return errors.New("email already registered")
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return errors.New("failed to hash password")
	}

	user := &models.CreateUser{
		Email:         req.Email,
		First_name:    req.FirstName,
		Last_name:     req.LastName,
		Middle_name:   req.MiddleName,
		Date_of_birth: req.DateOfBirth,
		Phone_number:  req.PhoneNumber,
		Avatar:        "",
		IIN:           req.IIN,
		Role:          "client",
		Is_active:     true,
		Password:      hashedPassword,
		Date_joined:   time.Now().Format("2025-01-02 00:00:01"),
	}

	if err := s.repo.CreateUser(user); err != nil {
		return err
	}

	return nil
}

func (s *AuthenticationService) UserMe(c *fiber.Ctx) (*models.User, error) {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return nil, errors.New("authorization header is missing")
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		return nil, errors.New("invalid authorization format")
	}

	claims, err := s.tokenMaker.VerifyToken(tokenString)
	if err != nil {
		return nil, errors.New("invalid or expired token")
	}

	user, err := s.repo.GetUserByID(claims.UserID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}
