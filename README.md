# OpenAi (community-maintained) [![Go Reference](https://pkg.go.dev/badge/github.com/GoFarsi/openai.svg)](https://pkg.go.dev/github.com/GoFarsi/openai)
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

## Install ![Go Version](https://img.shields.io/badge/go%20version-%3E=1.19-61CFDD.svg?style=flat-square)

```shell
$ go get -u github.com/GoFarsi/openai
```

## Example Completion

```go
package main

import (
	"context"
	"github.com/GoFarsi/openai"
	"github.com/GoFarsi/openai/client"
	"github.com/GoFarsi/openai/entity"
	"github.com/GoFarsi/openai/models"
	"log"
	"os"
)

func main() {
	cli, err := client.New(os.Getenv("OPENAI_API_KEY"))
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
	"github.com/GoFarsi/openai"
	"github.com/GoFarsi/openai/client"
	"github.com/GoFarsi/openai/utils"
	"log"
)

var code string = `
func add(a, b int) int {
	return a + b
}
`

func main() {
	cli, err := client.New(os.Getenv("OPENAI_API_KEY"))
	if err != nil {
		log.Fatalln(err)
	}

	c := openai.NewCompletion(cli)
	resp, err := c.CreateCompletionFromPattern(context.Background(), openai.ProgrammingLanguageTranslator(
		code,
		utils.Go,
		utils.Python,
	))

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp.Choices[0].Text)
}

```

See more details in [documentation](https://pkg.go.dev/github.com/GoFarsi/openai).

## TODO
- [ ] Stream Support
- [ ] Moderation API
- [x] Example API
- [ ] Fine-Tune API
- [x] File API
- [ ] Engine API
- [ ] Azure API Support
- [ ] Client, API Unit test

## Contributing

1. fork project in your GitHub account.
2. create new branch for new changes.
3. after change code, send Pull Request.
