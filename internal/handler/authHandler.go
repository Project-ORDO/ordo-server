package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/Project-ORDO/ORDO-backEnd/internal/helper"
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
		helper.Responce(c, http.StatusBadRequest, "invalid input", nil, "")
		return
	}

	err := h.userService.SignUp(context.Background(), &user)
	if err != nil {
		helper.Responce(c, http.StatusBadRequest, err.Error(), nil, "")
		return
	}

	helper.Responce(c, http.StatusOK, "User created. Verification email sent.", nil, "")
}

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (h *AuthHandler) VerifyEmail(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		helper.Responce(c, http.StatusBadRequest, "Token is required", nil, "")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := h.userService.VerifyEmail(ctx, token)
	if err != nil {
		helper.Responce(c, http.StatusBadRequest, err.Error(), nil, "")
		return
	}

	c.Redirect(http.StatusSeeOther, "http://localhost:8080/api/v1/auth/signup")
}

func (h *AuthHandler) Login(c *gin.Context) {
	var loginReq request.LoginRequest

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		helper.Responce(c, http.StatusBadRequest, "Invalid input", nil, "")
		return
	}

	user, err := h.userService.LoginUser(context.Background(), &loginReq)
	if err != nil {
		helper.Responce(c, http.StatusUnauthorized, err.Error(), nil, "")
		return
	}

	userResponse := response.UserResponse{
		Email: user.Email,
		Name:  user.Name,
	}

	helper.Responce(c, http.StatusOK, "Login successful", userResponse, "")
}
