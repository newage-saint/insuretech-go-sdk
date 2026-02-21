package models

// UserPoliciesListingResponse represents a user_policies_listing_response
type UserPoliciesListingResponse struct {
	Error      *Error    `json:"error,omitempty"`
	Policies   []*Policy `json:"policies,omitempty"`
	TotalCount int       `json:"total_count,omitempty"`
}
