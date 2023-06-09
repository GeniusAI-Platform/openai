package completion

import (
	"context"
	"github.com/GeniusAI-Platform/openai"
	"github.com/GeniusAI-Platform/openai/client"
	"github.com/GeniusAI-Platform/openai/types/programming"
	"log"
	"os"
)

func ExampleProgrammingLanguageTranslator() {
	var code string = `
func add(a, b int) int {
	return a + b
}
`

	apiKey := os.Getenv("OPENAI_API_KEY")
	cli, err := client.New([]string{apiKey})
	if err != nil {
		log.Fatalln(err)
	}

	c := openai.NewCompletion(cli)
	resp, err := c.CreateCompletionFromPattern(context.Background(), ProgrammingLanguageTranslator(
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

func ExampleTextToCommand() {
	var text string = `
create nginx pod with kubectl and 5 replica
`

	apiKey := os.Getenv("OPENAI_API_KEY")
	cli, err := client.New([]string{apiKey})
	if err != nil {
		log.Fatalln(err)
	}

	c := openai.NewCompletion(cli)
	resp, err := c.CreateCompletionFromPattern(context.Background(), TextToCommand(
		text,
		0,
	))

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp.Choices[0].Text)
}

func ExampleProgrammingBugFixer() {
	var code string = `
func add(a, b int) string {
	return a + b
}
`

	apiKey := os.Getenv("OPENAI_API_KEY")
	cli, err := client.New([]string{apiKey})
	if err != nil {
		log.Fatalln(err)
	}

	c := openai.NewCompletion(cli)
	resp, err := c.CreateCompletionFromPattern(context.Background(), ProgrammingBugFixer(
		code,
		programming.Go,
		0,
	))

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp.Choices[0].Text)
}

func ExampleGrammarCorrection() {
	var text string = `
Helo w0rld! how are to you?
`

	apiKey := os.Getenv("OPENAI_API_KEY")
	cli, err := client.New([]string{apiKey})
	if err != nil {
		log.Fatalln(err)
	}

	c := openai.NewCompletion(cli)
	resp, err := c.CreateCompletionFromPattern(context.Background(), GrammarCorrection(
		text,
		0,
	))

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp.Choices[0].Text)
}

func ExampleProgrammingAlgorithmOptimizer() {
	var code string = `
func BubbleSort(array[] int)[]int {
   for i:=0; i< len(array)-1; i++ {
      for j:=0; j < len(array)-i-1; j++ {
         if (array[j] > array[j+1]) {
            array[j], array[j+1] = array[j+1], array[j]
         }
      }
   }
   return array
}
`

	apiKey := os.Getenv("OPENAI_API_KEY")
	cli, err := client.New([]string{apiKey})
	if err != nil {
		log.Fatalln(err)
	}

	c := openai.NewCompletion(cli)
	resp, err := c.CreateCompletionFromPattern(context.Background(), ProgrammingAlgorithmOptimizer(
		code,
		programming.Go,
		0,
	))

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp.Choices[0].Text)
}
