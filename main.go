package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/repo", getRepo)

	router.Run("localhost:8080")
}
