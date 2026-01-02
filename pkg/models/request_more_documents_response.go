package models

// RequestMoreDocumentsResponse represents a request_more_documents_response
type RequestMoreDocumentsResponse struct {
	Error   *Error `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}
