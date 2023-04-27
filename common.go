package openai

import (
	"github.com/GoFarsi/openai/client"
	"github.com/GoFarsi/openai/entity"
	"github.com/GoFarsi/openai/errors"
	"net/http"
	"reflect"
)

func responseHandler[T any](resp *client.Response) (response T, err error) {
	errResp := new(entity.ErrorResponse)
	m, ok := reflect.New(reflect.TypeOf(response).Elem()).Interface().(T)
	if !ok {
		return response, errors.New(http.StatusInternalServerError, "", "response type is invalid", "", "")
	}
	if resp.GetHttpResponse().StatusCode != http.StatusOK {
		if err = resp.GetJSON(errResp); err != nil {
			return response, err
		}
		errResp.HttpCode = resp.GetHttpResponse().StatusCode
		return response, errResp
	}

	if err = resp.GetJSON(m); err != nil {
		return response, err
	}

	return m, nil
}
