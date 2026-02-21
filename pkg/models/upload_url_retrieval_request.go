package models

// UploadURLRetrievalRequest represents a upload_url_retrieval_request
type UploadURLRetrievalRequest struct {
	TenantId         string    `json:"tenant_id"`
	Filename         string    `json:"filename,omitempty"`
	ContentType      string    `json:"content_type,omitempty"`
	FileType         *FileType `json:"file_type,omitempty"`
	ExpiresInMinutes int       `json:"expires_in_minutes,omitempty"`
}
