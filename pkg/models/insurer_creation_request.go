package models

// InsurerCreationRequest represents a insurer_creation_request
type InsurerCreationRequest struct {
	Name               string `json:"name"`
	NameBn             string `json:"name_bn,omitempty"`
	TradeLicenseNumber string `json:"trade_license_number,omitempty"`
	Email              string `json:"email"`
	Phone              string `json:"phone,omitempty"`
	Code               string `json:"code,omitempty"`
	Type               string `json:"type"`
	TinNumber          string `json:"tin_number,omitempty"`
	Address            string `json:"address,omitempty"`
}
