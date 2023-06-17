package openai

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	Config Config
}

func NewClient(apiKey string) (*Client, error) {
	config := defaultConfig(apiKey)
	return &Client{
		Config: config,
	}, nil
}

func (c *Client) request(request *http.Request, v any) error {
	c.setHeaders(request)
	response, err := c.Config.HTTPClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if isFailureStatusCode(response) {
		return handleErrorResponse(response)
	}

	return decodeResponse(response.Body, v)
}

func (c *Client) setHeaders(req *http.Request) {
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Config.apiKey))
}

func isFailureStatusCode(resp *http.Response) bool {
	return resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest
}

func decodeResponse(body io.Reader, v any) error {
	if v == nil {
		return nil
	}

	if result, ok := v.(*string); ok {
		b, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		*result = string(b)
		return nil
	}

	return json.NewDecoder(body).Decode(v)
}

// unfinish
func handleErrorResponse(resp *http.Response) error {
	b, _ := io.ReadAll(resp.Body)
	return fmt.Errorf("%v\n", string(b))
}
