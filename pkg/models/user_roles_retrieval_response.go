package models

// UserRolesRetrievalResponse represents a user_roles_retrieval_response
type UserRolesRetrievalResponse struct {
	Error *Error  `json:"error,omitempty"`
	Roles []*Role `json:"roles,omitempty"`
}
