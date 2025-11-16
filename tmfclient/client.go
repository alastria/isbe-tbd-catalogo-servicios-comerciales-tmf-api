package tmfclient

import (
	"bytes"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

// Config holds the configuration for the tmfclient package
type Config struct {
	// BaseURL of the remote TMForum server
	BaseURL string `json:"base_url" yaml:"base_url"`

	// Timeout in seconds for HTTP requests
	Timeout int `json:"timeout" yaml:"timeout"`
}

// Client is a client for the TMForum API.
type Client struct {
	config *Config
	client *http.Client
}

// NewClient creates a new client.
func NewClient(config *Config) *Client {
	return &Client{
		config: config,
		client: &http.Client{
			Timeout: time.Duration(config.Timeout) * time.Second,
		},
	}
}

// Get sends a GET request to the remote server.
func (c *Client) Get(path string, headers map[string]string) (*http.Response, error) {
	return c.do("GET", path, nil, headers)
}

// Post sends a POST request to the remote server.
func (c *Client) Post(path string, body []byte, headers map[string]string) (*http.Response, error) {
	return c.do("POST", path, body, headers)
}

// Patch sends a PATCH request to the remote server.
func (c *Client) Patch(path string, body []byte, headers map[string]string) (*http.Response, error) {
	return c.do("PATCH", path, body, headers)
}

// Delete sends a DELETE request to the remote server.
func (c *Client) Delete(path string, headers map[string]string) (*http.Response, error) {
	return c.do("DELETE", path, nil, headers)
}

func (c *Client) do(method, path string, body []byte, headers map[string]string) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", c.config.BaseURL, path)
	slog.Debug("sending", slog.String("method", method), "url", url)

	var req *http.Request
	var err error

	if body != nil {
		req, err = http.NewRequest(method, url, bytes.NewReader(body))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return nil, err
	}

	// TODO: send the authorization header
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return c.client.Do(req)
}
