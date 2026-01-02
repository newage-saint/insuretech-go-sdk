package models

import (
	"time"
)

// ClaimDocument represents a claim_document
type ClaimDocument struct {
	UpdatedAt    time.Time `json:"updated_at"`
	DocumentId   string    `json:"document_id"`
	FileUrl      string    `json:"file_url"`
	CreatedAt    time.Time `json:"created_at"`
	ClaimId      string    `json:"claim_id"`
	DocumentType string    `json:"document_type"`
	FileHash     string    `json:"file_hash"`
	UploadedAt   time.Time `json:"uploaded_at"`
	Verified     bool      `json:"verified"`
	VerifiedBy   string    `json:"verified_by,omitempty"`
}
