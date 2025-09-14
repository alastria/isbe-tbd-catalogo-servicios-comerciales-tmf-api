package replicate

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Config holds the configuration for the replicate package
type Config struct {
	// Version of the TMForum API
	Version string `json:"version" yaml:"version"`

	// Database configuration
	DatabaseFile string `json:"database_file" yaml:"database_file"`

	// Remote server configuration (inherited from reporting package)
	BaseURL string `json:"base_url" yaml:"base_url"`
	Timeout int    `json:"timeout" yaml:"timeout"` // in seconds

	// Object types to replicate
	ObjectTypes []string `json:"object_types" yaml:"object_types"`

	// Pagination settings
	PaginationEnabled bool `json:"pagination_enabled" yaml:"pagination_enabled"`
	PageSize          int  `json:"page_size" yaml:"page_size"`     // objects per page
	MaxObjects        int  `json:"max_objects" yaml:"max_objects"` // maximum objects to retrieve

	// Validation settings
	ValidateRequiredFields bool `json:"validate_required_fields" yaml:"validate_required_fields"`
	ValidateRelatedParty   bool `json:"validate_related_party" yaml:"validate_related_party"`
	FixValidationErrors    bool `json:"fix_validation_errors" yaml:"fix_validation_errors"`

	// Replication settings
	SkipInvalidObjects bool `json:"skip_invalid_objects" yaml:"skip_invalid_objects"`
	UpdateExisting     bool `json:"update_existing" yaml:"update_existing"`
}

// DefaultConfig returns a default configuration
func DefaultConfig() *Config {
	return &Config{
		Version:                "v4",
		DatabaseFile:           "remotecache.db",
		BaseURL:                "https://tmf.dome-marketplace-sbx.org",
		Timeout:                30,
		ObjectTypes:            DefaultObjectTypes(),
		PaginationEnabled:      true,
		PageSize:               100,
		MaxObjects:             10000,
		ValidateRequiredFields: true,
		ValidateRelatedParty:   true,
		FixValidationErrors:    false,
		SkipInvalidObjects:     true,
		UpdateExisting:         true,
	}
}

// DefaultObjectTypes returns the default list of object types to replicate
func DefaultObjectTypes() []string {
	// Use the same object types as the reporting package
	return []string{
		"productOffering",
		"productSpecification",
		"productOfferingPrice",
		"category",
		"individual",
		"organization",
		"productCatalog",
		"customer",
		"product",
		"service",
	}
}

// LoadConfigFromEnv loads configuration from environment variables
func (c *Config) LoadConfigFromEnv() {
	if version := os.Getenv("REPLICATE_VERSION"); version != "" {
		c.Version = version
	}

	if databaseFile := os.Getenv("REPLICATE_DATABASE_FILE"); databaseFile != "" {
		c.DatabaseFile = databaseFile
	}

	if baseURL := os.Getenv("REPLICATE_BASE_URL"); baseURL != "" {
		c.BaseURL = baseURL
	}

	if timeout := os.Getenv("REPLICATE_TIMEOUT"); timeout != "" {
		if t, err := strconv.Atoi(timeout); err == nil && t > 0 {
			c.Timeout = t
		}
	}

	if objectTypes := os.Getenv("REPLICATE_OBJECT_TYPES"); objectTypes != "" {
		c.ObjectTypes = strings.Split(objectTypes, ",")
		for i, t := range c.ObjectTypes {
			c.ObjectTypes[i] = strings.TrimSpace(t)
		}
	}

	// Load pagination settings from environment
	if paginationEnabled := os.Getenv("REPLICATE_PAGINATION_ENABLED"); paginationEnabled != "" {
		if enabled, err := strconv.ParseBool(paginationEnabled); err == nil {
			c.PaginationEnabled = enabled
		}
	}

	if pageSize := os.Getenv("REPLICATE_PAGE_SIZE"); pageSize != "" {
		if size, err := strconv.Atoi(pageSize); err == nil && size > 0 {
			c.PageSize = size
		}
	}

	if maxObjects := os.Getenv("REPLICATE_MAX_OBJECTS"); maxObjects != "" {
		if max, err := strconv.Atoi(maxObjects); err == nil && max > 0 {
			c.MaxObjects = max
		}
	}

	// Load validation settings from environment
	if validateRequired := os.Getenv("REPLICATE_VALIDATE_REQUIRED"); validateRequired != "" {
		if enabled, err := strconv.ParseBool(validateRequired); err == nil {
			c.ValidateRequiredFields = enabled
		}
	}

	if validateRelatedParty := os.Getenv("REPLICATE_VALIDATE_RELATED_PARTY"); validateRelatedParty != "" {
		if enabled, err := strconv.ParseBool(validateRelatedParty); err == nil {
			c.ValidateRelatedParty = enabled
		}
	}

	if fixErrors := os.Getenv("REPLICATE_FIX_VALIDATION_ERRORS"); fixErrors != "" {
		if enabled, err := strconv.ParseBool(fixErrors); err == nil {
			c.FixValidationErrors = enabled
		}
	}

	// Load replication settings from environment
	if skipInvalid := os.Getenv("REPLICATE_SKIP_INVALID"); skipInvalid != "" {
		if enabled, err := strconv.ParseBool(skipInvalid); err == nil {
			c.SkipInvalidObjects = enabled
		}
	}

	if updateExisting := os.Getenv("REPLICATE_UPDATE_EXISTING"); updateExisting != "" {
		if enabled, err := strconv.ParseBool(updateExisting); err == nil {
			c.UpdateExisting = enabled
		}
	}
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if c.DatabaseFile == "" {
		return fmt.Errorf("database file is required")
	}

	if c.BaseURL == "" {
		return fmt.Errorf("base URL is required")
	}

	if len(c.ObjectTypes) == 0 {
		return fmt.Errorf("at least one object type must be specified")
	}

	if c.Timeout <= 0 {
		return fmt.Errorf("timeout must be positive")
	}

	c.BaseURL = strings.TrimRight(c.BaseURL, "/")

	return nil
}
