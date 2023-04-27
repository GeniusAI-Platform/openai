package openai

import (
	"context"
	"github.com/GoFarsi/openai/client"
	"github.com/GoFarsi/openai/entity"
)

const (
	moderationEndpoint = "/moderations"
)

type Moderation struct {
	client client.Transporter
}

func NewModeration(client client.Transporter) *Moderation {
	return &Moderation{
		client: client,
	}
}

// CreateModeration Classifies if text violates OpenAI's Content Policy
func (m *Moderation) CreateModeration(ctx context.Context, req entity.ModerationRequest) (*entity.ModerationResponse, error) {
	if err := m.client.GetValidator().Struct(req); err != nil {
		return nil, err
	}

	resp, err := m.client.Post(ctx, &client.APIConfig{Path: moderationEndpoint}, req)
	if err != nil {
		return nil, err
	}

	return responseHandler[*entity.ModerationResponse](resp)
}
