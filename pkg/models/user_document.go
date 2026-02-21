package models

import (
	"time"
)

// UserDocument represents a user_document
type UserDocument struct {
	UserDocumentId     string    `json:"user_document_id,omitempty"`
	PolicyId           string    `json:"policy_id,omitempty"`
	VerifiedBy         string    `json:"verified_by,omitempty"`
	VerifiedAt         time.Time `json:"verified_at,omitempty"`
	UserId             string    `json:"user_id,omitempty"`
	DocumentTypeId     string    `json:"document_type_id,omitempty"`
	FileUrl            string    `json:"file_url,omitempty"`
	VerificationStatus string    `json:"verification_status,omitempty"`
	CreatedAt          time.Time `json:"created_at,omitempty"`
	UpdatedAt          time.Time `json:"updated_at,omitempty"`
}
