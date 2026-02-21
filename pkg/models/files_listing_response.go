package models

// FilesListingResponse represents a files_listing_response
type FilesListingResponse struct {
	Page  *PaginationResponse `json:"page,omitempty"`
	Files []*StoredFile       `json:"files,omitempty"`
}
