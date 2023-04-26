package openai

import (
	"context"
	"github.com/GoFarsi/openai/client"
	"github.com/GoFarsi/openai/entity"
	"github.com/GoFarsi/openai/patterns/completion"
	"net/http"
)

const (
	createCompletionEndpoint = "/completions"
)

type Completion struct {
	client client.Transporter
}

// NewCompletion create Completion object to text completion using davinci
func NewCompletion(client client.Transporter) *Completion {
	return &Completion{
		client: client,
	}
}

// CreateCompletion Creates a job that fine-tunes a specified model from a given dataset.
// Response includes details of the enqueued job including job status and the name of the fine-tuned models once complete.
func (c *Completion) CreateCompletion(ctx context.Context, req entity.CompletionRequest) (*entity.CompletionResponse, error) {
	return c.request(ctx, req)
}

// CreateCompletionFromPattern create a completion using specific patterns
func (c *Completion) CreateCompletionFromPattern(ctx context.Context, pattern completion.CompletionPattern) (*entity.CompletionResponse, error) {
	return c.request(ctx, pattern())
}

func (c *Completion) request(ctx context.Context, req entity.CompletionRequest) (*entity.CompletionResponse, error) {
	if err := c.client.GetValidator().Struct(req); err != nil {
		return nil, err
	}

	resp, err := c.client.Post(ctx, &client.APIConfig{Path: createCompletionEndpoint}, req)
	if err != nil {
		return nil, err
	}

	response := new(entity.CompletionResponse)
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
