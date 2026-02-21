package models

// DocumentsListingRequest represents a documents_listing_request
type DocumentsListingRequest struct {
	EntityType string `json:"entity_type"`
	EntityId   string `json:"entity_id"`
	Status     string `json:"status,omitempty"`
	Page       int    `json:"page,omitempty"`
	PageSize   int    `json:"page_size,omitempty"`
}
