package entity

import "os"

type ImageRequest struct {
	// Prompt A text description of the desired image(s). The maximum length is 1000 characters
	Prompt string `json:"prompt" validate:"required"`
	// N The number of images to generate. Must be between 1 and 10
	N int `json:"n,omitempty"`
	// Size The size of the generated images. Must be one of ImageSize256x256, ImageSize512x512, or ImageSize1024x1024
	Size ImageSize `json:"size,omitempty"`
	// ResponseFormat The format in which the generated images are returned. Must be one of ImageResponseFormatURL or ImageResponseFormatB64JSON
	ResponseFormat ImageResponseFormat `json:"response_format,omitempty"`
	// User A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse
	User string `json:"user,omitempty"`
}

type ImageEditRequest struct {
	// Image The image to edit. Must be a valid PNG file, less than 4MB, and square. If mask is not provided,
	//image must have transparency, which will be used as the mask
	Image *os.File `json:"image" validate:"required"`
	// Mask An additional image whose fully transparent areas (e.g. where alpha is zero) indicate where image
	//should be edited. Must be a valid PNG file, less than 4MB, and have the same dimensions as image
	Mask *os.File `json:"mask,omitempty"`
	// Prompt A text description of the desired image(s). The maximum length is 1000 characters
	Prompt string `json:"prompt" validate:"required"`
	// N The number of images to generate. Must be between 1 and 10
	N int `json:"n,omitempty"`
	// Size The size of the generated images. Must be one of ImageSize256x256, ImageSize512x512, or ImageSize1024x1024
	Size ImageSize `json:"size,omitempty"`
	// ResponseFormat The format in which the generated images are returned. Must be one of ImageResponseFormatURL or ImageResponseFormatB64JSON
	ResponseFormat ImageResponseFormat `json:"response_format,omitempty"`
	// User A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse
	User string `json:"user,omitempty"`
}

type ImageVariationRequest struct {
	// Image The image to use as the basis for the variation(s). Must be a valid PNG file, less than 4MB, and square
	Image *os.File `json:"image" validate:"required"`
	// N The number of images to generate. Must be between 1 and 10
	N int `json:"n,omitempty"`
	// Size The size of the generated images. Must be one of ImageSize256x256, ImageSize512x512, or ImageSize1024x1024
	Size string `json:"size,omitempty"`
	// ResponseFormat The format in which the generated images are returned. Must be one of ImageResponseFormatURL or ImageResponseFormatB64JSON
	ResponseFormat ImageResponseFormat `json:"response_format,omitempty"`
	// User A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse
	User string `json:"user,omitempty"`
}

type ImageResponse struct {
	Created int64                    `json:"created,omitempty"`
	Data    []ImageResponseDataInner `json:"data,omitempty"`
}

type ImageResponseDataInner struct {
	URL     string `json:"url,omitempty"`
	B64JSON string `json:"b64_json,omitempty"`
}
