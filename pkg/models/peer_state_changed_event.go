package models

import (
	"time"
)

// PeerStateChangedEvent represents a peer_state_changed_event
type PeerStateChangedEvent struct {
	PeerId    string               `json:"peer_id,omitempty"`
	OldState  *PeerConnectionState `json:"old_state,omitempty"`
	NewState  *PeerConnectionState `json:"new_state,omitempty"`
	ChangedAt time.Time            `json:"changed_at,omitempty"`
	RoomId    string               `json:"room_id,omitempty"`
}
