package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    api := r.Group("/api/v1")
    authRouter := NewAuthRouter()
    authRouter.RegisterRoutes(api)
}
