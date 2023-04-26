/*
Package openai provides a Go SDK for the OpenAI API.this package supports several models, including GPT-4, GPT-3.5, GPT-3, DALL-E, and audio
models. You can specify the desired model using the `Model` field in the request object.

# Usage

To use this SDK, you will first need to obtain an API key from the OpenAI website.
You can then create a client object using the `New` function:

	client := client.New(apiKey)

The client object provides methods for making requests to the various endpoints of
the OpenAI API. For example, to generate text using the GPT-3.5 or GPT-4 model, you can use
the `CreateChatCompletion` method:

	c := openai.NewChat(cli)
	resp, err := c.CreateChatCompletion(context.Background(), entity.ChatRequest{
		Model: models.GPT35_TURBO,
		Messages: []entity.ChatMessage{
			{
				Role:    entity.USER,
				Content: "Hello",
			},
		},
	})

In addition to generating text and images, this package also supports fine-tuning
models and generating embeddings. For example, to fine-tune a GPT-3 model, you can
use the `CreateFineTune` method:

	c := openai.NewFineTune(cli)
	resp, err := c.CreateFineTune(context.Background(), entity.FineTuneRequest{})

For more information about the available methods and request/response objects, see
the documentation for the `Client` type and the various endpoint types.

# Authentication

Requests to the OpenAI API must include an API key in the `Authorization` header.
You can pass this key to the client constructor, or you can set the `OPENAI_API_KEY`
environment variable to automatically use it:

	os.Setenv("OPENAI_API_KEY", apiKey)
	client := client.New(os.GetEnv("OPENAI_API_KEY"))

# Concurrency

The client methods are safe to use concurrently from multiple goroutines.

# Errors

Any errors returned by the client methods will be of type `openai.Error`. This type
provides access to the raw HTTP response, as well as any JSON error response that
was returned by the API. For more information, see the documentation for the `Error`
type.

# Endpoint Types

The package defines types for each of the endpoints in the OpenAI API. These types
provide a convenient way to construct requests and parse responses for each endpoint.
For more information, see the documentation for each endpoint type.

# Examples

The `_examples` directory in the package source contains several examples of how to
use the SDK to perform various tasks with the OpenAI API. These examples can serve as
a starting point for your own usage of the SDK.
*/
package openai
