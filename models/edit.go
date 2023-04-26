package models

import "encoding/json"

type Edit uint8

const (
	TEXT_DAVINCI_EDIT_001 Edit = iota + 1
	CODE_DAVINCI_EDIT_001
)

func (e Edit) String() string {
	switch e {
	case TEXT_DAVINCI_EDIT_001:
		return "text-davinci-edit-001"
	case CODE_DAVINCI_EDIT_001:
		return "code-davinci-edit-001"
	default:
		return ""
	}
}

func (e Edit) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

func (e Edit) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	return nil
}
