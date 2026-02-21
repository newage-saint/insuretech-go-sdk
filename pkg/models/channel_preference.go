package models

// ChannelPreference represents a channel_preference
type ChannelPreference struct {
	Enabled       bool                 `json:"enabled,omitempty"`
	ExcludedTypes []*NotificationType  `json:"excluded_types,omitempty"`
	Channel       *NotificationChannel `json:"channel,omitempty"`
}
