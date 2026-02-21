package models

import (
	"time"
)

// Peer represents a peer
type Peer struct {
	PeerId      string                 `json:"peer_id"`
	DisplayName string                 `json:"display_name"`
	State       interface{}            `json:"state"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	LastSeenAt  time.Time              `json:"last_seen_at"`
	UserAgent   string                 `json:"user_agent,omitempty"`
	LeftAt      time.Time              `json:"left_at,omitempty"`
	RoomId      string                 `json:"room_id"`
	Tracks      []*Track               `json:"tracks,omitempty"`
	JoinedAt    time.Time              `json:"joined_at"`
}
