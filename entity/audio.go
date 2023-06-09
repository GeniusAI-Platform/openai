package entity

import (
	"github.com/GeniusAI-Platform/openai/models"
	"os"
)

type AudioRequest struct {
	Model models.Audio `json:"model" validate:"required"`
	// File The audio file to transcribe, in one of these formats: mp3, mp4, mpeg, mpga, m4a, wav, or webm
	File *os.File `json:"file" validate:"required"`
	// Prompt An optional text to guide the model's style or continue a previous audio segment. The prompt should match the audio language
	Prompt string `json:"prompt,omitempty"`
	// Temperature The sampling temperature, between 0 and 1. Higher values like 0.8 will make the output more random, while lower values
	//like 0.2 will make it more focused and deterministic. If set to 0, the model will use log probability to automatically increase
	//the temperature until certain thresholds are hit
	Temperature float32 `json:"temperature,omitempty"`
	// Language The language of the input audio. Supplying the input language in ISO-639-1 format will improve accuracy and latency, learn more:
	// https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes
	Language string `json:"language,omitempty"`
	// ResponseFormat The format of the transcript output, in one of these options: json, text, srt, verbose_json, or vtt
	ResponseFormat AudioFormat `json:"responseFormat,omitempty"`
}

type AudioResponse struct {
	Text string `json:"text"`
}
