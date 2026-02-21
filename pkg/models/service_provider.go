package models

import (
	"time"
)

// ServiceProvider represents a service_provider
type ServiceProvider struct {
	ProviderId                 string               `json:"provider_id,omitempty"`
	Email                      string               `json:"email,omitempty"`
	ServicesOffered            []string             `json:"services_offered,omitempty"`
	SupportedProductCategories []string             `json:"supported_product_categories,omitempty"`
	CreatedAt                  time.Time            `json:"created_at,omitempty"`
	ProviderType               *ServiceProviderType `json:"provider_type,omitempty"`
	District                   string               `json:"district,omitempty"`
	Latitude                   float64              `json:"latitude,omitempty"`
	UpdatedAt                  time.Time            `json:"updated_at,omitempty"`
	ProviderName               string               `json:"provider_name,omitempty"`
	Address                    string               `json:"address,omitempty"`
	City                       string               `json:"city,omitempty"`
	PhoneNumber                string               `json:"phone_number,omitempty"`
	Longitude                  float64              `json:"longitude,omitempty"`
	IsNetworkProvider          bool                 `json:"is_network_provider,omitempty"`
}
