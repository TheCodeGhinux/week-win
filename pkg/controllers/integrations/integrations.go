package controllers

import (
	"net/http"

	services "github.com/TheCodeGhinux/week-win/services/integrations"
	"github.com/gin-gonic/gin"
)

type IntegrationController struct {
	Service *services.IntegrationService
}

// func NewIntegrationController() *IntegrationController {
// 	service := services.NewIntegrationService() // Service is instantiated here
// 	return &IntegrationController{service: service}
// }

func (ic *IntegrationController) GetIntegrationJSON(c *gin.Context) {

	data := ic.Service.GetIntegrationsJSON(c)

	c.JSON(http.StatusOK, gin.H{"data": data})
}
