package models

import (
	"time"
)

// Room represents a room
type Room struct {
	ParticipantCount int                    `json:"participant_count"`
	CreatorId        string                 `json:"creator_id,omitempty"`
	Config           interface{}            `json:"config"`
	MaxParticipants  int                    `json:"max_participants"`
	CreatedAt        time.Time              `json:"created_at"`
	Metadata         map[string]interface{} `json:"metadata,omitempty"`
	State            interface{}            `json:"state"`
	ClosedAt         time.Time              `json:"closed_at,omitempty"`
	RoomId           string                 `json:"room_id"`
	Name             string                 `json:"name"`
}
