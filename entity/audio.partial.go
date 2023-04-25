package entity

type AudioFormat uint8

const (
	AudioJSONFormat AudioFormat = iota
	AudioSRTFormat
	AudioVTTFormat
	AudioTextFormat
	AudioVerboseJSONFormat
)

func (a AudioFormat) String() string {
	switch a {
	case AudioJSONFormat:
		return "json"
	case AudioSRTFormat:
		return "srt"
	case AudioVTTFormat:
		return "vtt"
	case AudioTextFormat:
		return "text"
	case AudioVerboseJSONFormat:
		return "verbose_json"
	default:
		return "json"
	}
}
