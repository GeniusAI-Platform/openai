package openai

import (
	"bytes"
	"context"
	"github.com/GoFarsi/openai/client"
	"github.com/GoFarsi/openai/entity"
	"github.com/GoFarsi/openai/utils"
	"strconv"
)

const (
	createImageEndpoint          = "/images/generations"
	editImageEndpoint            = "/images/edits"
	createImageVariationEndpoint = "/images/variations"
)

type Image struct {
	client client.Transporter
}

// NewImage create Image object to create, edit image or image variation using DALL·E 2
func NewImage(client client.Transporter) *Image {
	return &Image{
		client: client,
	}
}

// CreateImage Creates an image given a prompt
func (i *Image) CreateImage(ctx context.Context, req entity.ImageRequest) (*entity.ImageResponse, error) {
	if err := i.client.GetValidator().Struct(req); err != nil {
		return nil, err
	}

	resp, err := i.client.Post(ctx, &client.APIConfig{Path: createImageEndpoint}, req)
	if err != nil {
		return nil, err
	}

	return responseHandler[*entity.ImageResponse](resp)
}

// ImageEdit Creates an edited or extended image given an original image and a prompt
func (i *Image) ImageEdit(ctx context.Context, req entity.ImageEditRequest) (*entity.ImageResponse, error) {
	if err := i.client.GetValidator().Struct(req); err != nil {
		return nil, err
	}

	body := new(bytes.Buffer)
	fb := utils.NewFormBuilder(body)

	if err := fb.CreateFormFile("image", req.Image); err != nil {
		return nil, err
	}

	if req.Mask != nil {
		if err := fb.CreateFormFile("mask", req.Mask); err != nil {
			return nil, err
		}
	}

	if err := fb.WriteField("prompt", req.Prompt); err != nil {
		return nil, err
	}

	if req.N == 0 {
		req.N = 1
	}

	if err := fb.WriteField("n", strconv.Itoa(req.N)); err != nil {
		return nil, err
	}

	if err := fb.WriteField("size", req.Size.String()); err != nil {
		return nil, err
	}

	if err := fb.WriteField("response_format", req.ResponseFormat.String()); err != nil {
		return nil, err
	}

	if err := fb.WriteField("user", req.User); err != nil {
		return nil, err
	}

	if err := fb.Close(); err != nil {
		return nil, err
	}

	resp, err := i.client.PostFile(ctx, &client.APIConfig{Path: editImageEndpoint}, body, fb.FormDataContentType())
	if err != nil {
		return nil, err
	}

	return responseHandler[*entity.ImageResponse](resp)
}

// CreateImageVariation Creates a variation of a given image
func (i *Image) CreateImageVariation(ctx context.Context, req entity.ImageVariationRequest) (*entity.ImageResponse, error) {
	if err := i.client.GetValidator().Struct(req); err != nil {
		return nil, err
	}

	body := new(bytes.Buffer)
	fb := utils.NewFormBuilder(body)

	if err := fb.CreateFormFile("image", req.Image); err != nil {
		return nil, err
	}

	if req.N == 0 {
		req.N = 1
	}

	if err := fb.WriteField("n", strconv.Itoa(req.N)); err != nil {
		return nil, err
	}

	if err := fb.WriteField("size", req.Size.String()); err != nil {
		return nil, err
	}

	if err := fb.WriteField("response_format", req.ResponseFormat.String()); err != nil {
		return nil, err
	}

	if err := fb.WriteField("user", req.User); err != nil {
		return nil, err
	}

	if err := fb.Close(); err != nil {
		return nil, err
	}

	resp, err := i.client.PostFile(ctx, &client.APIConfig{Path: createImageVariationEndpoint}, body, fb.FormDataContentType())
	if err != nil {
		return nil, err
	}

	return responseHandler[*entity.ImageResponse](resp)
}
