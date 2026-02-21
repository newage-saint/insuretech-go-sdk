package models

// JoinRoomRequest represents a join_room_request
type JoinRoomRequest struct {
	RoomId       string                 `json:"room_id"`
	JoinToken    string                 `json:"join_token,omitempty"`
	DisplayName  string                 `json:"display_name,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
	Capabilities []*MediaCapability     `json:"capabilities,omitempty"`
}
