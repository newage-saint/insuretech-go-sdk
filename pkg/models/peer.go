package models

import (
	"time"
)

// Peer represents a peer
type Peer struct {
	JoinedAt    time.Time              `json:"joined_at"`
	LastSeenAt  time.Time              `json:"last_seen_at"`
	PeerId      string                 `json:"peer_id"`
	UserAgent   string                 `json:"user_agent,omitempty"`
	LeftAt      time.Time              `json:"left_at,omitempty"`
	RoomId      string                 `json:"room_id"`
	DisplayName string                 `json:"display_name"`
	State       interface{}            `json:"state"`
	Tracks      []*Track               `json:"tracks,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}
