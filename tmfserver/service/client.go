package service

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hesusruiz/isbetmf/internal/errl"
	"github.com/hesusruiz/isbetmf/tmfserver/repository"
)

// Config holds the configuration for the tmfclient package
type Config struct {
	// BaseURL of the remote TMForum server
	BaseURL string `json:"base_url" yaml:"base_url"`

	// Timeout in seconds for HTTP requests
	Timeout int `json:"timeout" yaml:"timeout"`
}

// // Client is a client for the TMForum API.
// type Client struct {
// 	config *Config
// 	client *http.Client
// }

// // NewClient creates a new client.
// func NewClient(config *Config) *Client {
// 	return &Client{
// 		config: config,
// 		client: &http.Client{
// 			Timeout: time.Duration(config.Timeout) * time.Second,
// 		},
// 	}
// }

// // Get sends a GET request to the remote server.
// func (c *Client) Get(path string, headers map[string]string) (*http.Response, error) {
// 	return c.do("GET", path, nil, headers)
// }

// // Post sends a POST request to the remote server.
// func (c *Client) Post(path string, body []byte, headers map[string]string) (*http.Response, error) {
// 	return c.do("POST", path, body, headers)
// }

func ForwardTMFPost(req *Request, remoteServer string, objMap repository.TMFObjectMap) (repository.TMFObjectMap, []error) {

	requestBody, err := json.Marshal(objMap)
	if err != nil {
		return nil, []error{errl.Errorf("failed to marshall object: %w", err)}

	}

	path := fmt.Sprintf("/tmf-api/%s/%s/%s", req.APIfamily, req.APIVersion, req.ResourceName)
	fullURL := remoteServer + path

	agent := fiber.Post(fullURL)
	agent.Body(requestBody)

	agent.Set("Authorization", "Bearer "+req.AccessToken)
	agent.Set("Content-Type", "application/json")
	agent.Set("Content-Length", strconv.Itoa(len(requestBody)))

	statusCode, responseBody, errs := agent.Bytes()
	if len(errs) > 0 {
		return nil, errs
	}

	if statusCode != fiber.StatusCreated {
		return nil, []error{errl.Errorf("unexpected status code: %d", statusCode)}
	}

	obj, err := repository.NewTMFObjectMapFromRequest(req.ResourceName, responseBody)
	if err != nil {
		return nil, []error{errl.Errorf("failed to bind request body: %w", err)}
	}

	validation := obj.Validate(req.ResourceName)
	if len(validation.Errors) > 0 {
		return nil, []error{errl.Errorf("validation errors: %v", validation.Errors)}
	}

	return obj, nil

}

// // Patch sends a PATCH request to the remote server.
// func (c *Client) Patch(path string, body []byte, headers map[string]string) (*http.Response, error) {
// 	return c.do("PATCH", path, body, headers)
// }

// // Delete sends a DELETE request to the remote server.
// func (c *Client) Delete(path string, headers map[string]string) (*http.Response, error) {
// 	return c.do("DELETE", path, nil, headers)
// }

// func (c *Client) do(method, path string, body []byte, headers map[string]string) (*http.Response, error) {
// 	url := fmt.Sprintf("%s%s", c.config.BaseURL, path)
// 	slog.Debug("sending", slog.String("method", method), "url", url)

// 	var req *http.Request
// 	var err error

// 	if body != nil {
// 		req, err = http.NewRequest(method, url, bytes.NewReader(body))
// 	} else {
// 		req, err = http.NewRequest(method, url, nil)
// 	}

// 	if err != nil {
// 		return nil, err
// 	}

// 	// TODO: send the authorization header
// 	for key, value := range headers {
// 		req.Header.Set(key, value)
// 	}

// 	return c.client.Do(req)
// }
