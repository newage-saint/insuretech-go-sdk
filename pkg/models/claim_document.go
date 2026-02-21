package models

import (
	"time"
)

// ClaimDocument represents a claim_document
type ClaimDocument struct {
	DocumentId   string    `json:"document_id,omitempty"`
	FileUrl      string    `json:"file_url,omitempty"`
	Verified     bool      `json:"verified,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	ClaimId      string    `json:"claim_id,omitempty"`
	DocumentType string    `json:"document_type,omitempty"`
	FileHash     string    `json:"file_hash,omitempty"`
	UploadedAt   time.Time `json:"uploaded_at,omitempty"`
	VerifiedBy   string    `json:"verified_by,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}
