package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/stats/language", langageStats)
	router.GET("/stats/contributions", contributionStats)

	router.Run(":8080")
}
