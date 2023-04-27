package openai

import (
	"context"
	"github.com/GoFarsi/openai/client"
	"github.com/GoFarsi/openai/entity"
)

const (
	createEmbeddingEndpoint = "/embeddings"
)

type Embedding struct {
	client client.Transporter
}

// NewEmbedding create embedding object to create embeddings
func NewEmbedding(client client.Transporter) *Embedding {
	return &Embedding{
		client: client,
	}
}

// CreateEmbedding Creates an embedding vector representing the input text
func (e *Embedding) CreateEmbedding(ctx context.Context, req entity.EmbeddingRequest) (*entity.EmbeddingResponse, error) {
	if err := e.client.GetValidator().Struct(req); err != nil {
		return nil, err
	}

	resp, err := e.client.Post(ctx, &client.APIConfig{Path: createEmbeddingEndpoint}, req)
	if err != nil {
		return nil, err
	}

	return responseHandler[*entity.EmbeddingResponse](resp)
}
