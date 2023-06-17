package openai

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const chatCompletionsSuffix = "/chat/completions"

const (
	ChatMessageRoleUser      = "user"
	ChatMessageRoleSystem    = "system"
	ChatMessageRoleAssistant = "assistant"
	ChatMessageRoleFunction  = "function"
)

type ChatCompletionMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatCompletionRequest struct {
	Model            string                  `json:"model"`
	Messages         []ChatCompletionMessage `json:"messages"`
	MaxTokens        int                     `json:"max_tokens,omitempty"`
	Temperature      float32                 `json:"temperature,omitempty"`
	TopP             float32                 `json:"top_p,omitempty"`
	N                int                     `json:"n,omitempty"`
	Stream           bool                    `json:"stream,omitempty"`
	Stop             []string                `json:"stop,omitempty"`
	PresencePenalty  float32                 `json:"presence_penalty,omitempty"`
	FrequencyPenalty float32                 `json:"frequency_penalty,omitempty"`
}

type FinishReason string

const (
	FinishReasonStop FinishReason = "stop"
	FinishReasonNull FinishReason = "null"
)

type ChatCompletionChoice struct {
	Index        int                   `json:"index"`
	Message      ChatCompletionMessage `json:"message"`
	FinishReason FinishReason          `json:"finish_reason"`
}

type ChatCompletionResponse struct {
	ID      string                 `json:"id"`
	Object  string                 `json:"object"`
	Created int64                  `json:"created"`
	Model   string                 `json:"model"`
	Choices []ChatCompletionChoice `json:"choices"`
	Usage   Usage                  `json:"usage"`
}

func (c *Client) CreateChatCompletion(request ChatCompletionRequest) (response ChatCompletionResponse, err error) {
	url := c.config.apiBase + chatCompletionsSuffix
	by, _ := json.Marshal(request)
	payload := bytes.NewReader(by)

	req, err := http.NewRequest(http.MethodPost, url, payload)
	if err != nil {
		return
	}

	err = c.request(req, &response)
	return
}

func (c *ChatCompletionRequest) AddUserMessage(content string) {
	c.Messages = append(c.Messages, ChatCompletionMessage{
		Role:    ChatMessageRoleUser,
		Content: content,
	})
}

func (c *ChatCompletionRequest) AddAssistantMessage(content string) {
	c.Messages = append(c.Messages, ChatCompletionMessage{
		Role:    ChatMessageRoleAssistant,
		Content: content,
	})
}

func (c *ChatCompletionRequest) ClearMessage(content string) {
	c.Messages = []ChatCompletionMessage{}
}

func NewChatCompletion() (ChatCompletionRequest, error) {
	return ChatCompletionRequest{
		Model:    GPT3Dot5Turbo,
		Messages: []ChatCompletionMessage{},
	}, nil
}
