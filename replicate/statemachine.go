package replicate

// LifecycleObjectBasePaths maps object names that have lifecycleStatus to their basePath
// This map is generated from TMF OpenAPI v4 swagger definitions and can be used to:
// - Identify which objects support lifecycle management
// - Determine the correct API endpoint for lifecycle operations
// - Validate that objects have the proper lifecycle status fields
//
// The map contains 53 objects from various TMF APIs including:
// - Product Catalog Management (ProductOffering, ProductSpecification, etc.)
// - Service Catalog Management (ServiceSpecification, ServiceCandidate, etc.)
// - Resource Catalog (ResourceSpecification, ResourceCandidate, etc.)
// - Agreement Management (AgreementSpecification)
// - Usage Management (UsageSpecification)
var LifecycleObjectBasePaths = map[string]string{
	"AgreementSpecification":         "/tmf-api/agreementManagement/v4/",
	"BundledProductOffering":         "/tmf-api/productCatalogManagement/v4/",
	"BundledProductSpecification":    "/tmf-api/productCatalogManagement/v4/",
	"Catalog":                        "/tmf-api/productCatalogManagement/v4/",
	"Category":                       "/tmf-api/productCatalogManagement/v4/",
	"EntitySpecification":            "/tmf-api/usageManagement/v4/",
	"LogicalResourceSpecification":   "/tmf-api/resourceCatalog/v4/",
	"POPCharge":                      "/tmf-api/productCatalogManagement/v4/",
	"PhysicalResourceSpecification":  "/tmf-api/resourceCatalog/v4/",
	"ProductOffering":                "/tmf-api/productCatalogManagement/v4/",
	"ProductOfferingPrice":           "/tmf-api/productCatalogManagement/v4/",
	"ProductOfferingPriceRefOrValue": "/tmf-api/productCatalogManagement/v4/",
	"ProductSpecification":           "/tmf-api/productCatalogManagement/v4/",
	"ResourceCandidate":              "/tmf-api/resourceCatalog/v4/",
	"ResourceCatalog":                "/tmf-api/resourceCatalog/v4/",
	"ResourceCategory":               "/tmf-api/resourceCatalog/v4/",
	"ResourceFunctionSpecification":  "/tmf-api/resourceCatalog/v4/",
	"ResourceSpecification":          "/tmf-api/resourceCatalog/v4/",
	"ServiceCandidate":               "/tmf-api/serviceCatalogManagement/v4/",
	"ServiceCatalog":                 "/tmf-api/serviceCatalogManagement/v4/",
	"ServiceCategory":                "/tmf-api/serviceCatalogManagement/v4/",
	"ServiceSpecification":           "/tmf-api/serviceCatalogManagement/v4/",
	"UsageSpecification":             "/tmf-api/usageManagement/v4/",
}

var LifecycleStatusMandatory = map[string]string{
	"ProductOffering":       "In Study",
	"ProductOfferingPrice":  "In Study",
	"ProductSpecification":  "In Study",
	"ResourceSpecification": "In Study",
	"ServiceSpecification":  "In Study",
}
