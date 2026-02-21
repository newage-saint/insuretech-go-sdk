package models

// DocumentAnalysisRequest represents a document_analysis_request
type DocumentAnalysisRequest struct {
	DocumentData string `json:"document_data,omitempty"`
	DocumentUrl  string `json:"document_url"`
	DocumentType string `json:"document_type,omitempty"`
}
