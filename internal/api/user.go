package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zaenalarifin12/test-ecommerce/internal/domain"
	authDto "github.com/zaenalarifin12/test-ecommerce/internal/dto/user"
	"net/http"
)

type userApi struct {
	userService domain.UserService
}

func NewUser(router *gin.Engine, userService domain.UserService, middlewares ...gin.HandlerFunc) {
	handler := userApi{userService: userService}

	v1 := router.Group("/api/v1")
	v1.Use(middlewares...)
	{
		v1.POST("/register", handler.Register)
		v1.POST("/login", handler.Login)
	}

}

func (u userApi) Register(ctx *gin.Context) {
	var userInput authDto.RegisterRequest
	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		fmt.Println(userInput)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Call the UserService to register the user
	if _, err := u.userService.Register(ctx, userInput); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "data": userInput})
}

func (u userApi) Login(ctx *gin.Context) {
	var loginInput authDto.LoginRequest
	if err := ctx.ShouldBindJSON(&loginInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the UserService to authenticate the user
	user, err := u.userService.Authenticate(ctx, loginInput)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Send the token in the response
	ctx.JSON(http.StatusOK, gin.H{"data": user})
}
