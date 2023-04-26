package models

import "encoding/json"

type Audio uint8

const (
	// WHISPER_1 Whisper is a general-purpose speech recognition model. It is trained on a large dataset of diverse audio and is also a multi-task model that can perform multilingual speech recognition as well as speech translation and language identification
	WHISPER_1 = iota + 1
)

func (a Audio) String() string {
	switch a {
	case WHISPER_1:
		return "whisper-1"
	default:
		return ""
	}
}

func (a Audio) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.String())
}

func (a Audio) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	return nil
}
