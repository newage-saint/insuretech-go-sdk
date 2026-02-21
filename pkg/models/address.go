package models

// Address represents a address
type Address struct {
	PostalCode   string  `json:"postal_code,omitempty"`
	Longitude    float64 `json:"longitude,omitempty"`
	AddressLine1 string  `json:"address_line1,omitempty"`
	AddressLine2 string  `json:"address_line2,omitempty"`
	City         string  `json:"city,omitempty"`
	District     string  `json:"district,omitempty"`
	Country      string  `json:"country,omitempty"`
	Latitude     float64 `json:"latitude,omitempty"`
	Division     string  `json:"division,omitempty"`
}
