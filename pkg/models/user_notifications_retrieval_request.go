package models

// UserNotificationsRetrievalRequest represents a user_notifications_retrieval_request
type UserNotificationsRetrievalRequest struct {
	UnreadOnly bool   `json:"unread_only,omitempty"`
	Limit      int    `json:"limit,omitempty"`
	Offset     int    `json:"offset,omitempty"`
	UserId     string `json:"user_id"`
}
