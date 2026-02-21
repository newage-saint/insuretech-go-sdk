package models

import (
	"time"
)

// ServiceProvider represents a service_provider
type ServiceProvider struct {
	ProviderName               string               `json:"provider_name,omitempty"`
	ServicesOffered            []string             `json:"services_offered,omitempty"`
	Email                      string               `json:"email,omitempty"`
	ProviderId                 string               `json:"provider_id,omitempty"`
	ProviderType               *ServiceProviderType `json:"provider_type,omitempty"`
	City                       string               `json:"city,omitempty"`
	District                   string               `json:"district,omitempty"`
	PhoneNumber                string               `json:"phone_number,omitempty"`
	Latitude                   float64              `json:"latitude,omitempty"`
	Longitude                  float64              `json:"longitude,omitempty"`
	CreatedAt                  time.Time            `json:"created_at,omitempty"`
	Address                    string               `json:"address,omitempty"`
	IsNetworkProvider          bool                 `json:"is_network_provider,omitempty"`
	SupportedProductCategories []string             `json:"supported_product_categories,omitempty"`
	UpdatedAt                  time.Time            `json:"updated_at,omitempty"`
}
