package controllers

import (
	"log"

	services "github.com/TheCodeGhinux/week-win/services/accomplishments"
	IS "github.com/TheCodeGhinux/week-win/services/integrations"
	"github.com/gin-gonic/gin"
)

type AccomplishmentHandler struct {
	Service *services.AccomplishmentsService
}

func (ac *AccomplishmentHandler) GetWeeklyAccomplishments(c *gin.Context) {
	var payload IS.MonitorPayload

	// Try to bind JSON, log error if it fails
	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Println("⚠️ Error binding JSON payload:", err)
		log.Println("⚠️ Proceeding without a valid payload...")
	}

	// Extract fields from payload
	channelID := payload.ChannelID
	returnURL := payload.ReturnURL
	settings := payload.Settings

	// Log missing required fields instead of returning errors
	if channelID == "" {
		log.Println("⚠️ Missing ChannelID in payload")
	}
	if returnURL == "" {
		log.Println("⚠️ Missing ReturnURL in payload")
	}
	if settings == nil {
		log.Println("⚠️ Missing Settings in payload")
	}

	log.Printf("📥 Received Payload - ChannelID: %s, ReturnURL: %s, Settings: %+v\n", channelID, returnURL, settings)

	// Call the service method with extracted data
	result, err := ac.Service.GetWeeklyAccomplishments(channelID, settings)
	if err != nil {
		log.Println("❌ Error processing weekly accomplishments:", err)
		return // No JSON response, just logs
	}

	log.Println("✅ Successfully processed:", result)
}
