# chater
Using the OpenAI API in Golang

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

    // Prompt
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
