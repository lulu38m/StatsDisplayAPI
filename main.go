package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/stats/contributions", langageStats)

	router.Run("localhost:8080")
}
