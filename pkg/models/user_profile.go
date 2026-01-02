package models

import (
	"time"
)

// UserProfile represents a user_profile
type UserProfile struct {
	KycVerified     bool         `json:"kyc_verified"`
	CreatedAt       time.Time    `json:"created_at"`
	AddressLine1    string       `json:"address_line1"`
	Division        string       `json:"division"`
	PostalCode      string       `json:"postal_code,omitempty"`
	FullName        string       `json:"full_name"`
	DateOfBirth     time.Time    `json:"date_of_birth"`
	Occupation      string       `json:"occupation,omitempty"`
	AddressLine2    string       `json:"address_line2,omitempty"`
	NidNumber       string       `json:"nid_number"`
	Gender          *AuthnGender `json:"gender"`
	City            string       `json:"city"`
	ProfilePhotoUrl string       `json:"profile_photo_url,omitempty"`
	KycVerifiedAt   time.Time    `json:"kyc_verified_at,omitempty"`
	UpdatedAt       time.Time    `json:"updated_at"`
	UserId          string       `json:"user_id"`
	District        string       `json:"district"`
	Country         string       `json:"country"`
}
