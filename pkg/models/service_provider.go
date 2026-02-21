package models

import (
	"time"
)

// ServiceProvider represents a service_provider
type ServiceProvider struct {
	Longitude                  float64              `json:"longitude,omitempty"`
	CreatedAt                  time.Time            `json:"created_at,omitempty"`
	ProviderId                 string               `json:"provider_id,omitempty"`
	ProviderName               string               `json:"provider_name,omitempty"`
	PhoneNumber                string               `json:"phone_number,omitempty"`
	ServicesOffered            []string             `json:"services_offered,omitempty"`
	IsNetworkProvider          bool                 `json:"is_network_provider,omitempty"`
	SupportedProductCategories []string             `json:"supported_product_categories,omitempty"`
	Address                    string               `json:"address,omitempty"`
	City                       string               `json:"city,omitempty"`
	District                   string               `json:"district,omitempty"`
	Email                      string               `json:"email,omitempty"`
	UpdatedAt                  time.Time            `json:"updated_at,omitempty"`
	ProviderType               *ServiceProviderType `json:"provider_type,omitempty"`
	Latitude                   float64              `json:"latitude,omitempty"`
}
