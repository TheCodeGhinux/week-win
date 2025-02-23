package routers

import (
	"fmt"

	controller "github.com/TheCodeGhinux/week-win/pkg/controllers/integrations"
	"github.com/gin-gonic/gin"
)

func Integration(router *gin.Engine, ApiVersion string) *gin.Engine {
	IntegrationController := controller.IntegrationController{}

	IntegrationGroup := router.Group(fmt.Sprintf("%v", ApiVersion))
	{
		IntegrationGroup.GET("/integration-spec", IntegrationController.GetIntegrationJSON)
	}

	return router
}
