package models

import (
	"time"
)

// Peer represents a peer
type Peer struct {
	PeerId      string                 `json:"peer_id"`
	DisplayName string                 `json:"display_name"`
	State       interface{}            `json:"state"`
	Tracks      []*Track               `json:"tracks,omitempty"`
	LastSeenAt  time.Time              `json:"last_seen_at"`
	UserAgent   string                 `json:"user_agent,omitempty"`
	RoomId      string                 `json:"room_id"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	JoinedAt    time.Time              `json:"joined_at"`
	LeftAt      time.Time              `json:"left_at,omitempty"`
}
