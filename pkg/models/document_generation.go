package models

import (
	"time"
)

// DocumentGeneration represents a document_generation
type DocumentGeneration struct {
	GeneratedBy        string      `json:"generated_by,omitempty"`
	AuditInfo          interface{} `json:"audit_info"`
	Id                 string      `json:"id"`
	EntityType         string      `json:"entity_type"`
	Data               string      `json:"data"`
	FileSizeBytes      string      `json:"file_size_bytes,omitempty"`
	GeneratedAt        time.Time   `json:"generated_at,omitempty"`
	DocumentTemplateId string      `json:"document_template_id"`
	EntityId           string      `json:"entity_id"`
	Status             interface{} `json:"status"`
	FileUrl            string      `json:"file_url,omitempty"`
	QrCodeData         string      `json:"qr_code_data,omitempty"`
}
