package errors

import (
	"errors"
	"github.com/GeniusAI-Platform/openai/entity"
)

var (
	ErrTooManyEmptyStreamMessages       = errors.New("stream has sent too many empty messages")
	ErrChatCompletionInvalidModel       = errors.New("this model is not supported with this method, please use CreateCompletion client method instead")
	ErrChatCompletionStreamNotSupported = errors.New("streaming is not supported with this method, please use CreateChatCompletionStream")
	ErrFailedToUnmarshalJSON            = errors.New("failed to unmarshal json response")
	ErrAPIKeyIsEmpty                    = errors.New("api key is empty")
	ErrFileIsInvalidFormat              = errors.New("file format is invalid, please create jsonl file and check training example: https://platform.openai.com/docs/guides/fine-tuning/prepare-training-data")
)

func New(httpCode int, providerCode string, message string, messageType string, param any) *entity.ErrorResponse {
	return &entity.ErrorResponse{
		HttpCode: httpCode,
		Err: &entity.ErrorPayload{
			Message: message,
			Type:    messageType,
			Param:   param,
			Code:    providerCode,
		},
	}
}
