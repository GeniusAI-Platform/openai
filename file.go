package openai

import (
	"bytes"
	"context"
	"fmt"
	"github.com/GeniusAI-Platform/openai/client"
	"github.com/GeniusAI-Platform/openai/entity"
	"github.com/GeniusAI-Platform/openai/errors"
	"github.com/GeniusAI-Platform/openai/types"
	"github.com/GeniusAI-Platform/openai/utils"
	"net/http"
	"path/filepath"
)

const (
	fileEndpoint                   = "/files"
	fileDynamicEndpoint            = "/files/%s"
	fileDynamicWithContentEndpoint = "/files/%s/content"
)

type File struct {
	client client.Transporter
}

// NewFile create File object to manage file (upload, retrieve, list)
func NewFile(client client.Transporter) *File {
	return &File{
		client: client,
	}
}

// ListFile Returns a list of files that belong to the user's organization
func (f *File) ListFile(ctx context.Context) (*entity.FilesListResponse, error) {
	resp, err := f.client.Get(ctx, &client.APIConfig{Path: fileEndpoint})
	if err != nil {
		return nil, err
	}

	return responseHandler[*entity.FilesListResponse](resp)
}

// UploadFile Upload a file that contains document(s) to be used across various endpoints/features. Currently, the size of all the files uploaded by one organization can be up to 1 GB. Please contact us if you need to increase the storage limit
func (f *File) UploadFile(ctx context.Context, req entity.FileUploadRequest) (*entity.FileResponse, error) {
	if err := f.client.GetValidator().Struct(req); err != nil {
		return nil, err
	}

	body := new(bytes.Buffer)
	fb := utils.NewFormBuilder(body)

	if filepath.Ext(req.File.Name()) != ".jsonl" {
		return nil, errors.ErrFileIsInvalidFormat
	}

	if err := fb.CreateFormFile("file", req.File); err != nil {
		return nil, err
	}

	if err := fb.WriteField("purpose", req.Purpose); err != nil {
		return nil, err
	}

	if err := fb.Close(); err != nil {
		return nil, err
	}

	resp, err := f.client.PostFile(ctx, &client.APIConfig{Path: fileEndpoint}, body, fb.FormDataContentType())
	if err != nil {
		return nil, err
	}

	return responseHandler[*entity.FileResponse](resp)
}

// RetrieveFile Returns information about a specific file or file content
func (f *File) RetrieveFile(ctx context.Context, fileID types.ID, content bool) (*entity.FileResponse, error) {
	if fileID.IsEmpty() {
		return nil, errors.New(http.StatusBadRequest, "", "fileID is empty", "", "")
	}

	path := fmt.Sprintf(fileDynamicEndpoint, fileID)

	if content {
		path = fmt.Sprintf(fileDynamicWithContentEndpoint, fileID)
	}

	resp, err := f.client.Get(ctx, &client.APIConfig{Path: path})
	if err != nil {
		return nil, err
	}

	return responseHandler[*entity.FileResponse](resp)
}
