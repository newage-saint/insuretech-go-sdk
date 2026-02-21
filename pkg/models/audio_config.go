package models

// AudioConfig represents a audio_config
type AudioConfig struct {
	MaxBitrateKbps         int           `json:"max_bitrate_kbps,omitempty"`
	EnableNoiseSuppression bool          `json:"enable_noise_suppression,omitempty"`
	EnableEchoCancellation bool          `json:"enable_echo_cancellation,omitempty"`
	EnableAutoGainControl  bool          `json:"enable_auto_gain_control,omitempty"`
	AllowedCodecs          []*AudioCodec `json:"allowed_codecs,omitempty"`
	DefaultCodec           *AudioCodec   `json:"default_codec,omitempty"`
}
