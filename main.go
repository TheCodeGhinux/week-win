package main

import (
	"net/http"

	"github.com/TheCodeGhinux/week-win/pkg/routing"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	routing.Route()

	r.Run("0.0.0.0:8080")
}
