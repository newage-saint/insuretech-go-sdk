package models

// UserRolesRetrievalRequest represents a user_roles_retrieval_request
type UserRolesRetrievalRequest struct {
	TenantId string `json:"tenant_id"`
	UserId   string `json:"user_id"`
}
