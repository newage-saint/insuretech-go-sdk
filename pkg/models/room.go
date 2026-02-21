package models

import (
	"time"
)

// Room represents a room
type Room struct {
	Name             string                 `json:"name"`
	ParticipantCount int                    `json:"participant_count"`
	MaxParticipants  int                    `json:"max_participants"`
	State            interface{}            `json:"state"`
	ClosedAt         time.Time              `json:"closed_at,omitempty"`
	CreatorId        string                 `json:"creator_id,omitempty"`
	RoomId           string                 `json:"room_id"`
	Config           interface{}            `json:"config"`
	CreatedAt        time.Time              `json:"created_at"`
	Metadata         map[string]interface{} `json:"metadata,omitempty"`
}
