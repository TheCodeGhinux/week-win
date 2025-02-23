package routers

import (
	"fmt"

	controller "github.com/TheCodeGhinux/week-win/pkg/controllers/accomplishments"
	"github.com/gin-gonic/gin"
)

func Accomplishments(router *gin.Engine, ApiVersion string) *gin.Engine {
	AccomplishmentsController := controller.AccomplishmentHandler{}

	AccomplishmentsGroup := router.Group(fmt.Sprintf("%v", ApiVersion))
	{
		AccomplishmentsGroup.POST("/tick", AccomplishmentsController.GetWeeklyAccomplishments)
	}

	return router
}
