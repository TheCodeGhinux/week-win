package service

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type IntegrationService struct{}

type IntegrationResponse struct {
	Date                DateInfo     `json:"date"`
	Descriptions        Descriptions `json:"descriptions"`
	IntegrationCategory string       `json:"integration_category"`
	IntegrationType     string       `json:"integration_type"`
	IsActive            bool         `json:"is_active"`
	Output              []Output     `json:"output"`
	KeyFeatures         []string     `json:"key_features"`
	Permissions         Permissions  `json:"permissions"`
	Settings            []Setting    `json:"settings"`
	TickURL             string       `json:"tick_url"`
	TargetURL           string       `json:"target_url"`
}

type DateInfo struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Descriptions struct {
	AppDescription  string `json:"app_description"`
	AppLogo         string `json:"app_logo"`
	AppName         string `json:"app_name"`
	AppURL          string `json:"app_url"`
	BackgroundColor string `json:"background_color"`
}

type Output struct {
	Label string `json:"label"`
	Value bool   `json:"value"`
}

type Permissions struct {
	MonitoringUser MonitoringUser `json:"monitoring_user"`
}

type MonitoringUser struct {
	AlwaysOnline bool   `json:"always_online"`
	DisplayName  string `json:"display_name"`
}

type Setting struct {
	Label    string   `json:"label"`
	Type     string   `json:"type"`
	Required bool     `json:"required"`
	Default  string   `json:"default"`
	Options  []string `json:"options,omitempty"`
}

type MonitorPayload struct {
	ChannelID string        `json:"channel_id,omitempty"`
	ReturnURL string        `json:"return_url,omitempty"`
	Settings  []interface{} `json:"settings,omitempty"`
}

func GetBaseURL(c *gin.Context) string {
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	host := c.Request.Host

	return scheme + "://" + host
}

func (is *IntegrationService) GetIntegrationsJSON(c *gin.Context) IntegrationResponse {
	baseUrl := GetBaseURL(c)
	log.Println("Base Url in integration: ", baseUrl)
	tickURL := fmt.Sprintf("%s/api/v1/tick", baseUrl)
	log.Println("Tick Url in integration: ", tickURL)

	return IntegrationResponse{
		Date: DateInfo{
			CreatedAt: "2025-02-21",
			UpdatedAt: "2025-02-21",
		},
		Descriptions: Descriptions{
			AppDescription:  "This is a weekly wins bot for teams",
			AppLogo:         "https://i.postimg.cc/L5bv01Px/gr-stocks-Iq9-Sa-Jezk-OE-unsplash.jpg",
			AppName:         "Week-win Bot",
			AppURL:          baseUrl,
			BackgroundColor: "#00a400",
		},
		IntegrationCategory: "Communication & Collaboration",
		IntegrationType:     "interval",
		IsActive:            true,
		Output: []Output{
			{Label: "output_channel_1", Value: true},
		},
		KeyFeatures: []string{
			"Feature description 1",
			"Feature description 2",
			"Feature description 3",
			"Feature description 4",
		},
		Permissions: Permissions{
			MonitoringUser: MonitoringUser{
				AlwaysOnline: true,
				DisplayName:  "Weekly wins Bot",
			},
		},
		Settings: []Setting{
			{
				Label:    "interval",
				Type:     "dropdown",
				Required: true,
				Default:  "2 * * * *",
				Options: []string{
					"0 9 * * 1", "* * * * *", "15 9 * * 1", "25 9 * * 1", "2 * * * *", "5 * * * *",
				},
			},
			{
				Label:    "source",
				Type:     "dropdown",
				Required: true,
				Default:  "Google Drive",
				Options: []string{
					"Google Form",
				},
			},
		},
		TickURL:   tickURL,
		TargetURL: baseUrl,
	}
}
