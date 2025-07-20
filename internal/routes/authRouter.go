package routes

import (
	"github.com/Project-ORDO/ORDO-backEnd/internal/handler"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(rg *gin.RouterGroup){
	auth := rg.Group("/auth")
	authHandler := handler.NewAuthHandler()

    auth.POST("/signup", authHandler.SignUp)
    auth.POST("/login", authHandler.Login)
}