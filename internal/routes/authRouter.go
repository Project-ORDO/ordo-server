package routes

import (
	"github.com/Project-ORDO/ORDO-backEnd/internal/handler"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(rg *gin.RouterGroup){
	auth := rg.Group("/auth")
	auth.POST("/signup", handler.SignUpHandler)
}