package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/hesusruiz/isbetmf/config"
	"github.com/hesusruiz/isbetmf/internal/errl"
	repo "github.com/hesusruiz/isbetmf/tmfserver/repository"
)

// PagingConfig holds the configuration for the proxy
type PagingConfig struct {

	// Remote server configuration
	BaseURL string `json:"base_url" yaml:"base_url"`
	Timeout int    `json:"timeout" yaml:"timeout"` // in seconds

	// All possible object types to retrieve
	ObjectTypes []string `json:"object_types" yaml:"object_types"`

	// Pagination settings
	PageSize   int `json:"page_size" yaml:"page_size"`     // objects per page
	MaxObjects int `json:"max_objects" yaml:"max_objects"` // maximum objects to retrieve

	// Validation settings
	ValidateRequiredFields bool `json:"validate_required_fields" yaml:"validate_required_fields"`
	ValidateRelatedParty   bool `json:"validate_related_party" yaml:"validate_related_party"`
	FixValidationErrors    bool `json:"fix_validation_errors" yaml:"fix_validation_errors"`

	// Output settings
	OutputDir  string `json:"output_dir" yaml:"output_dir"`
	ReportFile string `json:"report_file" yaml:"report_file"`
}

// ClientWithPaging represents an HTTP client for connecting to TMForum servers supporting pagination
type ClientWithPaging struct {
	httpClient   *http.Client
	baseURL      string
	timeout      time.Duration
	pagingConfig *PagingConfig
}

// NewClientWithPaging creates a new TMForum client
func NewClientWithPaging(config *PagingConfig) *ClientWithPaging {
	return &ClientWithPaging{
		httpClient: &http.Client{
			Timeout: time.Duration(config.Timeout) * time.Second,
		},
		baseURL:      config.BaseURL,
		timeout:      time.Duration(config.Timeout) * time.Second,
		pagingConfig: config,
	}
}

func (c *ClientWithPaging) GetObjectURL(objectType string, objectID string) string {
	pathPrefix, exists := config.GeneratedDefaultResourceToPathPrefix[objectType]
	if !exists {
		return ""
	}
	return fmt.Sprintf("%s%s/%s", c.baseURL, pathPrefix, objectID)
}

type procObj func(objType string, obj repo.TMFObjectMap) (repo.TMFObjectMap, bool, error)

// GetAllObjectsOfType retrieves all objects of a specific type using pagination
func (c *ClientWithPaging) GetAllObjectsOfType(ctx context.Context, objectType string, processObject procObj) ([]repo.TMFObjectMap, error) {
	// Get the path prefix for this object type from the routes map
	pathPrefix, exists := config.GeneratedDefaultResourceToPathPrefix[objectType]
	if !exists {
		return nil, fmt.Errorf("unknown object type: %s", objectType)
	}

	var allObjects []repo.TMFObjectMap
	limit := c.pagingConfig.PageSize
	offset := 0

	for {
		// Build URL with pagination parameters
		url := fmt.Sprintf("%s%s?limit=%d&offset=%d", c.baseURL, pathPrefix, limit, offset)
		fmt.Printf("Retrieving %s objects: offset=%d, limit=%d\n", objectType, offset, limit)

		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %w", err)
		}

		// Set common headers
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Content-Type", "application/json")

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return nil, fmt.Errorf("failed to execute request: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			return nil, fmt.Errorf("server returned status %d: %s", resp.StatusCode, string(body))
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %w", err)
		}

		// Try to parse as an array first
		var objects []repo.TMFObjectMap
		if err := json.Unmarshal(body, &objects); err != nil {
			// If it's not an array, try to parse as a single object
			var singleObject repo.TMFObjectMap
			if err := json.Unmarshal(body, &singleObject); err != nil {
				return nil, fmt.Errorf("failed to parse response as JSON: %w", err)
			}
			objects = []repo.TMFObjectMap{singleObject}
		}

		// Process each object with the user-supplied logic
		// We stop when an error is received or when false is returned
		var cont bool
		if processObject != nil {
			for i := range objects {
				objects[i], cont, err = processObject(objectType, objects[i])
				if err != nil {
					return nil, errl.Errorf("processing object %s", objects[i].ID())
				}
				if !cont {
					// Add objects processed from this page to the total collection
					allObjects = append(allObjects, objects...)
					fmt.Printf("Total %s objects retrieved: %d\n", objectType, len(allObjects))
					return allObjects, nil
				}
			}
		}

		// Add objects from this page to the total collection
		allObjects = append(allObjects, objects...)

		// If we got fewer objects than the limit, we've reached the end
		if len(objects) < limit {
			break
		}

		// Move to next page
		offset += limit

		// Safety check to prevent infinite loops
		if offset >= c.pagingConfig.MaxObjects {
			fmt.Printf("Warning: Reached maximum objects limit (%d) for %s\n", c.pagingConfig.MaxObjects, objectType)
			break
		}
	}

	fmt.Printf("Total %s objects retrieved: %d\n", objectType, len(allObjects))
	return allObjects, nil
}

func (c *ClientWithPaging) GetSingleObject(ctx context.Context, objectType string, id string) (repo.TMFObjectMap, error) {
	// Get the path prefix for this object type from the routes map
	pathPrefix, exists := config.GeneratedDefaultResourceToPathPrefix[objectType]
	if !exists {
		return nil, fmt.Errorf("unknown object type: %s", objectType)
	}

	// Build URL with pagination parameters
	url := fmt.Sprintf("%s%s/%s", c.baseURL, pathPrefix, id)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set common headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("server returned status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// If it's not an array, try to parse as a single object
	var singleObject repo.TMFObjectMap
	if err := json.Unmarshal(body, &singleObject); err != nil {
		return nil, fmt.Errorf("failed to parse response as JSON: %w", err)
	}

	return singleObject, nil
}

// GetPageOfObjects retrieves up to a single page of objects of a specific type
func (c *ClientWithPaging) GetPageOfObjects(ctx context.Context, objectType string, offset int, processObject procObj) ([]repo.TMFObjectMap, error) {
	// Get the path prefix for this object type from the routes map
	pathPrefix, exists := config.GeneratedDefaultResourceToPathPrefix[objectType]
	if !exists {
		return nil, fmt.Errorf("unknown object type: %s", objectType)
	}

	limit := c.pagingConfig.PageSize

	// Build URL with pagination parameters
	url := fmt.Sprintf("%s%s?limit=%d&offset=%d", c.baseURL, pathPrefix, limit, offset)
	fmt.Printf("Retrieving %s objects: offset=%d, limit=%d\n", objectType, offset, limit)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set common headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("server returned status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Try to parse as an array first
	var objects []repo.TMFObjectMap
	if err := json.Unmarshal(body, &objects); err != nil {
		// If it's not an array, try to parse as a single object
		var singleObject repo.TMFObjectMap
		if err := json.Unmarshal(body, &singleObject); err != nil {
			return nil, fmt.Errorf("failed to parse response as JSON: %w", err)
		}
		objects = []repo.TMFObjectMap{singleObject}
	}

	// TODO: this is a hack because the TMF API for listing collections does not work as expected
	// We need to retrieve each object individually, and can not trust on the object that has been retrieved with the list API
	for i := range objects {
		objects[i], err = c.GetSingleObject(ctx, objectType, objects[i].ID())
		if err != nil {
			return nil, err
		}
	}

	// Process each object with the user-supplied logic
	var cont bool
	if processObject != nil {
		for i := range objects {
			objects[i], cont, err = processObject(objectType, objects[i])
			if err != nil {
				return nil, errl.Errorf("processing object %s", objects[i].ID())
			}
			if !cont {
				return objects, nil
			}
		}
	}

	return objects, nil
}

// DefaultPagingConfig returns a default configuration
func DefaultPagingConfig() *PagingConfig {
	return &PagingConfig{
		BaseURL:                "https://tmf.dome-marketplace-sbx.org",
		Timeout:                30,
		ObjectTypes:            DefaultObjectTypes(),
		PageSize:               100,
		MaxObjects:             10000,
		ValidateRequiredFields: true,
		ValidateRelatedParty:   true,
		FixValidationErrors:    false,
		OutputDir:              "./reports",
		ReportFile:             "tmf_validation_report.md",
	}
}

// DefaultObjectTypes returns the default list of object types to retrieve
func DefaultObjectTypes() []string {

	objectTypes := []string{}
	for objectType := range config.GeneratedDefaultResourceToPathPrefix {
		objectTypes = append(objectTypes, objectType)
	}

	return objectTypes
}
