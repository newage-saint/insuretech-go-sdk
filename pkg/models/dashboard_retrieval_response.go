package models

// DashboardRetrievalResponse represents a dashboard_retrieval_response
type DashboardRetrievalResponse struct {
	Error      *Error                 `json:"error,omitempty"`
	Dashboard  *Dashboard             `json:"dashboard,omitempty"`
	WidgetData map[string]interface{} `json:"widget_data,omitempty"`
}
