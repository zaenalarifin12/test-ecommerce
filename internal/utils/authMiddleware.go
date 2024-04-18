package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware(ctx *gin.Context) {
	// Check authorization token in the request header
	token := ctx.GetHeader("Authorization")

	// Here, you would typically decode the token to extract user information such as user ID and level.
	// For the sake of example, let's assume you have a function `decodeToken` to decode the token.
	claim, err := decodeToken(token)

	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		ctx.Abort()
		return
	}

	// Check if user level is 2
	fmt.Println(claim.Level)
	if claim.Level != 2 {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		ctx.Abort()
		return
	}

	// Continue with the request chain
	ctx.Next()
}

// CustomClaims represents the custom claims in your JWT token
type CustomClaims struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Level    int64  `json:"level"`
	jwt.StandardClaims
}

// DecodeToken decodes the JWT token and returns the claims if valid
func decodeToken(tokenString string) (*CustomClaims, error) {
	// Remove the "Bearer " prefix from the token string
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
