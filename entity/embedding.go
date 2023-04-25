package entity

import "github.com/GoFarsi/openai/models"

// EmbeddingRequest is the input to a Create embeddings request.
type EmbeddingRequest struct {
	Model models.Embedding `json:"model" validate:"required"`
	// Input is a slice of strings for which you want to generate an EmbeddingData vector.
	// Each input must not exceed 2048 tokens in length.
	// OpenAPI suggests replacing newlines (\n) in your input with a single space, as they
	// have observed inferior results when newlines are present.
	// E.g.
	//	"The food was delicious and the waiter..."
	Input []string `json:"input" validate:"required"`
	// User A unique identifier representing your end-user, which will help OpenAI to monitor and detect abuse.
	User string `json:"user"`
}

// EmbeddingResponse is the response from a Create embeddings request.
type EmbeddingResponse struct {
	Model  models.Embedding `json:"model"`
	Object string           `json:"object"`
	Data   []EmbeddingData  `json:"data"`
	Usage  TokenUsage       `json:"usage"`
}

type EmbeddingData struct {
	Object    string    `json:"object"`
	Embedding []float32 `json:"embedding"`
	Index     int       `json:"index"`
}
