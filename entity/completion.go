package entity

import "github.com/GoFarsi/openai/models"

type CompletionRequest struct {
	Model models.Completion `json:"model" validate:"required"`
	// Prompt he prompt(s) to generate completions for, encoded as a string,
	// array of strings, array of tokens, or array of token arrays
	Prompt any `json:"prompt,omitempty" validate:"required"`
	// Suffix The suffix that comes after a completion of inserted text
	Suffix string `json:"suffix,omitempty"`
	// MaxTokens The maximum number of tokens to generate in the completion
	// The token count of your prompt plus max_tokens cannot exceed the model's
	// context length. Most models have a context length of 2048 tokens
	// (except for the newest models, which support 4096).
	MaxTokens int `json:"max_tokens,omitempty"`
	// Temperature What sampling temperature to use, between 0 and 2.
	//Higher values like 0.8 will make the output more random, while
	//lower values like 0.2 will make it more focused and deterministic
	Temperature float32 `json:"temperature,omitempty"`
	// TopP An alternative to sampling with temperature, called nucleus
	//sampling, where the model considers the results of the tokens with
	//top_p probability mass. So 0.1 means only the tokens comprising the
	//top 10% probability mass are considered
	TopP float32 `json:"top_p,omitempty"`
	// N How many completions to generate for each prompt
	N int `json:"n,omitempty"`
	// Stream Whether to stream back partial progress. If set,
	//tokens will be sent as data-only server-sent events as they become
	//available, with the stream terminated by a data: [DONE] message
	Stream bool `json:"stream,omitempty"`
	// LogProbs Include the log probabilities on the logprobs most likely
	//tokens, as well the chosen tokens. For example, if logprobs is 5,
	//the API will return a list of the 5 most likely tokens.
	//The API will always return the logprob of the sampled token,
	//so there may be up to logprobs+1 elements in the response
	LogProbs int `json:"logprobs,omitempty"`
	// Echo Echo back the prompt in addition to the completion
	Echo bool `json:"echo,omitempty"`
	// Stop Up to 4 sequences where the API will stop generating further tokens.
	//The returned text will not contain the stop sequence
	Stop []string `json:"stop,omitempty"`
	// PresencePenalty Number between -2.0 and 2.0. Positive values penalize
	//new tokens based on whether they appear in the text so far,
	//increasing the model's likelihood to talk about new topics
	PresencePenalty float32 `json:"presence_penalty,omitempty"`
	// FrequencyPenalty Number between -2.0 and 2.0. Positive values penalize new tokens
	//based on their existing frequency in the text so far,
	//decreasing the model's likelihood to repeat the same line verbatim
	FrequencyPenalty float32 `json:"frequency_penalty,omitempty"`
	// BestOf Generates best_of completions server-side and returns the "best"
	//(the one with the highest log probability per token). Results
	//cannot be streamed
	BestOf int `json:"best_of,omitempty"`
	// LogitBias Modify the likelihood of specified tokens appearing in the completion
	LogitBias map[string]int `json:"logit_bias,omitempty"`
	// User A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse.
	User string `json:"user,omitempty"`
}

type CompletionResponse struct {
	ID      string             `json:"id"`
	Object  string             `json:"object"`
	Created int64              `json:"created"`
	Model   string             `json:"model"`
	Choices []CompletionChoice `json:"choices"`
	Usage   TokenUsage         `json:"usage"`
}

type CompletionChoice struct {
	Text         string        `json:"text"`
	Index        int           `json:"index"`
	FinishReason string        `json:"finish_reason"`
	LogProbs     LogprobResult `json:"logprobs"`
}

type LogprobResult struct {
	Tokens        []string             `json:"tokens"`
	TokenLogprobs []float32            `json:"token_logprobs"`
	TopLogprobs   []map[string]float32 `json:"top_logprobs"`
	TextOffset    []int                `json:"text_offset"`
}
