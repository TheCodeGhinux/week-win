package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Greet(c *gin.Context) {
	appName := "Telex Integration app"
	appDesc := "Telex integration using Golang"
	c.JSON(http.StatusOK, gin.H{
		"message":         "Telex Integration app API",
		"app name":        `Welcome to ` + appName,
		"app description": appDesc,
	})
}
