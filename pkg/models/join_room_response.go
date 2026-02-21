package models

// JoinRoomResponse represents a join_room_response
type JoinRoomResponse struct {
	Room          *Room   `json:"room,omitempty"`
	ExistingPeers []*Peer `json:"existing_peers,omitempty"`
	Peer          *Peer   `json:"peer,omitempty"`
}
