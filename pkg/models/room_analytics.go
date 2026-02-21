package models

// RoomAnalytics represents a room_analytics
type RoomAnalytics struct {
	RoomId           string `json:"room_id,omitempty"`
	PeakParticipants int    `json:"peak_participants,omitempty"`
	TotalSessions    string `json:"total_sessions,omitempty"`
}
