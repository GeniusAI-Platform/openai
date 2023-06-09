package completion

import (
	"fmt"
	"github.com/GeniusAI-Platform/openai/entity"
	"github.com/GeniusAI-Platform/openai/models"
	"github.com/GeniusAI-Platform/openai/types/programming"
)

type CompletionPattern func() entity.CompletionRequest

// ProgrammingLanguageTranslator convert programming language code from language to language (go > python)
func ProgrammingLanguageTranslator(code string, languageFrom, languageTo programming.ProgrammingLanguage, maxTokens int) CompletionPattern {
	style := `
##### Translate this function from %s into %s
### %s
    
%s
    
### %s
`
	prompt := fmt.Sprintf(style, languageFrom, languageTo, languageFrom, code, languageTo)

	return func() entity.CompletionRequest {
		return requestBuilder(
			models.TEXT_DAVINCI_003,
			prompt,
			1.0,
			0,
			0.0,
			0.0,
			0.0,
			maxTokens,
			[]string{"###"}...,
		)
	}
}

// ProgrammingBugFixer find bug in your programming code with specific language
func ProgrammingBugFixer(code string, language programming.ProgrammingLanguage, maxTokens int) CompletionPattern {
	style := `
##### Fix bugs in the below function %s language
 
### Buggy %s

%s
    
### Fixed %s
`
	prompt := fmt.Sprintf(style, language, language, code, language)

	return func() entity.CompletionRequest {
		return requestBuilder(
			models.TEXT_DAVINCI_003,
			prompt,
			1.0,
			0,
			0.0,
			0.0,
			0.0,
			maxTokens,
			[]string{"###"}...,
		)
	}
}

// ProgrammingAlgorithmOptimizer improve performance your algorithm function
func ProgrammingAlgorithmOptimizer(code string, language programming.ProgrammingLanguage, maxTokens int) CompletionPattern {
	style := `
##### Improve performance in the below function
 
### Performance %s
%s
    
### Improved %s
`
	prompt := fmt.Sprintf(style, language, code, language)

	return func() entity.CompletionRequest {
		return requestBuilder(
			models.TEXT_DAVINCI_003,
			prompt,
			1.0,
			0.7,
			0.0,
			0.0,
			1,
			maxTokens,
			[]string{"###"}...,
		)
	}
}

// TextToCommand create command using explained text
func TextToCommand(text string, maxTokens int) CompletionPattern {
	style := `
Convert this text to a programmatic command:

%s
`
	prompt := fmt.Sprintf(style, text)

	return func() entity.CompletionRequest {
		return requestBuilder(
			models.TEXT_DAVINCI_003,
			prompt,
			1.0,
			0,
			0.0,
			0.2,
			0.0,
			maxTokens,
			[]string{`\n`}...,
		)
	}
}

// GrammarCorrection check your english text grammar and do correction
func GrammarCorrection(text string, maxTokens int) CompletionPattern {
	style := `
Correct this to standard English:

%s
`
	prompt := fmt.Sprintf(style, text)

	return func() entity.CompletionRequest {
		return requestBuilder(
			models.TEXT_DAVINCI_003,
			prompt,
			1.0,
			0,
			0.0,
			0.0,
			0.0,
			maxTokens,
		)
	}
}

func requestBuilder(model models.Completion, prompt any, topP, temperature, frequencyPenalty, presencePenalty float32, bestOf, maxTokens int, stop ...string) entity.CompletionRequest {
	req := entity.CompletionRequest{
		Model:            model,
		Prompt:           prompt,
		TopP:             topP,
		Temperature:      temperature,
		FrequencyPenalty: frequencyPenalty,
		PresencePenalty:  presencePenalty,
		BestOf:           bestOf,
		Stop:             stop,
	}

	if maxTokens != 0 {
		req.MaxTokens = maxTokens
	}

	return req
}
