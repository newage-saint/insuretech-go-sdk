package models

import (
	"time"
)

// ServiceProvider represents a service_provider
type ServiceProvider struct {
	ServicesOffered            []string             `json:"services_offered,omitempty"`
	ProviderName               string               `json:"provider_name,omitempty"`
	Longitude                  float64              `json:"longitude,omitempty"`
	SupportedProductCategories []string             `json:"supported_product_categories,omitempty"`
	Address                    string               `json:"address,omitempty"`
	PhoneNumber                string               `json:"phone_number,omitempty"`
	Email                      string               `json:"email,omitempty"`
	IsNetworkProvider          bool                 `json:"is_network_provider,omitempty"`
	ProviderId                 string               `json:"provider_id,omitempty"`
	District                   string               `json:"district,omitempty"`
	CreatedAt                  time.Time            `json:"created_at,omitempty"`
	UpdatedAt                  time.Time            `json:"updated_at,omitempty"`
	ProviderType               *ServiceProviderType `json:"provider_type,omitempty"`
	City                       string               `json:"city,omitempty"`
	Latitude                   float64              `json:"latitude,omitempty"`
}
