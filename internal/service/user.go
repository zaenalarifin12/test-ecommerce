package service

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/zaenalarifin12/test-ecommerce/internal/domain"
	authDto "github.com/zaenalarifin12/test-ecommerce/internal/dto/user"
	"github.com/zaenalarifin12/test-ecommerce/internal/utils"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

type userService struct {
	userRepository domain.UserRepository
}

func NewUser(userRepository domain.UserRepository) domain.UserService {
	return &userService{userRepository: userRepository}
}

// Authenticate authenticates a user with the provided credentials and generates a JWT token.
func (u *userService) Authenticate(ctx context.Context, req authDto.LoginRequest) (authDto.LoginResponse, error) {
	// Find the user by email
	user, err := u.userRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		return authDto.LoginResponse{}, err
	}
	// Compare the stored password hash with the provided password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return authDto.LoginResponse{}, domain.ErrInvalidCredentials
	}

	// Generate JWT token claims
	claims := jwt.MapClaims{
		"email":    user.Email,
		"username": user.Username,
		"level":    user.Level,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	// Generate JWT token
	tokenString, err := utils.GenerateToken(claims)
	if err != nil {
		fmt.Println(err.Error())
		return authDto.LoginResponse{}, err
	}

	// Return a successful response with the generated token
	return authDto.LoginResponse{
		Email:    user.Email,
		Username: user.Username,
		Level:    strconv.FormatInt(user.Level, 10),
		Token:    tokenString,
	}, nil
}

// Register registers a new user.
func (u *userService) Register(ctx context.Context, req authDto.RegisterRequest) (authDto.RegisterResponse, error) {
	// Check if the email already exists
	err := u.userRepository.CheckEmailExistence(ctx, req.Email)
	if err != nil && err != domain.ErrEmailExists {
		return authDto.RegisterResponse{}, err
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return authDto.RegisterResponse{}, err
	}

	// Create a new user
	user := &domain.User{
		Email:    req.Email,
		Username: req.Username,
		Password: string(hashedPassword),
		Level:    req.Level,
	}

	// Insert the user into the database
	err = u.userRepository.Insert(ctx, user)
	if err != nil {
		fmt.Println(err.Error())

		return authDto.RegisterResponse{}, err
	}
	// Return a successful response
	return authDto.RegisterResponse{
		Email:    user.Email,
		Username: user.Username,
		Level:    user.Level,
	}, nil
}
