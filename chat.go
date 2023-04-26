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

func NewChat(client client.Transporter) *ChatCompletion {
	return &ChatCompletion{
		client: client,
	}
}

// CreateChatCompletion start chat with chatgpt
func (c *ChatCompletion) CreateChatCompletion(ctx context.Context, req entity.ChatRequest) (*entity.ChatResponse, error) {
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
