package models

import (
	"time"
)

// Room represents a room
type Room struct {
	MaxParticipants  int                    `json:"max_participants"`
	CreatedAt        time.Time              `json:"created_at"`
	Metadata         map[string]interface{} `json:"metadata,omitempty"`
	CreatorId        string                 `json:"creator_id,omitempty"`
	Name             string                 `json:"name"`
	State            interface{}            `json:"state"`
	ClosedAt         time.Time              `json:"closed_at,omitempty"`
	RoomId           string                 `json:"room_id"`
	Config           interface{}            `json:"config"`
	ParticipantCount int                    `json:"participant_count"`
}
