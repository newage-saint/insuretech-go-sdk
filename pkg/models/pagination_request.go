package models

// PaginationRequest represents a pagination_request
type PaginationRequest struct {
	SortOrder *SortOrder `json:"sort_order,omitempty"`
	Page      int        `json:"page"`
	PageSize  int        `json:"page_size,omitempty"`
	SortBy    string     `json:"sort_by,omitempty"`
}
