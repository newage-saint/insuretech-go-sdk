package models

// DocumentAnalysisRequest represents a document_analysis_request
type DocumentAnalysisRequest struct {
	DocumentType string `json:"document_type,omitempty"`
	DocumentData string `json:"document_data,omitempty"`
	DocumentUrl  string `json:"document_url"`
}
