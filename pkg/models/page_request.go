package models

// PageRequest represents a page_request
type PageRequest struct {
	PageSize  int    `json:"page_size,omitempty"`
	PageToken string `json:"page_token,omitempty"`
	Page      int    `json:"page"`
}
