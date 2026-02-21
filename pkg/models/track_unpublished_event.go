package models

import (
	"time"
)

// TrackUnpublishedEvent represents a track_unpublished_event
type TrackUnpublishedEvent struct {
	TrackId       string    `json:"track_id,omitempty"`
	UnpublishedAt time.Time `json:"unpublished_at,omitempty"`
	RoomId        string    `json:"room_id,omitempty"`
	PeerId        string    `json:"peer_id,omitempty"`
}
