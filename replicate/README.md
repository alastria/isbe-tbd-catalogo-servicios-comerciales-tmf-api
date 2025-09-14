# TMForum Replicate Package

The TMForum Replicate package provides functionality to retrieve objects from remote TMForum servers, validate them, and store them in a local SQLite database. It builds upon the existing reporting package for validation and the tmfserver repository for database operations.

## Features

- **Remote Object Retrieval**: Retrieve objects from any TMForum-compliant server
- **Object Validation**: Validate objects using the same validation rules as the reporting package
- **Automatic Fixing**: Optionally fix validation errors automatically during validation
- **SQLite Storage**: Store validated objects in a local SQLite database
- **Pagination Support**: Handle large datasets with automatic pagination
- **Configurable**: Support for configuration files, environment variables, and command-line options
- **Statistics**: Track replication progress and database statistics
- **Database Management**: Clear database, update existing objects, skip invalid objects

## Architecture

The package is organized into several components:

- **Config**: Configuration management with environment variable support
- **Database**: SQLite database operations using the existing repository structure
- **Replicator**: Main orchestrator that coordinates retrieval, validation, and storage
- **CLI**: Command-line interface for easy operation

## Installation

```bash
go get github.com/hesusruiz/isbetmf/replicate
```

## Quick Start

### Basic Usage

```go
package main

import (
    "context"
    "log"
    
    "github.com/hesusruiz/isbetmf/replicate"
)

func main() {
    // Create configuration
    config := replicate.DefaultConfig()
    config.BaseURL = "https://tmf.example.com"
    config.DatabaseFile = "remotecache.db"
    
    // Create replicator instance
    replicator, err := replicate.NewReplicator(config)
    if err != nil {
        log.Fatal(err)
    }
    defer replicator.Close()
    
    // Run replication
    ctx := context.Background()
    if err := replicator.Run(ctx); err != nil {
        log.Fatal(err)
    }
}
```

### Command Line Interface

A command-line interface is provided in `cmd/replicate/`:

```bash
# Build the binary
go build -o replicate cmd/replicate/main.go

# Run with basic options
./replicate -base-url "https://tmf.example.com"

# Run with custom database file
./replicate -base-url "https://tmf.example.com" -database "custom.db"

# Run with specific object types
./replicate -base-url "https://tmf.example.com" -object-types "productOffering,productSpecification"

# Show statistics
./replicate -stats

# Clear database before replication
./replicate -clear -base-url "https://tmf.example.com"

# Show help
./replicate -help
```

## Configuration

### Default Configuration

The package provides sensible defaults:

```go
config := replicate.DefaultConfig()
// Returns:
// - DatabaseFile: "remotecache.db"
// - BaseURL: "https://tmf.dome-marketplace-sbx.org"
// - Timeout: 30 seconds
// - ObjectTypes: Default object types from reporting package
// - PaginationEnabled: true
// - PageSize: 100
// - MaxObjects: 10000
// - ValidateRequiredFields: true
// - ValidateRelatedParty: true
// - SkipInvalidObjects: true
// - UpdateExisting: true
```

### Environment Variables

Set environment variables to override configuration:

```bash
export REPLICATE_BASE_URL="https://tmf.example.com"
export REPLICATE_DATABASE_FILE="custom.db"
export REPLICATE_TIMEOUT="60"
export REPLICATE_OBJECT_TYPES="productOffering,productSpecification"
export REPLICATE_PAGINATION_ENABLED="true"
export REPLICATE_PAGE_SIZE="100"
export REPLICATE_MAX_OBJECTS="10000"
export REPLICATE_VALIDATE_REQUIRED="true"
export REPLICATE_VALIDATE_RELATED_PARTY="true"
export REPLICATE_FIX_VALIDATION_ERRORS="false"
export REPLICATE_SKIP_INVALID="true"
export REPLICATE_UPDATE_EXISTING="true"
```

### Programmatic Configuration

```go
config := &replicate.Config{
    DatabaseFile:           "remotecache.db",
    BaseURL:                "https://tmf.example.com",
    Timeout:                30,
    ObjectTypes:            []string{"productOffering", "productSpecification"},
    PaginationEnabled:      true,
    PageSize:               100,
    MaxObjects:             10000,
    ValidateRequiredFields: true,
    ValidateRelatedParty:   true,
    FixValidationErrors:    false,
    SkipInvalidObjects:     true,
    UpdateExisting:         true,
}

// Load from environment variables
config.LoadConfigFromEnv()

// Validate configuration
if err := config.Validate(); err != nil {
    log.Fatal(err)
}
```

## Database Schema

The package uses the same database schema as the existing tmfserver repository:

```sql
CREATE TABLE IF NOT EXISTS tmf_object (
    "id" TEXT NOT NULL,
    "type" TEXT NOT NULL,
    "version" TEXT,
    "api_version" TEXT,
    "last_update" TEXT,
    "content" BLOB NOT NULL,
    "created_at" DATETIME NOT NULL,
    "updated_at" DATETIME NOT NULL,
    PRIMARY KEY ("id", "type", "version")
);
```

## Object Types

The package supports the same object types as the reporting package:

- `productOffering`
- `productSpecification`
- `productOfferingPrice`
- `category`
- `individual`
- `organization`
- `productCatalog`
- `customer`
- `product`
- `service`

## Validation

The package uses the same validation rules as the reporting package:

- **Required Fields**: All objects must have `id`, `href`, `lastUpdate`, and `version`
- **Related Party Requirements**: Objects must include related party information with required roles
- **Version Support**: Supports both TMForum API v4 and v5 validation rules
- **Automatic Fixing**: Optionally fix validation errors automatically (disabled by default)
- **Fixed Errors Tracking**: Track which errors and warnings were automatically fixed

## Statistics

The package provides comprehensive statistics:

```go
stats, err := replicator.GetStatistics()
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Database File: %s\n", stats.DatabaseFile)
fmt.Printf("Base URL: %s\n", stats.BaseURL)
fmt.Printf("Total Objects: %d\n", stats.TotalObjects)
fmt.Printf("Objects by Type: %v\n", stats.ObjectsByType)
```

## CLI Options

The command-line interface provides comprehensive options:

```bash
# Basic options
-base-url string          Base URL of the TMForum server
-database string          SQLite database file path (default "remotecache.db")
-timeout int              Timeout in seconds for HTTP requests (default 30)
-object-types string      Comma-separated list of object types to replicate

# Pagination options
-pagination               Enable pagination for object retrieval (default true)
-page-size int            Number of objects per page (default 100)
-max-objects int          Maximum objects to retrieve per type (default 10000)

# Validation options
-validate-required        Validate required fields (default true)
-validate-related-party   Validate related party requirements (default true)
-fix-errors               Fix validation errors automatically (default false)

# Replication options
-skip-invalid             Skip invalid objects (default true)
-update-existing          Update existing objects (default true)

# Other options
-stats                    Show replication statistics
-clear                    Clear the database before replication
-help                     Show help information
```

## Error Handling

The package provides comprehensive error handling:

- **Connection Errors**: Network connectivity issues
- **Validation Errors**: Object validation failures
- **Database Errors**: SQLite operation failures
- **Configuration Errors**: Invalid configuration settings

All errors include context and can be wrapped for additional information.

## Testing

Run the test suite:

```bash
go test ./replicate/...
```

Run with coverage:

```bash
go test -cover ./replicate/...
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass
6. Submit a pull request

## License

This package is licensed under the same license as the parent project.

## Support

For issues and questions:

1. Check the documentation
2. Search existing issues
3. Create a new issue with detailed information
4. Include configuration and error logs
