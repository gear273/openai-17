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