package entity

import "github.com/GoFarsi/openai/models"

type EditsRequest struct {
	Model models.Edit `json:"model" validate:"required"`
	// Input The input text to use as a starting point for the edit
	Input string `json:"input,omitempty"`
	// Instruction The instruction that tells the model how to edit the prompt
	Instruction string `json:"instruction" validate:"required"`
	// N How many edits to generate for the input and instruction
	N int `json:"n,omitempty"`
	// Temperature What sampling temperature to use, between 0 and 2.
	//Higher values like 0.8 will make the output more random, while
	//lower values like 0.2 will make it more focused and deterministic
	Temperature float32 `json:"temperature,omitempty"`
	// TopP An alternative to sampling with temperature, called nucleus
	//sampling, where the model considers the results of the tokens with
	//top_p probability mass. So 0.1 means only the tokens comprising the
	//top 10% probability mass are considered
	TopP float32 `json:"top_p,omitempty"`
}

type EditsResponse struct {
	Object  string        `json:"object"`
	Created int64         `json:"created"`
	Usage   TokenUsage    `json:"usage"`
	Choices []EditsChoice `json:"choices"`
}

type EditsChoice struct {
	Text  string `json:"text"`
	Index int    `json:"index"`
}
