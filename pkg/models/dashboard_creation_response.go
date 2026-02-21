package models

// DashboardCreationResponse represents a dashboard_creation_response
type DashboardCreationResponse struct {
	Message     string `json:"message,omitempty"`
	Error       *Error `json:"error,omitempty"`
	DashboardId string `json:"dashboard_id,omitempty"`
}
