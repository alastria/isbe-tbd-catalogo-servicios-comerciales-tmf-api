package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/hesusruiz/isbetmf/replicate"
)

func main() {
	// Parse command line flags
	var (
		version              = flag.String("version", "v4", "TMForum API version (v4 or v5)")
		baseURL              = flag.String("base-url", "", "Base URL of the TMForum server")
		databaseFile         = flag.String("database", "remotecache.db", "SQLite database file path")
		timeout              = flag.Int("timeout", 30, "Timeout in seconds for HTTP requests")
		objectTypes          = flag.String("object-types", "", "Comma-separated list of object types to replicate")
		paginationEnabled    = flag.Bool("pagination", true, "Enable pagination for object retrieval")
		pageSize             = flag.Int("page-size", 100, "Number of objects per page")
		maxObjects           = flag.Int("max-objects", 10000, "Maximum objects to retrieve per type")
		validateRequired     = flag.Bool("validate-required", true, "Validate required fields")
		validateRelatedParty = flag.Bool("validate-related-party", true, "Validate related party requirements")
		fixValidationErrors  = flag.Bool("fix-errors", true, "Fix validation errors automatically")
		skipInvalid          = flag.Bool("skip-invalid", true, "Skip invalid objects")
		updateExisting       = flag.Bool("update-existing", true, "Update existing objects")
		showStats            = flag.Bool("stats", false, "Show replication statistics")
		clearDatabase        = flag.Bool("clear", false, "Clear the database before replication")
		help                 = flag.Bool("help", false, "Show help information")
	)

	flag.Parse()

	if *help {
		showHelp()
		return
	}

	// Load configuration
	config := replicate.DefaultConfig()

	// Override with command line flags
	if *version != "" {
		config.Version = *version
	}
	if *baseURL != "" {
		config.BaseURL = *baseURL
	}
	if *databaseFile != "" {
		config.DatabaseFile = *databaseFile
	}
	if *timeout > 0 {
		config.Timeout = *timeout
	}
	if *objectTypes != "" {
		config.ObjectTypes = parseObjectTypes(*objectTypes)
	}
	config.PaginationEnabled = *paginationEnabled
	config.PageSize = *pageSize
	config.MaxObjects = *maxObjects
	config.ValidateRequiredFields = *validateRequired
	config.ValidateRelatedParty = *validateRelatedParty
	config.FixValidationErrors = *fixValidationErrors
	config.SkipInvalidObjects = *skipInvalid
	config.UpdateExisting = *updateExisting

	// Load from environment variables
	config.LoadConfigFromEnv()

	// Validate configuration
	if err := config.Validate(); err != nil {
		log.Fatalf("Configuration error: %v", err)
	}

	// Create replicator instance
	replicator, err := replicate.NewReplicator(config)
	if err != nil {
		log.Fatalf("Failed to create replicator: %v", err)
	}
	defer replicator.Close()

	// Set up context with cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigChan
		log.Printf("Received signal %v, shutting down gracefully...", sig)
		cancel()
	}()

	// Show statistics if requested
	if *showStats {
		if err := showStatistics(replicator); err != nil {
			log.Fatalf("Failed to show statistics: %v", err)
		}
		return
	}

	// Clear database if requested
	if *clearDatabase {
		log.Printf("Clearing database...")
		if err := replicator.ClearDatabase(); err != nil {
			log.Fatalf("Failed to clear database: %v", err)
		}
		log.Printf("Database cleared successfully")
	}

	// Run replication process
	log.Printf("Starting replication process...")
	startTime := time.Now()

	if err := replicator.Run(ctx); err != nil {
		log.Fatalf("Replication failed: %v", err)
	}

	duration := time.Since(startTime)
	log.Printf("Replication completed successfully in %v", duration)

	// Show final statistics
	if err := showStatistics(replicator); err != nil {
		log.Printf("Warning: Failed to show final statistics: %v", err)
	}
}

// showStatistics displays replication statistics
func showStatistics(replicator *replicate.Replicator) error {
	stats, err := replicator.GetStatistics()
	if err != nil {
		return fmt.Errorf("failed to get statistics: %w", err)
	}

	fmt.Printf("\n=== Replication Statistics ===\n")
	fmt.Printf("Database File: %s\n", stats.DatabaseFile)
	fmt.Printf("Base URL: %s\n", stats.BaseURL)
	fmt.Printf("Total Objects: %d\n", stats.TotalObjects)
	fmt.Printf("Object Types: %s\n", strings.Join(stats.ObjectTypes, ", "))
	fmt.Printf("\nObjects by Type:\n")
	for objectType, count := range stats.ObjectsByType {
		fmt.Printf("  %s: %d\n", objectType, count)
	}
	fmt.Printf("Generated at: %s\n", stats.GeneratedAt.Format("2006-01-02 15:04:05 UTC"))
	fmt.Printf("=============================\n")

	return nil
}

// parseObjectTypes parses a comma-separated string of object types
func parseObjectTypes(typesStr string) []string {
	if typesStr == "" {
		return nil
	}

	types := strings.Split(typesStr, ",")
	for i, t := range types {
		types[i] = strings.TrimSpace(t)
	}
	return types
}

// showHelp displays help information
func showHelp() {
	fmt.Printf(`TMForum Replicator

Usage: replicate [options]

Options:
  -version string
        TMForum API version (v4 or v5) (default "v4")
  -base-url string
        Base URL of the TMForum server
  -database string
        SQLite database file path (default "remotecache.db")
  -timeout int
        Timeout in seconds for HTTP requests (default 30)
  -object-types string
        Comma-separated list of object types to replicate
  -pagination
        Enable pagination for object retrieval (default true)
  -page-size int
        Number of objects per page (default 100)
  -max-objects int
        Maximum objects to retrieve per type (default 10000)
  -validate-required
        Validate required fields (default true)
  -validate-related-party
        Validate related party requirements (default true)
  -fix-errors
        Fix validation errors automatically (default false)
  -skip-invalid
        Skip invalid objects (default true)
  -update-existing
        Update existing objects (default true)
  -stats
        Show replication statistics
  -clear
        Clear the database before replication
  -help
        Show this help information

Environment Variables:
  REPLICATE_VERSION           TMForum API version (v4 or v5)
  REPLICATE_BASE_URL          Base URL of the TMForum server
  REPLICATE_DATABASE_FILE     SQLite database file path
  REPLICATE_TIMEOUT           Timeout in seconds
  REPLICATE_OBJECT_TYPES      Comma-separated list of object types
  REPLICATE_PAGINATION_ENABLED Enable pagination (true/false)
  REPLICATE_PAGE_SIZE         Number of objects per page
  REPLICATE_MAX_OBJECTS       Maximum objects to retrieve per type
  REPLICATE_VALIDATE_REQUIRED Validate required fields (true/false)
  REPLICATE_VALIDATE_RELATED_PARTY Validate related party requirements (true/false)
  REPLICATE_FIX_VALIDATION_ERRORS Fix validation errors automatically (true/false)
  REPLICATE_SKIP_INVALID      Skip invalid objects (true/false)
  REPLICATE_UPDATE_EXISTING   Update existing objects (true/false)

Examples:
  replicate -base-url "https://tmf.example.com" -object-types "productOffering,productSpecification"
  replicate -database "custom.db" -base-url "https://tmf.example.com"
  replicate -stats
  replicate -clear -base-url "https://tmf.example.com"

`)
}
