package openai

import (
	"context"
	"fmt"
	"github.com/GoFarsi/openai/client"
	"github.com/GoFarsi/openai/entity"
	"github.com/GoFarsi/openai/errors"
	"github.com/GoFarsi/openai/types"
	"net/http"
)

const (
	fineTuneEndpoint            = "/fine-tunes"
	fineTuneDynamicEndpoint     = "/fine-tunes/%s"
	fineTuneCancelEndpoint      = "/fine-tunes/%s/cancel"
	fineTuneEventEndpoint       = "/fine-tunes/%s/events"
	deleteFineTuneModelEndpoint = "/models/%s"
)

type FineTune struct {
	client client.Transporter
}

func NewFineTune(client client.Transporter) *FineTune {
	return &FineTune{
		client: client,
	}
}

// CreateFineTune Creates a job that fine-tunes a specified model from a given dataset
func (f *FineTune) CreateFineTune(ctx context.Context, req entity.FineTuneRequest) (*entity.FineTuneResponse, error) {
	if err := f.client.GetValidator().Struct(req); err != nil {
		return nil, err
	}

	resp, err := f.client.Post(ctx, &client.APIConfig{Path: fineTuneEndpoint}, req)
	if err != nil {
		return nil, err
	}

	return responseHandler[*entity.FineTuneResponse](resp)
}

// ListFineTunes List your organization's fine-tuning jobs
func (f *FineTune) ListFineTunes(ctx context.Context) (*entity.FineTuneList, error) {
	resp, err := f.client.Get(ctx, &client.APIConfig{Path: fineTuneEndpoint})
	if err != nil {
		return nil, err
	}

	return responseHandler[*entity.FineTuneList](resp)
}

// RetrieveFineTune Gets info about the fine-tune job
func (f *FineTune) RetrieveFineTune(ctx context.Context, fineTuneID types.ID) (*entity.FineTuneResponse, error) {
	if fineTuneID.IsEmpty() {
		return nil, errors.New(http.StatusBadRequest, "", "fineTuneID is empty", "", "")
	}

	resp, err := f.client.Get(ctx, &client.APIConfig{Path: fmt.Sprintf(fineTuneDynamicEndpoint, fineTuneID)})
	if err != nil {
		return nil, err
	}

	return responseHandler[*entity.FineTuneResponse](resp)
}

// CancelFineTune Immediately cancel a fine-tune job
func (f *FineTune) CancelFineTune(ctx context.Context, fineTuneID types.ID) (*entity.FineTuneResponse, error) {
	if fineTuneID.IsEmpty() {
		return nil, errors.New(http.StatusBadRequest, "", "fineTuneID is empty", "", "")
	}

	resp, err := f.client.Post(ctx, &client.APIConfig{Path: fmt.Sprintf(fineTuneCancelEndpoint, fineTuneID)}, nil)
	if err != nil {
		return nil, err
	}

	return responseHandler[*entity.FineTuneResponse](resp)
}

// ListFineTuneEvent Get fine-grained status updates for a fine-tune job
func (f *FineTune) ListFineTuneEvent(ctx context.Context, fineTuneID types.ID) (*entity.FineTuneEventList, error) {
	if fineTuneID.IsEmpty() {
		return nil, errors.New(http.StatusBadRequest, "", "fineTuneID is empty", "", "")
	}

	resp, err := f.client.Get(ctx, &client.APIConfig{Path: fmt.Sprintf(fineTuneEventEndpoint, fineTuneID)})
	if err != nil {
		return nil, err
	}

	return responseHandler[*entity.FineTuneEventList](resp)
}

// DeleteFineTuneModel Delete a fine-tuned model. You must have the Owner role in your organization
func (f *FineTune) DeleteFineTuneModel(ctx context.Context, model string) (*entity.FineTuneDeleteResponse, error) {
	if len(model) == 0 {
		return nil, errors.New(http.StatusBadRequest, "", "model is empty", "", "")
	}

	resp, err := f.client.Delete(ctx, &client.APIConfig{Path: fmt.Sprintf(deleteFineTuneModelEndpoint, model)})
	if err != nil {
		return nil, err
	}

	return responseHandler[*entity.FineTuneDeleteResponse](resp)
}
