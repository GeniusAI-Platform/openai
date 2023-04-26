package openai

import (
	"context"
	"github.com/GoFarsi/openai/client"
	"github.com/GoFarsi/openai/entity"
	"github.com/GoFarsi/openai/models"
	"github.com/GoFarsi/openai/utils"
	"log"
	"os"
)

func ExampleNewChat() {
	cli, err := client.New("OPENAI_API_KEY")
	if err != nil {
		log.Fatalln(err)
	}

	c := NewChat(cli)
	resp, err := c.CreateChatCompletion(context.Background(), entity.ChatRequest{
		Model: models.GPT35_TURBO,
		Messages: []entity.ChatMessage{
			{
				Role:    entity.USER,
				Content: "Hello!!",
			},
		},
	})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp)
}

func ExampleNewCompletion() {
	cli, err := client.New("OPENAI_API_KEY")
	if err != nil {
		log.Fatalln(err)
	}

	c := NewCompletion(cli)
	resp, err := c.CreateCompletion(context.Background(), entity.CompletionRequest{
		Model:  models.TEXT_DAVINCI_002,
		Prompt: "Golang history",
	})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp)
}

func ExampleNewImage() {
	cli, err := client.New("OPENAI_API_KEY")
	if err != nil {
		log.Fatalln(err)
	}

	c := NewImage(cli)
	resp, err := c.CreateImage(context.Background(), entity.ImageRequest{
		Prompt: "Create a gopher baby",
	})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp)
}

func ExampleNewAudio() {
	cli, err := client.New("OPENAI_API_KEY")
	if err != nil {
		log.Fatalln(err)
	}

	f, err := os.Open("./testdata/file.mp3")
	if err != nil {
		log.Fatalln(err)
	}

	c := NewAudio(cli)
	resp, err := c.CreateTranscription(context.Background(), entity.AudioRequest{
		Model:    models.WHISPER_1,
		File:     f,
		Language: "fa",
	})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp)
}

func ExampleNewEmbedding() {
	cli, err := client.New("OPENAI_API_KEY")
	if err != nil {
		log.Fatalln(err)
	}

	c := NewEmbedding(cli)
	resp, err := c.CreateEmbedding(context.Background(), entity.EmbeddingRequest{
		Model: models.TEXT_EMBEDDING_ADA_002,
		Input: []string{"example input"},
	})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp)
}

func ExampleNewFile() {
	cli, err := client.New("OPENAI_API_KEY")
	if err != nil {
		log.Fatalln(err)
	}

	f, err := os.Open("./testdata/file.jsonl")
	if err != nil {
		log.Fatalln(err)
	}

	c := NewFile(cli)
	resp, err := c.UploadFile(context.Background(), entity.FileUploadRequest{
		File:    f,
		Purpose: "fine-tune",
	})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp)
}

func ExampleCompletion_CreateCompletionFromPattern() {
	var code string = `
func add(a, b int) int {
	return a + b
}
`

	cli, err := client.New("OPENAI_API_KEY")
	if err != nil {
		log.Fatalln(err)
	}

	c := NewCompletion(cli)
	resp, err := c.CreateCompletionFromPattern(context.Background(), ProgrammingLanguageTranslator(
		code,
		utils.Go,
		utils.Python,
	))

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp.Choices[0].Text)
}
