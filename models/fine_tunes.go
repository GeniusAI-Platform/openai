package models

import "encoding/json"

type FineTunes uint8

const (
	// DAVINCI (GPT-3 model) Most capable GPT-3 model. Can do any task the other models can do, often with higher quality.
	// training data: Up to Oct 2019
	DAVINCI FineTunes = iota + 1
	// CURIE (GPT-3 model) Very capable, but faster and lower cost than Davinci.
	// training data: Up to Oct 2019
	CURIE
	// BABBAGE (GPT-3 model) Capable of straightforward tasks, very fast, and lower cost.
	// training data: Up to Oct 2019
	BABBAGE
	// ADA (GPT-3 model) Capable of very simple tasks, usually the fastest model in the GPT-3 series, and lowest cost.
	// training data: Up to Oct 2019
	ADA
)

func (f FineTunes) String() string {
	switch f {
	case DAVINCI:
		return "davinci"
	case CURIE:
		return "curie"
	case BABBAGE:
		return "babbage"
	case ADA:
		return "ada"
	default:
		return ""
	}
}

func (f FineTunes) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.String())
}

func (f FineTunes) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	return nil
}
