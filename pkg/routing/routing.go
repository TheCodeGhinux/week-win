package routing

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TheCodeGhinux/week-win/pkg/routers"
	"github.com/gin-gonic/gin"
)

func Route() {

	// configs := config.LoadConfig()
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Weekly win telex integration",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	// Swagger docs route
	routeRegister(r)

	err := r.Run(fmt.Sprintf("%s:%s", "localhost", "8080"))

	// **Logging registered routes**
	for _, route := range r.Routes() {
		log.Printf("Registered route: %s %s\n", route.Method, route.Path)
	}

	if err != nil {
		log.Fatal("Error starting server in routing: ", err)
		return
	}

}

func routeRegister(router *gin.Engine) {

	apiVersion := "api/v1"
	routers.Greeting(router, apiVersion)
	routers.Integration(router, apiVersion)
	routers.Accomplishments(router, apiVersion)

}
