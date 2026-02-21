package models

// RunQueryRequest represents a run_query_request
type RunQueryRequest struct {
	Query      string                 `json:"query"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	Limit      int                    `json:"limit,omitempty"`
}
