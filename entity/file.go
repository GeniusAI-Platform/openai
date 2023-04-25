package entity

import "os"

type FileUploadRequest struct {
	// File Name of the JSON Lines file to be uploaded,
	// If the purpose is set to "fine-tune", each line is a JSON record with "prompt" and "completion" fields
	//representing your training examples: https://platform.openai.com/docs/guides/fine-tuning/prepare-training-data
	File *os.File `json:"file" validate:"required"`
	// Purpose The intended purpose of the uploaded documents, Use "fine-tune" for Fine-tuning. This allows us to validate the format of the uploaded file
	Purpose string `json:"purpose" validate:"required"`
}

type FileResponse struct {
	ID        string `json:"id"`
	Bytes     int    `json:"bytes"`
	FileName  string `json:"filename"`
	Object    string `json:"object"`
	Owner     string `json:"owner"`
	Purpose   string `json:"purpose"`
	CreatedAt int64  `json:"created_at"`
}

type FilesList struct {
	Files []FileResponse `json:"data"`
}
