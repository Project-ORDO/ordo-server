package handler

import (
	"context"
	"net/http"

	models "github.com/Project-ORDO/ORDO-backEnd/internal/model"
	"github.com/Project-ORDO/ORDO-backEnd/internal/model/request"
	"github.com/Project-ORDO/ORDO-backEnd/internal/model/response"
	"github.com/Project-ORDO/ORDO-backEnd/internal/repository/implementations"
	"github.com/Project-ORDO/ORDO-backEnd/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
    userService *service.UserService
}

func NewAuthHandler() *AuthHandler {
    return &AuthHandler{
        userService: service.NewUserService(implementations.NewUserRepo()),
    }
}

func (h *AuthHandler) SignUp(c *gin.Context) {
    var user models.User

    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
        return
    }

    err := h.userService.SignUp(context.Background(), &user)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func (h *AuthHandler) Login(c *gin.Context) {
    var loginReq request.LoginRequest

    if err := c.ShouldBindJSON(&loginReq); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    user, err := h.userService.LoginUser(context.Background(), &loginReq)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    userResponse := response.UserResponse{
        Email: user.Email,
        Name:  user.Name,
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Login successful",
        "user":    userResponse,
    })
}
