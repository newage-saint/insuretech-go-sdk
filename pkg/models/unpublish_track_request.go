package models

// UnpublishTrackRequest represents a unpublish_track_request
type UnpublishTrackRequest struct {
	RoomId  string `json:"room_id"`
	PeerId  string `json:"peer_id"`
	TrackId string `json:"track_id"`
}
