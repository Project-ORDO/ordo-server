package main

import (
	"github.com/Project-ORDO/ORDO-backEnd/config"
	"github.com/Project-ORDO/ORDO-backEnd/internal/routes"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func init() {
	config.LoadEnvFile()
	r = gin.Default()
	config.ConnectDB()
}

func main() {
	routes.SetupRoutes(r)
	r.Run()
}
