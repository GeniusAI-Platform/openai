package entity

import (
	"encoding/json"
)

type Role uint8

const (
	SYSTEM = iota
	USER
	ASSISTANT
)

func (r Role) String() string {
	switch r {
	case SYSTEM:
		return "system"
	case USER:
		return "user"
	case ASSISTANT:
		return "assistant"
	default:
		return "assistant"
	}
}

func (r Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r Role) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	return nil
}
