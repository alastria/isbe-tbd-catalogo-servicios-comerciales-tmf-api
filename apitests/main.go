package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

const (
	serverURL = "http://tmf.mycredential.eu/tmf-api/productCatalogManagement/v4"
	apiToken  = "fake-token"
)

// ProductSpecification represents the structure for a product specification.
// This is a simplified version based on the TMF620 swagger file.
type ProductSpecification struct {
	ID              string `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	Href            string `json:"href,omitempty"`
	Brand           string `json:"brand,omitempty"`
	Description     string `json:"description,omitempty"`
	LastUpdate      string `json:"lastUpdate,omitempty"`
	LifecycleStatus string `json:"lifecycleStatus,omitempty"`
	Version         string `json:"version,omitempty"`
}

func main() {
	// httpexpect requires a *testing.T instance. We can create a mock one.
	t := &testing.T{}

	// Create a new httpexpect instance.
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  serverURL,
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewCurlPrinter(t),
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	// --- Test ProductSpecification ---
	fmt.Println("--- Testing ProductSpecification ---")

	// 1. Create (POST)
	specToCreate := ProductSpecification{
		Name:            "My Test Product Specification",
		Brand:           "TestBrand",
		Description:     "A detailed description of my test product specification.",
		LifecycleStatus: "Active",
	}

	createdSpecObj := e.POST("/productSpecification").
		WithHeader("Authorization", "Bearer "+apiToken).
		WithJSON(specToCreate).
		Expect().
		Status(http.StatusCreated).
		JSON().Object()

	createdSpecObj.Value("id").String().NotEmpty()
	createdSpecObj.Value("name").String().IsEqual(specToCreate.Name)
	createdSpecObj.Value("brand").String().IsEqual(specToCreate.Brand)

	createdSpecID := createdSpecObj.Value("id").String().Raw()
	fmt.Printf("Successfully created ProductSpecification with ID: %s\n", createdSpecID)

	// 2. Get (GET)
	e.GET("/productSpecification/{id}", createdSpecID).
		WithHeader("Authorization", "Bearer "+apiToken).
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		Value("id").String().IsEqual(createdSpecID)
	fmt.Printf("Successfully retrieved ProductSpecification with ID: %s\n", createdSpecID)

	// 3. Update (PATCH)
	updatePayload := map[string]interface{}{"description": "An updated description."}
	e.PATCH("/productSpecification/{id}", createdSpecID).
		WithHeader("Authorization", "Bearer "+apiToken).
		WithJSON(updatePayload).
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		Value("description").String().IsEqual("An updated description.")
	fmt.Printf("Successfully patched ProductSpecification with ID: %s\n", createdSpecID)

	// 4. Delete (DELETE)
	e.DELETE("/productSpecification/{id}", createdSpecID).
		WithHeader("Authorization", "Bearer "+apiToken).
		Expect().
		Status(http.StatusNoContent)
	fmt.Printf("Successfully deleted ProductSpecification with ID: %s\n", createdSpecID)

	// 5. List (GET)
	fmt.Println("--- Testing ProductSpecification List ---")

	// Create two new product specifications to test the list functionality
	spec1 := ProductSpecification{
		Name:  "List Spec 1",
		Brand: "TestBrand",
	}
	spec2 := ProductSpecification{
		Name:  "List Spec 2",
		Brand: "TestBrand",
	}

	createdSpec1Obj := e.POST("/productSpecification").WithHeader("Authorization", "Bearer "+apiToken).WithJSON(spec1).Expect().Status(http.StatusCreated).JSON().Object()
	createdSpec2Obj := e.POST("/productSpecification").WithHeader("Authorization", "Bearer "+apiToken).WithJSON(spec2).Expect().Status(http.StatusCreated).JSON().Object()

	createdSpec1ID := createdSpec1Obj.Value("id").String().Raw()
	createdSpec2ID := createdSpec2Obj.Value("id").String().Raw()

	fmt.Printf("Successfully created list ProductSpecification 1 with ID: %s\n", createdSpec1ID)
	fmt.Printf("Successfully created list ProductSpecification 2 with ID: %s\n", createdSpec2ID)

	// Retrieve list and assert it contains at least two items
	list := e.GET("/productSpecification").
		WithHeader("Authorization", "Bearer "+apiToken).
		Expect().
		Status(http.StatusOK).
		JSON().Array()

	list.Length().Ge(2)
	fmt.Printf("Successfully listed ProductSpecifications, found %d items.\n", len(list.Raw()))

	// Cleanup
	e.DELETE("/productSpecification/{id}", createdSpec1ID).WithHeader("Authorization", "Bearer "+apiToken).Expect().Status(http.StatusNoContent)
	e.DELETE("/productSpecification/{id}", createdSpec2ID).WithHeader("Authorization", "Bearer "+apiToken).Expect().Status(http.StatusNoContent)

	fmt.Printf("Successfully deleted list ProductSpecifications with IDs: %s, %s\n", createdSpec1ID, createdSpec2ID)

	fmt.Println("\n--- All tests passed ---")
}
