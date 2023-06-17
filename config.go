package openai

import (
	"net/http"
)

const openaiAPIBaseV1 = "https://api.openai.com/v1"

type Config struct {
	apiKey  string
	apiBase string

	HTTPClient *http.Client
}

func defaultConfig(APIKey string) Config {
	return Config{
		apiKey:  APIKey,
		apiBase: openaiAPIBaseV1,

		HTTPClient: &http.Client{},
	}
}
