package replicate

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hesusruiz/isbetmf/reporting"
)

// Replicator orchestrates the replication process from remote TMForum servers to local database
type Replicator struct {
	config    *Config
	database  *Database
	client    *reporting.ClientPaging
	validator *Validator
	converter *Converter
}

// NewReplicator creates a new replicator instance
func NewReplicator(config *Config) (*Replicator, error) {
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	// Create database connection
	database, err := NewDatabase(config.DatabaseFile)
	if err != nil {
		return nil, fmt.Errorf("failed to create database: %w", err)
	}

	// Create reporting configuration for client and validator
	reportingConfig := &reporting.Config{
		Version:                config.Version,
		BaseURL:                config.BaseURL,
		Timeout:                config.Timeout,
		ObjectTypes:            config.ObjectTypes,
		PaginationEnabled:      config.PaginationEnabled,
		PageSize:               config.PageSize,
		MaxObjects:             config.MaxObjects,
		ValidateRequiredFields: config.ValidateRequiredFields,
		ValidateRelatedParty:   config.ValidateRelatedParty,
		FixValidationErrors:    config.FixValidationErrors,
	}

	// Create client, validator, and converter
	client := reporting.NewClientPaging(reportingConfig)
	validator := NewValidator(config)
	converter := NewConverter()

	return &Replicator{
		config:    config,
		database:  database,
		client:    client,
		validator: validator,
		converter: converter,
	}, nil
}

// Close closes the replicator and its database connection
func (r *Replicator) Close() error {
	if r.database != nil {
		return r.database.Close()
	}
	return nil
}

// Run performs the complete replication process
func (r *Replicator) Run(ctx context.Context) error {
	log.Printf("Starting TMForum object replication process")
	log.Printf("Base URL: %s", r.config.BaseURL)
	log.Printf("Database: %s", r.config.DatabaseFile)
	log.Printf("Object types: %v", r.config.ObjectTypes)

	// Test connection to remote server
	log.Printf("Testing connection to remote server...")
	if err := r.client.TestConnection(ctx); err != nil {
		return fmt.Errorf("connection test failed: %w", err)
	}
	log.Printf("Connection successful")

	// Process each object type
	var totalObjects int
	var totalValid int
	var totalInvalid int
	var totalErrors int
	var totalWarnings int

	for _, objectType := range r.config.ObjectTypes {
		log.Printf("Processing object type: %s", objectType)

		objects, valid, invalid, errors, warnings, err := r.replicateObjectType(ctx, objectType)
		if err != nil {
			log.Printf("Warning: Failed to replicate %s objects: %v", objectType, err)
			continue
		}

		totalObjects += objects
		totalValid += valid
		totalInvalid += invalid
		totalErrors += errors
		totalWarnings += warnings

		log.Printf("Replication complete for %s: %d objects, %d valid, %d invalid, %d errors, %d warnings",
			objectType, objects, valid, invalid, errors, warnings)
	}

	// Log final summary
	log.Printf("Replication process complete")
	log.Printf("Total objects processed: %d", totalObjects)
	log.Printf("Valid objects: %d", totalValid)
	log.Printf("Invalid objects: %d", totalInvalid)
	log.Printf("Total errors: %d", totalErrors)
	log.Printf("Total warnings: %d", totalWarnings)

	return nil
}

// replicateObjectType replicates all objects of a specific type
func (r *Replicator) replicateObjectType(ctx context.Context, objectType string) (int, int, int, int, int, error) {
	var totalObjects int
	var totalValid int
	var totalInvalid int
	var totalErrors int
	var totalWarnings int

	// Process objects page by page to avoid loading all into memory
	offset := 0

	for {
		// Get a single page of objects
		objects, err := r.client.GetObjectsPage(ctx, objectType, &reporting.Config{
			BaseURL:           r.config.BaseURL,
			Timeout:           r.config.Timeout,
			PaginationEnabled: r.config.PaginationEnabled,
			PageSize:          r.config.PageSize,
			MaxObjects:        r.config.MaxObjects,
		}, offset)
		if err != nil {
			return totalObjects, totalValid, totalInvalid, totalErrors, totalWarnings, err
		}

		// If no objects returned, we've reached the end
		if len(objects) == 0 {
			break
		}

		log.Printf("Retrieved %d %s objects (offset %d)", len(objects), objectType, offset)
		totalObjects += len(objects)

		// Process each object
		for _, obj := range objects {
			valid, invalid, errors, warnings, err := r.processObject(ctx, obj, objectType)
			if err != nil {
				log.Printf("Warning: Failed to process object %s: %v", obj.ID, err)
				continue
			}

			totalValid += valid
			totalInvalid += invalid
			totalErrors += errors
			totalWarnings += warnings
		}

		// If we got fewer objects than the page size, we've reached the end
		if len(objects) < r.config.PageSize {
			break
		}

		// Move to next page
		offset += r.config.PageSize

		// Safety check to prevent infinite loops
		if offset >= r.config.MaxObjects {
			log.Printf("Warning: Reached maximum objects limit (%d) for %s", r.config.MaxObjects, objectType)
			break
		}
	}

	return totalObjects, totalValid, totalInvalid, totalErrors, totalWarnings, nil
}

// processObject processes a single object: validates it and stores it in the database
func (r *Replicator) processObject(ctx context.Context, obj reporting.TMFObject, objectType string) (int, int, int, int, error) {
	// Convert reporting.TMFObject to replicate.TMFObject
	replicateObj, err := r.converter.ReportingTMFObjectToReplicateTMFObject(obj)
	if err != nil {
		return 0, 1, 1, 0, fmt.Errorf("failed to convert object: %w", err)
	}

	// Validate and optionally fix the object
	result := r.validator.ValidateAndFixObject(replicateObj, objectType)

	// Count validation results
	valid := 0
	invalid := 0
	errors := len(result.Errors)
	warnings := len(result.Warnings)

	if result.Valid {
		valid = 1
	} else {
		invalid = 1
	}

	// Log fixing results if any
	if len(result.ErrorsFixed) > 0 || len(result.WarningsFixed) > 0 {
		log.Printf("Fixed object %s (%s): %d errors fixed, %d warnings fixed",
			replicateObj.GetID(), objectType, len(result.ErrorsFixed), len(result.WarningsFixed))
	}

	// If object is invalid and we're configured to skip invalid objects, return early
	if !result.Valid && r.config.SkipInvalidObjects {
		log.Printf("Skipping invalid object %s: %d errors, %d warnings", replicateObj.GetID(), errors, warnings)
		return valid, invalid, errors, warnings, nil
	}

	// Convert replicate.TMFObject to repository.TMFObject
	repoObj, err := r.converter.ReplicateTMFObjectToRepositoryTMFObject(replicateObj, objectType)
	if err != nil {
		return valid, invalid, errors, warnings, fmt.Errorf("failed to convert to repository object: %w", err)
	}

	// Store the object in the database
	if err := r.database.StoreObject(repoObj); err != nil {
		return valid, invalid, errors, warnings, fmt.Errorf("failed to store object: %w", err)
	}

	log.Printf("Stored object %s (%s): valid=%t, errors=%d, warnings=%d, errors_fixed=%d, warnings_fixed=%d",
		replicateObj.GetID(), objectType, result.Valid, errors, warnings, len(result.ErrorsFixed), len(result.WarningsFixed))
	return valid, invalid, errors, warnings, nil
}

// ClearDatabase removes all objects from the database
func (r *Replicator) ClearDatabase() error {
	return r.database.ClearDatabase()
}

// GetStatistics returns replication statistics
func (r *Replicator) GetStatistics() (*Statistics, error) {
	stats := &Statistics{
		DatabaseFile: r.config.DatabaseFile,
		BaseURL:      r.config.BaseURL,
		ObjectTypes:  r.config.ObjectTypes,
	}

	// Get total object count
	totalCount, err := r.database.GetTotalObjectCount()
	if err != nil {
		return nil, fmt.Errorf("failed to get total object count: %w", err)
	}
	stats.TotalObjects = totalCount

	// Get counts by object type
	stats.ObjectsByType = make(map[string]int)
	for _, objectType := range r.config.ObjectTypes {
		count, err := r.database.GetObjectCount(objectType)
		if err != nil {
			log.Printf("Warning: Failed to get count for %s: %v", objectType, err)
			continue
		}
		stats.ObjectsByType[objectType] = count
	}

	return stats, nil
}

// Statistics holds replication statistics
type Statistics struct {
	DatabaseFile  string         `json:"database_file"`
	BaseURL       string         `json:"base_url"`
	ObjectTypes   []string       `json:"object_types"`
	TotalObjects  int            `json:"total_objects"`
	ObjectsByType map[string]int `json:"objects_by_type"`
	GeneratedAt   time.Time      `json:"generated_at"`
}
