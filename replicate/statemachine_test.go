package replicate

import (
	"testing"
)

func TestLifecycleObjectBasePaths(t *testing.T) {
	// Test that the map is not empty
	if len(LifecycleObjectBasePaths) == 0 {
		t.Error("LifecycleObjectBasePaths map should not be empty")
	}

	// Test some expected objects and their base paths
	testCases := []struct {
		objectName       string
		expectedBasePath string
	}{
		{"ProductOffering", "/tmf-api/productCatalogManagement/v4/"},
		{"ServiceSpecification", "/tmf-api/serviceCatalogManagement/v4/"},
		{"ResourceSpecification", "/tmf-api/resourceCatalog/v4/"},
		{"AgreementSpecification", "/tmf-api/agreementManagement/v4/"},
		{"UsageSpecification", "/tmf-api/usageManagement/v4/"},
	}

	for _, tc := range testCases {
		if basePath, exists := LifecycleObjectBasePaths[tc.objectName]; !exists {
			t.Errorf("Expected object %s to exist in LifecycleObjectBasePaths", tc.objectName)
		} else if basePath != tc.expectedBasePath {
			t.Errorf("Expected base path %s for object %s, got %s", tc.expectedBasePath, tc.objectName, basePath)
		}
	}

	// Test that all base paths start with "/tmf-api/"
	for objectName, basePath := range LifecycleObjectBasePaths {
		if len(basePath) == 0 {
			t.Errorf("Base path for object %s should not be empty", objectName)
		}
		if !startsWith(basePath, "/tmf-api/") {
			t.Errorf("Base path for object %s should start with '/tmf-api/', got: %s", objectName, basePath)
		}
	}
}

// Helper function to check if a string starts with a prefix
func startsWith(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}
