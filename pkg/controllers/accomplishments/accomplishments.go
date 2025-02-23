package controllers

import (
	"log"
	"net/http"

	services "github.com/TheCodeGhinux/week-win/services/accomplishments"
	"github.com/gin-gonic/gin"
)

type AccomplishmentHandler struct {
	Service *services.AccomplishmentsService
}

func (ac *AccomplishmentHandler) GetWeeklyAccomplishments(c *gin.Context) {

	var body struct {
		ReturnURL string `json:"return_url"`
		ChannelID string `json:"channel_id"`
		Settings  any    `json:"settings"` // Change `any` to the correct type if known
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	// Check if both return_url and channel_id are missing
	if body.ReturnURL == "" && body.ChannelID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "No return URL or channel ID provided",
		})
		return
	}

	result, err := ac.Service.GetWeeklyAccomplishments(body.ChannelID, body.Settings)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("Response:", result)
	c.JSON(http.StatusOK, result)
}
