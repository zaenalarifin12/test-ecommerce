package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/zaenalarifin12/test-ecommerce/internal/config"
	"strconv"
)

func GetUserID(ctx *gin.Context) (int64, error) {
	userID, ok := ctx.Get("user_id")
	if !ok {
		return 0, errors.New("user_id not found in context")
	}

	idStr, ok := userID.(string)
	if !ok {
		return 0, errors.New("user_id is not a string")
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// GenerateToken generates a JWT token for the given user claims.
func GenerateToken(claims jwt.MapClaims) (string, error) {
	// Create a new JWT token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims
	token.Claims = claims

	// Get the secret key
	secret := GetJWTSecret()

	// Check if secret is valid
	if secret == nil {
		return "", errors.New("failed to get JWT secret")
	}

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetJWTSecret() []byte {
	conf, err := config.LoadConfig(".")
	if err != nil {
		panic("can't load config")
	}

	return []byte(conf.SecretKey)
}
