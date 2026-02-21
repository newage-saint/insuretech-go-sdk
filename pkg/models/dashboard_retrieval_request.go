package models

// DashboardRetrievalRequest represents a dashboard_retrieval_request
type DashboardRetrievalRequest struct {
	StartDate   string `json:"start_date,omitempty"`
	EndDate     string `json:"end_date,omitempty"`
	DashboardId string `json:"dashboard_id"`
}
