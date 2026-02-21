package models

// FileRetrievalRequest represents a file_retrieval_request
type FileRetrievalRequest struct {
	TenantId string `json:"tenant_id"`
	FileId   string `json:"file_id"`
}
