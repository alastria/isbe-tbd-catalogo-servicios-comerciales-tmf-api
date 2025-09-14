# TMForum Reporting Package

The TMForum Reporting package provides functionality to connect to remote TMForum servers, retrieve objects of various types, validate them against requirements, and generate comprehensive reports. It supports both TMForum API v4 and v5 specifications with different validation rules for each version.

## Features

- **Multi-Version API Support**: Full support for TMForum API v4 and v5 with version-specific validation rules
- **Remote Server Connection**: Connect to any TMForum-compliant server with configurable base URL
- **Comprehensive Object Retrieval**: Retrieve and validate 49+ different object types from various TMForum domains
- **Smart Pagination**: Automatic pagination support to retrieve all objects efficiently (100 per page by default)
- **Advanced Validation**: Validate objects for required fields and related party requirements with version-specific rules
- **Comprehensive Reporting**: Generate detailed Markdown reports with statistics and error details
- **Flexible Configuration**: Support for configuration files (JSON/YAML), environment variables, and command-line options
- **Progress Tracking**: Real-time progress reporting with detailed stages for long-running operations
- **CLI Interface**: Complete command-line interface with comprehensive options and help system
- **Connection Testing**: Built-in connection testing and server health checks

## Architecture

The package is organized into several well-defined components:

- **Config**: Configuration management with environment variable support and validation
- **Client**: HTTP client for connecting to TMForum servers with automatic path prefix resolution
- **Validator**: Object validation against requirements with version-specific rules (v4/v5)
- **Reporter**: Report generation in Markdown format with comprehensive statistics
- **Proxy**: Main orchestrator that coordinates all components and manages the validation workflow
- **Routes**: Automatic path prefix mapping for 49+ different resource types across all TMForum domains

## Installation

```bash
go get github.com/hesusruiz/isbetmf/reporting
```

## Quick Start

### Basic Usage

```go
package main

import (
    "context"
    "log"
    
    "github.com/hesusruiz/isbetmf/reporting"
)

func main() {
    // Create configuration
    config := reporting.DefaultConfig()
    config.BaseURL = "https://tmf.example.com"
    
    // Create proxy instance
    proxyInstance, err := reporting.NewProxy(config)
    if err != nil {
        log.Fatal(err)
    }
    
    // Run validation
    ctx := context.Background()
    if err := proxyInstance.Run(ctx); err != nil {
        log.Fatal(err)
    }
}
```

### Command Line Interface

A comprehensive command-line interface is provided in `cmd/reporting/`:

```bash
# Build the binary
go build -o tmfproxy cmd/reporting/main.go

# Run with basic options
./tmfproxy -base-url "https://tmf.example.com"

# Run with progress tracking
./tmfproxy -base-url "https://tmf.example.com" -progress

# Run with custom object types
./tmfproxy -base-url "https://tmf.example.com" -object-types "productOffering,productSpecification"

# Run with configuration file
./tmfproxy -config config.yaml

# Run with environment variables
TMF_BASE_URL="https://tmf.example.com" TMF_OBJECT_TYPES="productOffering,productSpecification" ./tmfproxy

# Show help
./tmfproxy -help
```

### CLI Options

The command-line interface provides comprehensive options:

```bash
# Basic options
-base-url string          Base URL of the TMForum server
-timeout int              Timeout in seconds for HTTP requests (default 30)
-object-types string      Comma-separated list of object types to validate
-output-dir string        Output directory for reports (default "./reports")
-report-file string       Name of the report file (default "tmf_validation_report.md")

# Pagination options
-pagination               Enable pagination for object retrieval (default true)
-page-size int            Number of objects per page (default 100)
-max-objects int          Maximum objects to retrieve per type (default 10000)

# Validation options
-validate-required        Validate required fields (default true)
-validate-related-party   Validate related party requirements (default true)

# Other options
-progress                 Show progress updates
-config string            Configuration file (JSON or YAML)
-help                     Show help information
```

## Configuration

### Default Configuration

The package provides sensible defaults:

```go
config := reporting.DefaultConfig()
// Returns:
// - Version: "v4"
// - BaseURL: "https://tmf.dome-marketplace-sbx.org"
// - Timeout: 30 seconds
// - ObjectTypes: All 49+ supported object types
// - PaginationEnabled: true
// - PageSize: 100
// - MaxObjects: 10000
// - ValidateRequiredFields: true
// - ValidateRelatedParty: true
// - OutputDir: "./reports"
// - ReportFile: "tmf_validation_report.md"
```

### Configuration File

Create a configuration file (YAML or JSON):

```yaml
# config.yaml
version: "v4"  # API version: "v4" or "v5"
base_url: "https://tmf.example.com"
timeout: 30
object_types:
  - "productOffering"
  - "productSpecification"
  - "productOfferingPrice"
  - "category"
  - "individual"
  - "organization"
pagination_enabled: true
page_size: 100
max_objects: 10000
validate_required_fields: true
validate_related_party: true
output_dir: "./reports"
report_file: "tmf_validation_report.md"
```

### Environment Variables

Set environment variables to override configuration:

```bash
export TMF_BASE_URL="https://tmf.example.com"
export TMF_TIMEOUT="60"
export TMF_OBJECT_TYPES="productOffering,productSpecification,productOfferingPrice"
export TMF_OUTPUT_DIR="./custom_reports"
export TMF_REPORT_FILE="custom_report.md"
export TMF_PAGINATION_ENABLED="true"
export TMF_PAGE_SIZE="100"
export TMF_MAX_OBJECTS="10000"
```

### Programmatic Configuration

```go
config := &reporting.Config{
    Version:                "v4", // or "v5"
    BaseURL:                "https://tmf.example.com",
    Timeout:                30,
    ObjectTypes:            []string{"productOffering", "productSpecification"},
    PaginationEnabled:      true,
    PageSize:               100,
    MaxObjects:             10000,
    ValidateRequiredFields: true,
    ValidateRelatedParty:   true,
    OutputDir:              "./reports",
    ReportFile:             "tmf_validation_report.md",
}

// Load from environment variables
config.LoadConfigFromEnv()

// Validate configuration
if err := config.Validate(); err != nil {
    log.Fatal(err)
}
```

## Object Types

The package supports 49+ different object types across all TMForum domains. Here are the main categories:

### Product Catalog Management
- `productOffering`, `productSpecification`, `productOfferingPrice`, `category`, `catalog`

### Party Management
- `individual`, `organization`, `partyRole`

### Product & Service Inventory
- `product`, `service`, `resource`

### Order Management
- `productOrder`, `serviceOrder`, `resourceOrder`, `cancelProductOrder`, `cancelServiceOrder`, `cancelResourceOrder`

### Customer & Account Management
- `customer`, `billingAccount`, `financialAccount`, `partyAccount`, `settlementAccount`

### Billing & Usage
- `customerBill`, `customerBillOnDemand`, `appliedCustomerBillingRate`, `usage`, `usageSpecification`

### Resource Management
- `resourceCandidate`, `resourceCatalog`, `resourceCategory`, `resourceSpecification`, `resourceFunction`

### Service Management
- `serviceCandidate`, `serviceCatalog`, `serviceCategory`, `serviceSpecification`

### Agreement & Quote Management
- `agreement`, `agreementSpecification`, `quote`

### Resource Function Activation
- `heal`, `migrate`, `monitor`, `scale`

### Account Management
- `billFormat`, `billPresentationMedia`, `billingCycleSpecification`

Each object type automatically maps to the correct TMForum API endpoint using predefined path prefixes. The path prefixes are automatically generated and maintained, ensuring compatibility with the latest TMForum specifications.

You can customize this list in your configuration, but all object types must exist in the routes map to be processed.

## URL Building

The proxy automatically constructs the correct URLs for each object type using predefined path prefixes. For example:

- `productOffering` → `/tmf-api/productCatalogManagement/v4/productOffering`
- `individual` → `/tmf-api/party/v4/individual`
- `category` → `/tmf-api/productCatalogManagement/v4/category`
- `customerBill` → `/tmf-api/customerBillManagement/v4/customerBill`
- `resourceFunction` → `/tmf-api/resourceFunctionActivation/v4/resourceFunction`

The path prefixes are automatically generated and maintained, ensuring compatibility with the latest TMForum specifications. If an unknown object type is specified, the proxy will return an error.

## Pagination

The proxy automatically handles pagination to retrieve all objects from TMForum servers:

- **Default Page Size**: 100 objects per page (configurable)
- **Automatic Detection**: Stops when fewer objects than the page size are returned
- **Efficient Retrieval**: Uses `limit` and `offset` parameters in API calls
- **Safety Limits**: Configurable maximum objects per type (default: 10,000)
- **Configurable**: Can be disabled or customized via configuration

Example pagination flow:
1. Request: `?limit=100&offset=0` → Get first 100 objects
2. Request: `?limit=100&offset=100` → Get next 100 objects
3. Request: `?limit=100&offset=200` → Get next 100 objects
4. Continue until fewer than 100 objects are returned

## Validation Rules

### API Version Support

The package supports both TMForum API v4 and v5 with different validation rules:

- **v4 API**: Uses `RelatedPartyV4` structure with direct field validation
- **v5 API**: Uses `RelatedPartyV5` structure with nested party references

### Required Fields

All objects must have the following fields:
- `id`: Unique identifier
- `href`: Resource reference
- `lastUpdate`: Last update timestamp
- `version`: Object version

### Related Party Requirements

#### API v4 Objects
Objects must include related party information with the following roles:
- `seller`: The selling party (required for most objects)
- `selleroperator`: The operator responsible for selling (required for most objects)
- `buyer`: The buying party (required for some objects)
- `buyeroperator`: The operator responsible for buying (required for some objects)

#### API v5 Objects
Objects must include related party information with the following roles:
- `Seller`: The selling party
- `SellerOperator`: The operator responsible for selling

### Object-Specific Rules

Some object types have special validation rules:
- **Excluded from Related Party Validation**: `category`, `individual`, `organization`
- **Excluded from Buyer Info**: `catalog`, `productOffering`, `productSpecification`, `productOfferingPrice`, `resourceSpecification`, `serviceSpecification`

## Report Output

Reports are generated in Markdown format and include:

- **Summary Statistics**: Total objects, valid/invalid counts, errors, warnings
- **Statistics by Object Type**: Breakdown by each object type
- **Error Summary**: Count of each error type
- **Warning Summary**: Count of each warning type
- **Detailed Results**: Individual validation results for each object

### Report Structure

```
# TMForum Object Validation Report

## Summary Statistics
| Metric | Value |
|--------|-------|
| Total Objects | 150 |
| Valid Objects | 142 |
| Invalid Objects | 8 |
| Total Errors | 12 |
| Total Warnings | 25 |

## Statistics by Object Type
| Object Type | Count | Valid | Invalid | Errors | Warnings |
|-------------|-------|-------|---------|--------|----------|
| productOffering | 50 | 48 | 2 | 3 | 8 |

## Error Summary
| Error Code | Count |
|-------------|-------|
| MISSING_REQUIRED_FIELD | 8 |
| MISSING_RELATED_PARTY | 4 |

## Detailed Validation Results
### productOffering Objects
#### Object: PO-001
- **Type:** ProductOffering
- **Valid:** true
- **Timestamp:** 2024-01-15 10:30:00 UTC
```

## Error Codes

### Validation Errors

- `MISSING_REQUIRED_FIELD`: Required field is missing
- `MISSING_RELATED_PARTY`: Related party information is missing
- `MISSING_REQUIRED_ROLE`: Required related party role is missing
- `MISSING_PARTY_ID`: Related party ID is missing
- `MISSING_PARTY_HREF`: Related party href is missing
- `MISSING_PARTY_NAME`: Related party name is missing (v4 only)
- `MISSING_PARTY_REFERRED_TYPE`: Related party referred type is missing (v4 only)

### Validation Warnings

- `MISSING_REQUIRED_ROLE`: Required related party role is missing
- `EMPTY_ROLE`: Related party role is empty
- `MISSING_PARTY_ID`: Related party ID is missing
- `MISSING_PARTY_HREF`: Related party href is missing

## Advanced Usage

### Progress Tracking

```go
progressChan := make(chan reporting.ProgressUpdate)

go func() {
    if err := proxyInstance.RunWithProgress(ctx, progressChan); err != nil {
        log.Printf("Error: %v", err)
    }
}()

for update := range progressChan {
    fmt.Printf("[%s] %s - %d%%\n", 
        update.Stage, update.Message, update.Progress)
}
```

### Custom Validation

```go
// Create custom validator
validator := reporting.NewValidator(config)

// Validate individual objects
result := validator.ValidateObject(obj, objectType)

// Validate multiple objects
results := validator.ValidateObjects(objects, objectType)
```

### Custom Reporting

```go
// Create custom reporter
reporter := reporting.NewReporter(config)

// Generate report
report, err := reporter.GenerateReport(results)
if err != nil {
    log.Fatal(err)
}

// Access statistics
fmt.Printf("Total objects: %d\n", report.Statistics.TotalObjects)
fmt.Printf("Valid objects: %d\n", report.Statistics.ValidObjects)
```

## Error Handling

The package provides comprehensive error handling:

- **Connection Errors**: Network connectivity issues
- **Validation Errors**: Object validation failures
- **Configuration Errors**: Invalid configuration settings
- **Report Generation Errors**: File system or formatting issues

All errors include context and can be wrapped for additional information.

## Testing

Run the test suite:

```bash
go test ./reporting/...
```

Run with coverage:

```bash
go test -cover ./reporting/...
```

Run specific tests:

```bash
go test -v ./reporting/... -run TestValidator
go test -v ./reporting/... -run TestClient
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
