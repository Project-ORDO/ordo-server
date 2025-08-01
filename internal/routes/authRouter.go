package routes

import (
	"github.com/Project-ORDO/ORDO-backEnd/internal/handler"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	handler *handler.AuthHandler
}

func NewAuthRouter() *AuthRouter {
	return &AuthRouter{
		handler: handler.NewAuthHandler(),
	}
}

func (ar *AuthRouter) RegisterRoutes(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")
	auth.POST("/signup", ar.handler.SignUp)
	auth.GET("/verify", ar.handler.VerifyEmail)
	auth.POST("/login", ar.handler.Login)
}
