package models

import (
	"time"
)

// UserProfile represents a user_profile
type UserProfile struct {
	Gender                   *AuthnGender `json:"gender,omitempty"`
	KycVerified              bool         `json:"kyc_verified,omitempty"`
	UpdatedAt                time.Time    `json:"updated_at,omitempty"`
	EmergencyContactName     string       `json:"emergency_contact_name,omitempty"`
	IdUploadFrontUrl         string       `json:"id_upload_front_url,omitempty"`
	ProofOfAddressUrl        string       `json:"proof_of_address_url,omitempty"`
	DateOfBirth              time.Time    `json:"date_of_birth,omitempty"`
	AddressLine1             string       `json:"address_line1,omitempty"`
	ProfilePhotoUrl          string       `json:"profile_photo_url,omitempty"`
	MaritalStatus            string       `json:"marital_status,omitempty"`
	FullName                 string       `json:"full_name,omitempty"`
	District                 string       `json:"district,omitempty"`
	CreatedAt                time.Time    `json:"created_at,omitempty"`
	PermanentAddress         string       `json:"permanent_address,omitempty"`
	KycVerifiedAt            time.Time    `json:"kyc_verified_at,omitempty"`
	IdType                   string       `json:"id_type,omitempty"`
	ConsentPrivacyAcceptance bool         `json:"consent_privacy_acceptance,omitempty"`
	PostalCode               string       `json:"postal_code,omitempty"`
	NidNumber                string       `json:"nid_number,omitempty"`
	Occupation               string       `json:"occupation,omitempty"`
	City                     string       `json:"city,omitempty"`
	Division                 string       `json:"division,omitempty"`
	Country                  string       `json:"country,omitempty"`
	IdUploadBackUrl          string       `json:"id_upload_back_url,omitempty"`
	UserId                   string       `json:"user_id,omitempty"`
	AddressLine2             string       `json:"address_line2,omitempty"`
	Employer                 string       `json:"employer,omitempty"`
	EmergencyContactNumber   string       `json:"emergency_contact_number,omitempty"`
	PhotographSelfieUrl      string       `json:"photograph_selfie_url,omitempty"`
}
