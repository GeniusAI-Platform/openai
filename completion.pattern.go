package openai

import (
	"fmt"
	"github.com/GoFarsi/openai/entity"
	"github.com/GoFarsi/openai/models"
	"github.com/GoFarsi/openai/utils"
)

type CompletionPattern func() entity.CompletionRequest

// ProgrammingLanguageTranslator convert programming language code from language to language (go > python)
func ProgrammingLanguageTranslator(code string, languageFrom, languageTo utils.ProgrammingLanguage) CompletionPattern {
	style := `
##### Translate this function from %s into %s
### %s
    
%s
    
### %s
`
	prompt := fmt.Sprintf(style, languageFrom, languageTo, languageFrom, code, languageTo)

	return func() entity.CompletionRequest {
		return entity.CompletionRequest{
			Model:       models.TEXT_DAVINCI_003,
			Prompt:      prompt,
			TopP:        1.0,
			Temperature: 0,
			Stop:        []string{"###"},
		}
	}
}

// ProgrammingBugFixer find bug in your programming code with specific language
func ProgrammingBugFixer(code string, language utils.ProgrammingLanguage) CompletionPattern {
	style := `
##### Fix bugs in the below function
 
### Buggy %s

%s
    
### Fixed %s
`
	prompt := fmt.Sprintf(style, language, code, language)

	return func() entity.CompletionRequest {
		return entity.CompletionRequest{
			Model:       models.TEXT_DAVINCI_003,
			Prompt:      prompt,
			TopP:        1.0,
			Temperature: 0,
			Stop:        []string{"###"},
		}
	}
}

// GrammarCorrection check your english text grammar and do correction
func GrammarCorrection(text string) CompletionPattern {
	style := `
Correct this to standard English:

%s
`
	prompt := fmt.Sprintf(style, text)

	return func() entity.CompletionRequest {
		return entity.CompletionRequest{
			Model:            models.TEXT_DAVINCI_003,
			Prompt:           prompt,
			TopP:             1.0,
			Temperature:      0,
			FrequencyPenalty: 0.0,
			PresencePenalty:  0.0,
		}
	}
}
