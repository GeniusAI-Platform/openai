package models

import "encoding/json"

type Chat uint8

const (
	// GPT4 (GPT-4 model) More capable than any GPT-3.5 model, able to do more complex tasks, and optimized for chat. Will be updated with our latest model iteration.
	// training data: Up to Sep 2021
	GPT4 Chat = iota + 1
	// GPT4_0314 (GPT-4 model) Snapshot of gpt-4 from March 14th 2023. Unlike gpt-4, this model will not receive updates, and will be deprecated 3 months after a new version is released.
	// training data: Up to Sep 2021
	GPT4_0314
	// GPT4_32K (GPT-4 model) Same capabilities as the base gpt-4 mode but with 4x the context length. Will be updated with our latest model iteration.
	// training data: Up to Sep 2021
	GPT4_32K
	// GPT4_32K_0314 (GPT-4 model) Snapshot of gpt-4-32 from March 14th 2023. Unlike gpt-4-32k, this model will not receive updates, and will be deprecated 3 months after a new version is released.
	// training data: Up to Sep 2021
	GPT4_32K_0314

	// GPT35_TURBO (GPT-3.5 model) Most capable GPT-3.5 model and optimized for chat at 1/10th the cost of text-davinci-003. Will be updated with our latest model iteration.
	// training data: Up to Sep 2021
	GPT35_TURBO
	// GPT35_TURBO_0301 (GPT-3.5 model) Snapshot of gpt-3.5-turbo from March 1st 2023. Unlike gpt-3.5-turbo, this model will not receive updates, and will be deprecated 3 months after a new version is released.
	// training data: Up to Sep 2021
	GPT35_TURBO_0301
)

func (c Chat) String() string {
	switch c {
	case GPT4:
		return "gpt-4"
	case GPT4_0314:
		return "gpt-4-0314"
	case GPT4_32K:
		return "gpt-4-32k"
	case GPT4_32K_0314:
		return "gpt-4-32k-0314"
	case GPT35_TURBO:
		return "gpt-3.5-turbo"
	case GPT35_TURBO_0301:
		return "gpt-3.5-turbo-0301"
	default:
		return ""
	}
}

func (c Chat) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

func (c Chat) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	return nil
}
