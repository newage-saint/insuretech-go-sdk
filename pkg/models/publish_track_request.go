package models

// PublishTrackRequest represents a publish_track_request
type PublishTrackRequest struct {
	Track  *Track `json:"track,omitempty"`
	RoomId string `json:"room_id"`
	PeerId string `json:"peer_id"`
}
