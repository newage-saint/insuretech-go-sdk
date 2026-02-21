package models

// DocumentGenerationRequest represents a document_generation_request
type DocumentGenerationRequest struct {
	IncludeQrCode bool                   `json:"include_qr_code,omitempty"`
	TemplateId    string                 `json:"template_id"`
	EntityType    string                 `json:"entity_type"`
	EntityId      string                 `json:"entity_id"`
	Data          map[string]interface{} `json:"data,omitempty"`
}
