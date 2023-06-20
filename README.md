## Using the OpenAI API in Golang

## Chat Completion
Use the same format for the messages as you would for the official [OpenAI API](https://platform.openai.com/docs/guides/gpt/chat-completions-api).

## Example
```bash
go get -u github.com/Sricor/openai
```

```go
package main

import (
	"fmt"

	"github.com/Sricor/openai"
)

func main() {
	client := openai.NewClient("your api key")
	chat := openai.NewChatCompletion()

	chat.AddUserMessage("Hello world!")

	response, err := client.CreateChatCompletion(chat)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	// Support context
	// chat.AddAssistantMessage(response.Choices[0].Message.Content)

	fmt.Println(response.Choices[0].Message.Content)
}

```

## Function Call Example

```go
package main

import (
	"fmt"

	"github.com/Sricor/openai"
)

func main() {
	client := openai.NewClient("")
	chat := openai.NewChatCompletion()

	// https://platform.openai.com/docs/guides/gpt/function-calling
	functions := &openai.ChatCompletionFunction{
		Name:        "get_current_weather",
		Description: "Get the current weather in a given location",

		// JSON Schema
		Parameters: &openai.ChatCompletionFunctionParams{
			Type: openai.JSONSchemaTypeObject,

			Properties: map[string]*openai.JSONSchemaDefine{
				"location": &openai.JSONSchemaDefine{
					Type:        openai.JSONSchemaTypeString,
					Description: "The city and state, e.g. San Francisco, CA",
				},
				"unit": &openai.JSONSchemaDefine{
					Type: openai.JSONSchemaTypeString,
					Enum: []string{"celsius", "fahrenheit"},
				},
			},
			Required: []string{"location"},
		},
	}
	chat.SetModel(openai.GPT3Dot5Turbo0613)                       // Set model to gpt-3.5-turbo-0613
	chat.AddFunction(functions) // Add functions to request

	chat.AddUserMessage("What's the weather like in Boston?")

	response, err := client.CreateChatCompletion(chat)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}
	fmt.Println(response.Choices[0].FinishReason)

	if response.Choices[0].Message.FunctionCall != nil {
		fmt.Println(response.Choices[0].Message.FunctionCall.Name)
		fmt.Println(response.Choices[0].Message.FunctionCall.Arguments)
	}

}

```