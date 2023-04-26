package openai

import (
	"context"
	"github.com/GoFarsi/openai/client"
	"github.com/GoFarsi/openai/entity"
	"net/http"
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

	response := new(entity.ChatResponse)
	errResp := new(entity.ErrorResponse)

	if resp.GetHttpResponse().StatusCode != http.StatusOK {
		if err = resp.GetJSON(errResp); err != nil {
			return nil, err
		}
		errResp.HttpCode = resp.GetHttpResponse().StatusCode
		return nil, errResp
	}

	if err = resp.GetJSON(response); err != nil {
		return nil, err
	}

	return response, nil
}
