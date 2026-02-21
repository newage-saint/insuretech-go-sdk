package models

// RoomRetrievalResponse represents a room_retrieval_response
type RoomRetrievalResponse struct {
	Room  *Room   `json:"room,omitempty"`
	Peers []*Peer `json:"peers,omitempty"`
}
