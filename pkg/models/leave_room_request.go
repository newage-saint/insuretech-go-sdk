package models

// LeaveRoomRequest represents a leave_room_request
type LeaveRoomRequest struct {
	Reason string `json:"reason,omitempty"`
	RoomId string `json:"room_id"`
	PeerId string `json:"peer_id"`
}
