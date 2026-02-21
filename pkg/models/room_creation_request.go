package models

// RoomCreationRequest represents a room_creation_request
type RoomCreationRequest struct {
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	CreatorId string                 `json:"creator_id"`
	Name      string                 `json:"name"`
	Config    *RoomConfig            `json:"config,omitempty"`
}
