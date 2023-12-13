package main

import "github.com/gin-gonic/gin"

type Contributions struct {
}

func contributionStats(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "contributionStats",
	})
}
