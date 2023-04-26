package openai

import (
	"bytes"
	"context"
	"github.com/GoFarsi/openai/client"
	"github.com/GoFarsi/openai/entity"
	"github.com/GoFarsi/openai/utils"
	"net/http"
	"strconv"
)

const (
	createTranscriptionEndpoint = "/audio/transcriptions"
	createTranslationEndpoint   = "/audio/translations"
)

type Audio struct {
	client client.Transporter
}

func NewAudio(client client.Transporter) *Audio {
	return &Audio{
		client: client,
	}
}

// CreateTranscription Transcribes audio into the input language
func (a *Audio) CreateTranscription(ctx context.Context, req entity.AudioRequest) (*entity.AudioResponse, error) {
	if err := a.client.GetValidator().Struct(req); err != nil {
		return nil, err
	}

	body, contentType, err := a.createForm(req)
	if err != nil {
		return nil, err
	}

	resp, err := a.client.PostFile(ctx, &client.APIConfig{Path: createTranscriptionEndpoint}, body, contentType)
	if err != nil {
		return nil, err
	}

	response := new(entity.AudioResponse)
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

// CreateTranslation Translates audio into into English
func (a *Audio) CreateTranslation(ctx context.Context, req entity.AudioRequest) (*entity.AudioResponse, error) {
	if err := a.client.GetValidator().Struct(req); err != nil {
		return nil, err
	}

	body, contentType, err := a.createForm(req)
	if err != nil {
		return nil, err
	}

	resp, err := a.client.PostFile(ctx, &client.APIConfig{Path: createTranslationEndpoint}, body, contentType)
	if err != nil {
		return nil, err
	}

	response := new(entity.AudioResponse)
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

func (a *Audio) createForm(req entity.AudioRequest) (*bytes.Buffer, string, error) {
	body := new(bytes.Buffer)
	fb := utils.NewFormBuilder(body)

	if err := fb.CreateFormFile("file", req.File); err != nil {
		return nil, "", err
	}

	if err := fb.WriteField("model", req.Model.String()); err != nil {
		return nil, "", err
	}

	if err := fb.WriteField("prompt", req.Prompt); err != nil {
		return nil, "", err
	}

	if err := fb.WriteField("temperature", strconv.FormatFloat(float64(req.Temperature), 'E', -1, 64)); err != nil {
		return nil, "", err
	}

	if err := fb.WriteField("language", req.Language); err != nil {
		return nil, "", err
	}

	if err := fb.WriteField("response_format", req.ResponseFormat.String()); err != nil {
		return nil, "", err
	}

	if err := fb.Close(); err != nil {
		return nil, "", err
	}

	return body, fb.FormDataContentType(), nil
}
