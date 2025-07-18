package main

import "github.com/gin-gonic/gin"

var r *gin.Engine

func init() {
	r = gin.Default()
}

func main() {
	r.Run(":8080")
}
