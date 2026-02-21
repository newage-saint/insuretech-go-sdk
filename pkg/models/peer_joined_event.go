package models

import (
	"time"
)

// PeerJoinedEvent represents a peer_joined_event
type PeerJoinedEvent struct {
	RoomId   string    `json:"room_id,omitempty"`
	Peer     *Peer     `json:"peer,omitempty"`
	JoinedAt time.Time `json:"joined_at,omitempty"`
}
