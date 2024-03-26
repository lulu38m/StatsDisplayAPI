package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println()
	router := gin.Default()

	router.GET("/stats/language", langageStats)

	router.Run(":8080")
}
