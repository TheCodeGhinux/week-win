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

	log.Println("Base Url in integration: ", baseUrl)
	targetURL := fmt.Sprintf("%s/api/v1/tick", baseUrl)

	return map[string]interface{}{
		"date": map[string]string{
			"created_at": "2025-02-21",
			"updated_at": "2025-02-21",
		},
		"descriptions": map[string]string{
			"app_description":  "This is a weekly wins bot for teams",
			"app_logo":         "https://i.postimg.cc/L5bv01Px/gr-stocks-Iq9-Sa-Jezk-OE-unsplash.jpg",
			"app_name":         "Week-win Bot",
			"app_url":          "https://week-win.onrender.com/api/v1/integration-spec",
			"background_color": "#00a400",
		},
		"integration_category": "Communication & Collaboration",
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
				"display_name":  "Weekly wins Bot",
			},
		},
		"settings": []map[string]interface{}{
			{
				"label":    "Delivery Time",
				"type":     "dropdown",
				"required": true,
				"default":  "2 * * * *",
				"options": []string{
					"0 9 * * 1",
					"15 9 * * 1",
					"25 9 * * 1",
					"2 * * * *",
					"5 * * * *",
				},
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
