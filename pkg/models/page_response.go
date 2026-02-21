package models

// PageResponse represents a page_response
type PageResponse struct {
	PageSize      int    `json:"page_size,omitempty"`
	TotalPages    int    `json:"total_pages,omitempty"`
	TotalItems    string `json:"total_items,omitempty"`
	NextPageToken string `json:"next_page_token,omitempty"`
	Page          int    `json:"page,omitempty"`
}
