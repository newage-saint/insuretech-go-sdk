package models

import (
	"time"
)

// DocumentGeneration represents a document_generation
type DocumentGeneration struct {
	Id                 string      `json:"id"`
	DocumentTemplateId string      `json:"document_template_id"`
	Data               string      `json:"data"`
	Status             interface{} `json:"status"`
	FileUrl            string      `json:"file_url,omitempty"`
	QrCodeData         string      `json:"qr_code_data,omitempty"`
	EntityType         string      `json:"entity_type"`
	EntityId           string      `json:"entity_id"`
	FileSizeBytes      string      `json:"file_size_bytes,omitempty"`
	GeneratedBy        string      `json:"generated_by,omitempty"`
	GeneratedAt        time.Time   `json:"generated_at,omitempty"`
	AuditInfo          interface{} `json:"audit_info"`
}
