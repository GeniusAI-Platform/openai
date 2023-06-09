package openai

import (
	"context"
	"github.com/GeniusAI-Platform/openai/client"
	"github.com/GeniusAI-Platform/openai/entity"
)

const (
	chatCompletionEndpoint = "/chat/completions"
)

type ChatCompletion struct {
	client client.Transporter
}

// NewChat create chat completion object to create chat with chatgpt
func NewChat(client client.Transporter) *ChatCompletion {
	return &ChatCompletion{
		client: client,
	}
}

// CreateChatCompletion Creates a completion for the provided prompt and parameters
func (c *ChatCompletion) CreateChatCompletion(ctx context.Context, req entity.ChatRequest) (*entity.ChatResponse, error) {
	if err := c.client.GetValidator().Struct(req); err != nil {
		return nil, err
	}

	resp, err := c.client.Post(ctx, &client.APIConfig{Path: chatCompletionEndpoint}, req)
	if err != nil {
		return nil, err
	}

	return responseHandler[*entity.ChatResponse](resp)
}
