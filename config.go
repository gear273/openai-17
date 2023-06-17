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

func (c *Config) SetAPIKey(key string) {
	c.apiKey = key
}

func (c *Config) SetAPIBase(base string) {
	c.apiBase = base
}

func defaultConfig(key string) Config {
	return Config{
		apiKey:  key,
		apiBase: openaiAPIBaseV1,

		HTTPClient: &http.Client{},
	}
}
