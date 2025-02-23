package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type AccomplishmentsService struct{}

type WebhookPayload struct {
	EventName string `json:"event_name"`
	Message   string `json:"message"`
	Status    string `json:"status"`
	Username  string `json:"username"`
	Gif       string `json:"gif,omitempty"`
}

// func NewAccomplishments() *AccomplishmentsService {
// 	return &AccomplishmentsService{}
// }

func (s *AccomplishmentsService) collectAccomplishments() []string {
	allAccomplishments := []string{
		"@UserA closed 10 deals",
		"@UserB shipped a new feature",
		"@UserC resolved 5 critical bugs",
		"@UserD led a successful team meeting",
		"@UserE improved system performance by 20%",
		"@UserF deployed a major release",
		"@UserG optimized database queries",
		"@UserH refactored legacy codebase",
		"@UserI conducted a successful training session",
		"@UserJ fixed a major security vulnerability",
		"@UserK launched a marketing campaign",
		"@UserL onboarded 5 new clients",
	}

	// Seed random generator
	rand.Seed(time.Now().UnixNano())

	// Randomly shuffle the list
	rand.Shuffle(len(allAccomplishments), func(i, j int) {
		allAccomplishments[i], allAccomplishments[j] = allAccomplishments[j], allAccomplishments[i]
	})

	// Pick a random number of accomplishments (between 1 and 5)
	n := rand.Intn(5) + 1 // Generates a number between 1 and 5
	return allAccomplishments[:n]
}

func (s *AccomplishmentsService) formatMessages(responses []string) string {
	message := "This week's wins:\n"

	for _, response := range responses {
		message += "- " + response + "\n"
	}

	return message
}

// Get a random celebratory GIF (mocked list)
func (s *AccomplishmentsService) getCelebratoryGif() string {
	gifs := []string{
		"https://media0.giphy.com/media/v1.Y2lkPTc5MGI3NjExZDN5Mml1ejA4ZjVscXdzbncweHMwcjJudDc1Z2Z2emkzOGcxdWxobSZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/o75ajIFH0QnQC3nCeD/giphy.gif",
		"https://media1.giphy.com/media/v1.Y2lkPTc5MGI3NjExZXhudXIwMHVwNG5hY3F2Ynh6emg2Z25oYnh6eHByc2ZiOW4wd2tjeCZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/t3sZxY5zS5B0z5zMIz/giphy.gif",
		"https://media3.giphy.com/media/v1.Y2lkPTc5MGI3NjExNHE0aHdmdG5mdDZmaHcxdWcxY3YxOWF2OHRmMjMyZG9iajdkNG9tMiZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/l41lYCDgxP6OFBruE/giphy.gif",
		"https://media0.giphy.com/media/v1.Y2lkPTc5MGI3NjExbnRqbDZramdjNXg4dHF5amxhb2E2enVjemM3N3BpeDBhNXkzM2Z3eiZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/3oGRFw7Ypjh6LKFlwQ/giphy.gif",
		"https://media1.giphy.com/media/v1.Y2lkPTc5MGI3NjExMWF2NnV6YXQxMDVxdmhseHlqejk0eHplbjZwcjhwM25qN3lsZnowNSZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/l44Q6Etd5kdSGttXa/giphy.gif",
	}
	rand.Seed(time.Now().UnixNano())
	return gifs[rand.Intn(len(gifs))]
}

// Post message to Telex channel (mock function)
func (s *AccomplishmentsService) postToTelexChannel(message, channel_id, gif string) error {

	baseWebookUrl := "https://ping.telex.im/v1/webhooks/"
	webhookUrl := fmt.Sprintf("%s/%s", baseWebookUrl, channel_id)

	log.Println("Webhook URL: ", webhookUrl)

	payload := WebhookPayload{
		EventName: "Weekly Accomplishments",
		Message:   message,
		Status:    "success",
		Username:  "Accomplishment Bot",
		Gif:       gif,
	}

	// Convert payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	
	// Replace this with an actual API call to post to Telex
	fmt.Println("Message to send to telex: ", message)
	fmt.Println("Posting to telex...... ")


	// Send HHTP POST request
	resp, err := http.Post(webhookUrl, "application/json", bytes.NewBuffer(payloadBytes))

	if err != nil {
		log.Println("Error posting to Telex: ", err)
	}

	defer resp.Body.Close()

		// Log response status
	log.Println("Webhook Response Status:", resp.Status)
	return nil

}

// Weekly scheduled function
func (s *AccomplishmentsService) GetWeeklyAccomplishments(channel_id string, settings any) (map[string]interface{}, error) {

	log.Println("Channel ID: ", channel_id)
	log.Println("Settings: ", settings)
	accomplishments := s.collectAccomplishments()
	if len(accomplishments) == 0 {
		return nil, errors.New("No accomplishments to report this week")
	}

	message := s.formatMessages(accomplishments)
	gif := s.getCelebratoryGif()
	s.postToTelexChannel(message, channel_id, gif)

	result := map[string]interface{}{
		"message":    message,
		"username":   "Weekly Accomplishment",
		"event_name": "Weekly Accomplishment",
		"status":     "success",
	}

	// Log the result
	log.Println("Weekly Accomplishments:", result)

	return result, nil

}

// func (s *AccomplishmentsService) GetWeeklyAccomplishments() (map[string]interface{}, error) {
// 	accomplishments := s.collectAccomplishments()
// 	if len(accomplishments) == 0 {
// 		return nil, errors.New("No accomplishments to report this week")
// 	}

// 	message := s.formatMessages(accomplishments)
// 	gif := s.getCelebratoryGif()
// 	s.postToTelexChannel(message, gif)

// 	return map[string]interface{}{
// 		"message": message,
// 		"username": "Weekly Accomplishment",
// 		"event_name": "Weekly Accomplishment",
// 		"status":  "sucess",
// 	}, nil

// }
