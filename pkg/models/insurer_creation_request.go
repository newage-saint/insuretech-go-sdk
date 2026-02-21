package models

// InsurerCreationRequest represents a insurer_creation_request
type InsurerCreationRequest struct {
	NameBn             string `json:"name_bn,omitempty"`
	TinNumber          string `json:"tin_number,omitempty"`
	Address            string `json:"address,omitempty"`
	Code               string `json:"code,omitempty"`
	Name               string `json:"name"`
	Type               string `json:"type"`
	TradeLicenseNumber string `json:"trade_license_number,omitempty"`
	Email              string `json:"email"`
	Phone              string `json:"phone,omitempty"`
}
