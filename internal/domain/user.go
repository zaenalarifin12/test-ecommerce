package domain

import (
	"context"
	authDto "github.com/zaenalarifin12/test-ecommerce/internal/dto/user"
)

type User struct {
	ID       int64  `db:"id"`
	Email    string `db:"email"`
	Username string `db:"username"`
	Password string `db:"password"`
	Level    int64  `db:"level"`
}

type UserRepository interface {
	CheckEmailExistence(ctx context.Context, email string) error
	Insert(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email string) (User, error)
}

type UserService interface {
	Register(ctx context.Context, req authDto.RegisterRequest) (authDto.RegisterResponse, error)
	Authenticate(ctx context.Context, req authDto.LoginRequest) (authDto.LoginResponse, error)
}
