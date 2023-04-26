package main

import (
	"context"
	"github.com/GoFarsi/openai"
	"github.com/GoFarsi/openai/client"
	"github.com/GoFarsi/openai/entity"
	"github.com/GoFarsi/openai/models"
	"log"
)

func main() {
	cli, err := client.New("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")
	if err != nil {
		log.Fatalln(err)
	}

	c := openai.NewCompletion(cli)
	resp, err := c.CreateCompletion(context.Background(), entity.CompletionRequest{
		Model:  models.TEXT_DAVINCI_002,
		Prompt: "can you explain bubble sort algorithm?",
	})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp)
}
