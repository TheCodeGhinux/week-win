package service

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type IntegrationService struct{}

func GetBaseURL(c *gin.Context) string {
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	host := c.Request.Host

	return scheme + "://" + host
}

func (is *IntegrationService) GetIntegrationsJSON(c *gin.Context) map[string]interface{} {

	baseUrl := GetBaseURL(c)

	log.Println(baseUrl)
	targetURL := fmt.Sprintf("%s/api/v1/tick", baseUrl)

	return map[string]interface{}{
		"date": map[string]string{
			"created_at": "2025-02-21",
			"updated_at": "2025-02-21",
		},
		"descriptions": map[string]string{
			"app_description":  "A brief description of the application functionality.",
			"app_logo":         "https://asubeb.anambrastate.gov.ng/?c=107938849870",
			"app_name":         "Week-win Bot",
			"app_url":          "https://6vxj0rsr-8000.uks1.devtunnels.ms/api/v1/integration-spec",
			"background_color": "#00a400",
		},
		"integration_category": "Social Media Management",
		"integration_type":     "interval",
		"is_active":            true,
		"output": []map[string]interface{}{
			{
				"label": "output_channel_1",
				"value": true,
			},
		},
		"key_features": []string{
			"Feature description 1",
			"Feature description 2",
			"Feature description 3",
			"Feature description 4",
		},
		"permissions": map[string]interface{}{
			"monitoring_user": map[string]interface{}{
				"always_online": true,
				"display_name":  "Performance Monitor",
			},
		},
		"settings": []map[string]interface{}{
			{
				"label":    "Delivery Time",
				"type":     "text",
				"required": true,
				"default":  "2 * * * *",
			},
			{
				"label":    "Source",
				"type":     "dropdown",
				"required": true,
				"default":  "Random",
				"options": []string{
					"Random",
					"Psalms",
					"Proverbs",
					"Gospels",
					"Hope",
					"Comfort",
					"Wisdom",
				},
			},
			// {
			// 	"label":    "Alert Admin",
			// 	"type":     "multi-checkbox",
			// 	"required": true,
			// 	"default":  "Super-Admin",
			// 	"options":  []string{"Super-Admin", "Admin", "Manager", "Developer"},
			// },
		},
		"tick_url": targetURL,
		"target_url":  nil,
	}
}
