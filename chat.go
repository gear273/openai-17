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

type ChatCompletionFinishReason string

const (
	FinishReasonStop         ChatCompletionFinishReason = "stop"
	FinishReasonLength       ChatCompletionFinishReason = "length"
	FinishReasonFunctionCall ChatCompletionFinishReason = "function_call"
)

type ChatCompletionMessage struct {
	Role         string                      `json:"role"`
	Content      string                      `json:"content"`
	Name         string                      `json:"name,omitempty"`
	FunctionCall *ChatCompletionFunctionCall `json:"function_call,omitempty"`
}

type ChatCompletionFunctionCall struct {
	Name      string `json:"name,omitempty"`
	Arguments string `json:"arguments,omitempty"`
}

type ChatCompletionRequest struct {
	Model            string                    `json:"model"`
	Messages         []ChatCompletionMessage   `json:"messages"`
	Functions        []*ChatCompletionFunction `json:"functions,omitempty"`
	FunctionCall     string                    `json:"function_call,omitempty"`
	MaxTokens        int                       `json:"max_tokens,omitempty"`
	Temperature      float32                   `json:"temperature,omitempty"`
	TopP             float32                   `json:"top_p,omitempty"`
	N                int                       `json:"n,omitempty"`
	Stream           bool                      `json:"stream,omitempty"`
	Stop             []string                  `json:"stop,omitempty"`
	PresencePenalty  float32                   `json:"presence_penalty,omitempty"`
	FrequencyPenalty float32                   `json:"frequency_penalty,omitempty"`
}

// https://platform.openai.com/docs/guides/gpt/function-calling
type ChatCompletionFunction struct {
	Name        string                        `json:"name"`
	Description string                        `json:"description,omitempty"`
	Parameters  *ChatCompletionFunctionParams `json:"parameters,omitempty"`
}

type ChatCompletionFunctionParams struct {
	Type       JSONSchemaType               `json:"type"`
	Properties map[string]*JSONSchemaDefine `json:"properties,omitempty"`
	Required   []string                     `json:"required,omitempty"`
}

type ChatCompletionChoice struct {
	Index        int                        `json:"index"`
	Message      ChatCompletionMessage      `json:"message"`
	FinishReason ChatCompletionFinishReason `json:"finish_reason"`
}

type ChatCompletionResponse struct {
	ID      string                 `json:"id"`
	Object  string                 `json:"object"`
	Created int64                  `json:"created"`
	Model   string                 `json:"model"`
	Choices []ChatCompletionChoice `json:"choices"`
	Usage   Usage                  `json:"usage"`
}

func (c *Client) CreateChatCompletion(cr *ChatCompletionRequest) (cresp ChatCompletionResponse, err error) {
	url := c.Config.apiBase + chatCompletionsSuffix
	by, _ := json.Marshal(cr)
	payload := bytes.NewReader((by))

	req, err := http.NewRequest(http.MethodPost, url, payload)
	if err != nil {
		return
	}

	err = c.request(req, &cresp)
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

func (c *ChatCompletionRequest) AddFunctionMessage(functionName string, content string) {
	c.Messages = append(c.Messages, ChatCompletionMessage{
		Role:    ChatMessageRoleFunction,
		Name:    functionName,
		Content: content,
	})
}

func (c *ChatCompletionRequest) ClearMessage() {
	c.Messages = []ChatCompletionMessage{}
}

func (c *ChatCompletionRequest) SetModel(model string) {
	c.Model = model
}

func (c *ChatCompletionRequest) AddFunction(v any) {
	if v == nil {
		return
	}

	switch v := v.(type) {
	case *ChatCompletionFunction:
		c.Functions = append(c.Functions, v)
	case []*ChatCompletionFunction:
		c.Functions = append(c.Functions, v...)
	default:
		return
	}
}

func (c *ChatCompletionRequest) Close() {
	c = nil
}

func NewChatCompletion() *ChatCompletionRequest {
	return &ChatCompletionRequest{
		Model:    GPT3Dot5Turbo,
		Messages: []ChatCompletionMessage{},
	}
}
