# OpenAi (community-maintained) [![Go Reference](https://pkg.go.dev/badge/github.com/GeniusAI-Platform/openai.svg)](https://pkg.go.dev/github.com/GeniusAI-Platform/openai)
Package openai provides a Go SDK for the OpenAI API.this package supports several models, including GPT-4, GPT-3.5, GPT-3, DALL-E, and audio
models. You can specify the desired model using the `Model` field in the request object.


## Feature

- ChatGPT (GPT-3, GPT-3.5, GPT-4)
- DALL·E 2
- Embedding
- Audio
- Fine-Tune
- File
- Moderations
- Completion Patterns
- Multiple API keys support

## Install ![Go Version](https://img.shields.io/badge/go%20version-%3E=1.19-61CFDD.svg?style=flat-square)

```shell
$ go get -u github.com/GeniusAI-Platform/openai
```

## Example Completion

```go
package main

import (
	"context"
	"github.com/GeniusAI-Platform/openai"
	"github.com/GeniusAI-Platform/openai/client"
	"github.com/GeniusAI-Platform/openai/entity"
	"github.com/GeniusAI-Platform/openai/models"
	"log"
	"os"
)

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	cli, err := client.New([]string{apiKey})
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

```

Example Completion Patterns

```go
package main

import (
	"context"
	"github.com/GeniusAI-Platform/openai"
	"github.com/GeniusAI-Platform/openai/client"
	"github.com/GeniusAI-Platform/openai/patterns/completion"
	"github.com/GeniusAI-Platform/openai/types/programming"
	"log"
	"os"
)

var code string = `
func add(a, b int) int {
	return a + b
}
`

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	cli, err := client.New([]string{apiKey})
	if err != nil {
		log.Fatalln(err)
	}

	c := openai.NewCompletion(cli)
	resp, err := c.CreateCompletionFromPattern(context.Background(), completion.ProgrammingLanguageTranslator(
		code,
		programming.Go,
		programming.Python,
		0,
	))

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp.Choices[0].Text)
}

```

See more details in [documentation](https://pkg.go.dev/github.com/GeniusAI-Platform/openai).

## TODO
- [ ] Stream Support
- [x] Moderation API
- [x] Example API
- [x] Fine-Tune API
- [x] File API
- [ ] Engine API
- [ ] Azure API Support
- [ ] Client, API Unit test

## Contributing

1. fork project in your GitHub account.
2. create new branch for new changes.
3. after change code, send Pull Request.
