package models

// MyTasksListingRequest represents a my_tasks_listing_request
type MyTasksListingRequest struct {
	PageSize int    `json:"page_size,omitempty"`
	Status   string `json:"status"`
	Priority string `json:"priority,omitempty"`
	Page     int    `json:"page,omitempty"`
}
