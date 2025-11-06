package replicate

import (
	"os"
	"testing"
	"time"

	repo "github.com/hesusruiz/isbetmf/tmfserver/repository"
)

func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig()

	if config.DatabaseFile != "remotecache.db" {
		t.Errorf("Expected DatabaseFile to be 'remotecache.db', got %s", config.DatabaseFile)
	}

	if config.BaseURL != "https://tmf.dome-marketplace-sbx.org" {
		t.Errorf("Expected BaseURL to be 'https://tmf.dome-marketplace-sbx.org', got %s", config.BaseURL)
	}

	if config.Timeout != 30 {
		t.Errorf("Expected Timeout to be 30, got %d", config.Timeout)
	}

	if len(config.ObjectTypes) == 0 {
		t.Error("Expected ObjectTypes to have default values")
	}

	if !config.PaginationEnabled {
		t.Error("Expected PaginationEnabled to be true")
	}

	if config.PageSize != 100 {
		t.Errorf("Expected PageSize to be 100, got %d", config.PageSize)
	}

	if config.MaxObjects != 10000 {
		t.Errorf("Expected MaxObjects to be 10000, got %d", config.MaxObjects)
	}

	if !config.ValidateRequiredFields {
		t.Error("Expected ValidateRequiredFields to be true")
	}

	if !config.ValidateRelatedParty {
		t.Error("Expected ValidateRelatedParty to be true")
	}

	if !config.SkipInvalidObjects {
		t.Error("Expected SkipInvalidObjects to be true")
	}

	if !config.UpdateExisting {
		t.Error("Expected UpdateExisting to be true")
	}
}

func TestConfigValidation(t *testing.T) {
	config := &Config{}

	// Test empty database file
	if err := config.Validate(); err == nil {
		t.Error("Expected error for empty database file")
	}

	// Test empty base URL
	config.DatabaseFile = "test.db"
	if err := config.Validate(); err == nil {
		t.Error("Expected error for empty base URL")
	}

	// Test empty object types
	config.BaseURL = "https://example.com"
	if err := config.Validate(); err == nil {
		t.Error("Expected error for empty object types")
	}

	// Test invalid timeout
	config.ObjectTypes = []string{"productOffering"}
	config.Timeout = 0
	if err := config.Validate(); err == nil {
		t.Error("Expected error for invalid timeout")
	}

	// Test valid configuration
	config.Timeout = 30
	if err := config.Validate(); err != nil {
		t.Errorf("Expected valid configuration, got error: %v", err)
	}
}

func TestDatabaseOperations(t *testing.T) {
	// Create a temporary database file
	dbFile := "test_replicate.db"
	defer os.Remove(dbFile)

	// Create database
	db, err := NewDatabase(dbFile)
	if err != nil {
		t.Fatalf("Failed to create database: %v", err)
	}
	defer db.Close()

	// Test storing an object
	obj := &repo.TMFRecord{
		ID:         "test-id",
		Type:       "productOffering",
		Version:    "1.0",
		APIVersion: "v4",
		LastUpdate: time.Now().Format(time.RFC3339),
		Content:    []byte(`{"id":"test-id","name":"Test Product"}`),
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	err = db.StoreObject(obj)
	if err != nil {
		t.Fatalf("Failed to store object: %v", err)
	}

	// Test retrieving the object
	retrieved, err := db.GetObject("test-id", "productOffering")
	if err != nil {
		t.Fatalf("Failed to get object: %v", err)
	}

	if retrieved == nil {
		t.Fatal("Expected to retrieve object, got nil")
	}

	if retrieved.ID != "test-id" {
		t.Errorf("Expected ID 'test-id', got %s", retrieved.ID)
	}

	// Test listing objects
	objects, err := db.ListObjects("productOffering")
	if err != nil {
		t.Fatalf("Failed to list objects: %v", err)
	}

	if len(objects) != 1 {
		t.Errorf("Expected 1 object, got %d", len(objects))
	}

	// Test object count
	count, err := db.GetObjectCount("productOffering")
	if err != nil {
		t.Fatalf("Failed to get object count: %v", err)
	}

	if count != 1 {
		t.Errorf("Expected count 1, got %d", count)
	}

	// Test total object count
	totalCount, err := db.GetTotalObjectCount()
	if err != nil {
		t.Fatalf("Failed to get total object count: %v", err)
	}

	if totalCount != 1 {
		t.Errorf("Expected total count 1, got %d", totalCount)
	}

	// Test deleting the object
	err = db.DeleteObject("test-id", "productOffering")
	if err != nil {
		t.Fatalf("Failed to delete object: %v", err)
	}

	// Test that object is deleted
	retrieved, err = db.GetObject("test-id", "productOffering")
	if err != nil {
		t.Fatalf("Failed to get object after deletion: %v", err)
	}

	if retrieved != nil {
		t.Error("Expected object to be deleted, but it still exists")
	}
}

func TestReplicatorCreation(t *testing.T) {
	// Create a temporary database file
	dbFile := "test_replicator.db"
	defer os.Remove(dbFile)

	config := &Config{
		DatabaseFile:           dbFile,
		BaseURL:                "https://example.com",
		Timeout:                30,
		ObjectTypes:            []string{"productOffering"},
		PaginationEnabled:      true,
		PageSize:               100,
		MaxObjects:             1000,
		ValidateRequiredFields: true,
		ValidateRelatedParty:   true,
		SkipInvalidObjects:     true,
		UpdateExisting:         true,
	}

	replicator, err := NewReplicator(config)
	if err != nil {
		t.Fatalf("Failed to create replicator: %v", err)
	}
	defer replicator.Close()

	// Test getting statistics
	stats, err := replicator.GetStatistics()
	if err != nil {
		t.Fatalf("Failed to get statistics: %v", err)
	}

	if stats.DatabaseFile != dbFile {
		t.Errorf("Expected database file %s, got %s", dbFile, stats.DatabaseFile)
	}

	if stats.BaseURL != "https://example.com" {
		t.Errorf("Expected base URL 'https://example.com', got %s", stats.BaseURL)
	}

	if stats.TotalObjects != 0 {
		t.Errorf("Expected total objects 0, got %d", stats.TotalObjects)
	}
}

func TestReplicatorClearDatabase(t *testing.T) {
	// Create a temporary database file
	dbFile := "test_clear.db"
	defer os.Remove(dbFile)

	config := &Config{
		DatabaseFile:           dbFile,
		BaseURL:                "https://example.com",
		Timeout:                30,
		ObjectTypes:            []string{"productOffering"},
		PaginationEnabled:      true,
		PageSize:               100,
		MaxObjects:             1000,
		ValidateRequiredFields: true,
		ValidateRelatedParty:   true,
		SkipInvalidObjects:     true,
		UpdateExisting:         true,
	}

	replicator, err := NewReplicator(config)
	if err != nil {
		t.Fatalf("Failed to create replicator: %v", err)
	}
	defer replicator.Close()

	// Add some objects to the database
	db := replicator.database
	obj := &repo.TMFRecord{
		ID:         "test-id",
		Type:       "productOffering",
		Version:    "1.0",
		APIVersion: "v4",
		LastUpdate: time.Now().Format(time.RFC3339),
		Content:    []byte(`{"id":"test-id","name":"Test Product"}`),
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	err = db.StoreObject(obj)
	if err != nil {
		t.Fatalf("Failed to store object: %v", err)
	}

	// Verify object exists
	count, err := db.GetTotalObjectCount()
	if err != nil {
		t.Fatalf("Failed to get total object count: %v", err)
	}

	if count != 1 {
		t.Errorf("Expected 1 object, got %d", count)
	}

	// Clear database
	err = replicator.ClearDatabase()
	if err != nil {
		t.Fatalf("Failed to clear database: %v", err)
	}

	// Verify database is empty
	count, err = db.GetTotalObjectCount()
	if err != nil {
		t.Fatalf("Failed to get total object count after clear: %v", err)
	}

	if count != 0 {
		t.Errorf("Expected 0 objects after clear, got %d", count)
	}
}
