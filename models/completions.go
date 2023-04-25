package models

type CompletionsModel uint8

const (
	// TEXT_DAVINCI_003 (GPT-3.5 model) Can do any language task with better quality, longer output, and consistent instruction-following than the curie, babbage, or ada models. Also supports inserting completions within text.
	// training data: Up to Jun 2021
	TEXT_DAVINCI_003 CompletionsModel = iota + 1
	// TEXT_DAVINCI_002 (GPT-3.5 model) Similar capabilities to text-davinci-003 but trained with supervised fine-tuning instead of reinforcement learning.
	// training data: Up to Jun 2021
	TEXT_DAVINCI_002

	// TEXT_CURIE_001 (GPT-3 model) Very capable, faster and lower cost than Davinci.
	// training data: Up to Oct 2019
	TEXT_CURIE_001
	// TEXT_BABBAGE_001 (GPT-3 model) Capable of straightforward tasks, very fast, and lower cost.
	// training data: Up to Oct 2019
	TEXT_BABBAGE_001
	// TEXT_ADA_001 (GPT-3 model) Capable of very simple tasks, usually the fastest model in the GPT-3 series, and lowest cost.
	// training data: Up to Oct 2019
	TEXT_ADA_001

	// CODE_DAVINCI_002 (GPT-3.5 model) Optimized for code-completion tasks
	CODE_DAVINCI_002
	// CODE_DAVINCI_001 (GPT-3 model) Earlier version of code-davinci-002
	CODE_DAVINCI_001
	// CODE_CUSHMAN_002 (GPT-3 model) Almost as capable as Davinci Codex, but slightly faster. This speed advantage may make it preferable for real-time applications
	CODE_CUSHMAN_002
	// CODE_CUSHMAN_001 (GPT-3 model) Earlier version of code-cushman-002
	CODE_CUSHMAN_001
)

func (c CompletionsModel) String() string {
	switch c {
	case TEXT_DAVINCI_003:
		return "text-davinci-003"
	case TEXT_DAVINCI_002:
		return "text-davinci-002"
	case TEXT_CURIE_001:
		return "text-curie-001"
	case TEXT_BABBAGE_001:
		return "text-babbage-001"
	case TEXT_ADA_001:
		return "text-ada-001"
	case CODE_DAVINCI_002:
		return "code-davinci-002"
	case CODE_DAVINCI_001:
		return "code-davinci-002"
	case CODE_CUSHMAN_002:
		return "code-cushman-002"
	case CODE_CUSHMAN_001:
		return "code-cushman-001"
	default:
		return ""
	}
}
