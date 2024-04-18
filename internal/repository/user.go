// file: user_repository.go
package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zaenalarifin12/test-ecommerce/internal/domain"
)

// userRepository represents the PostgreSQL implementation of UserRepository.
type userRepository struct {
	db *pgxpool.Pool
}

// NewUser creates a new instance of UserRepository with the provided database connection pool.
func NewUser(pool *pgxpool.Pool) domain.UserRepository {
	return &userRepository{db: pool}
}

// FindByEmail retrieves a user from the database by their email address.
func (u *userRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User
	err := u.db.QueryRow(ctx, "SELECT id, email, username, password, level FROM users WHERE email = $1", email).Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.Level)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

// CheckEmailExistence checks if an email already exists in the database.
func (u *userRepository) CheckEmailExistence(ctx context.Context, email string) error {
	var exists bool
	err := u.db.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", email).Scan(&exists)
	if err != nil {
		return err
	}

	if exists == false {
		return domain.ErrEmailExists
	}

	return nil
}

// Insert inserts a new user into the database.
func (u *userRepository) Insert(ctx context.Context, user *domain.User) error {
	_, err := u.db.Exec(ctx, "INSERT INTO users (email, username, password, level) VALUES ($1, $2, $3, $4)", user.Email, user.Username, user.Password, user.Level)
	return err
}
