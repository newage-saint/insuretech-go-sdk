package models

// DocumentRetrievalResponse represents a document_retrieval_response
type DocumentRetrievalResponse struct {
	Error    *Error              `json:"error,omitempty"`
	Document *DocumentGeneration `json:"document,omitempty"`
}
