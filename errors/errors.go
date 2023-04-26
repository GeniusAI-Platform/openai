package errors

import (
	"errors"
	"github.com/GoFarsi/openai/entity"
)

var (
	ErrTooManyEmptyStreamMessages       = errors.New("stream has sent too many empty messages")
	ErrChatCompletionInvalidModel       = errors.New("this model is not supported with this method, please use CreateCompletion client method instead")
	ErrChatCompletionStreamNotSupported = errors.New("streaming is not supported with this method, please use CreateChatCompletionStream")
	ErrFailedToUnmarshalJSON            = errors.New("failed to unmarshal json response")
	ErrAPIKeyIsEmpty                    = errors.New("api key is empty")
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
