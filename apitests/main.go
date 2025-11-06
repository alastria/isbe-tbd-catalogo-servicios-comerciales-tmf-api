package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	serverURL = "https://tmf.mycredential.eu/tmf-api/productCatalogManagement/v4"
	apiToken  = `eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJodHRwczovL2NhdGFsb2cuaXNiZW9uYm9hcmQuY29tIiwiZXhwIjoxNzYxMjI5NDE3LCJpYXQiOjE3NjEyMjU4MTcsImlzcyI6Imh0dHBzOi8vY2VydGF1dGguZXZpZGVuY2VsZWRnZXIuZXUiLCJqdGkiOiJUQ1RJVktTQjZQQTNTRlJYNVlCVUZPUkdZNyIsIm5vbmNlIjoiZDA4Mjg1N2MwNDIxYmViZTBkYTU5YTA5ZmYwYWVhM2NlOW1SczZVN1AiLCJzY29wZSI6Im9wZW5pZCBlaWRhcyIsInN1YiI6Imh0dHBzOi8vY2F0YWxvZy5pc2Jlb25ib2FyZC5jb20iLCJ2YyI6eyJAY29udGV4dCI6WyJodHRwczovL3d3dy53My5vcmcvbnMvY3JlZGVudGlhbHMvIiwiaHR0cHM6Ly9jcmVkZW50aWFscy5ldWRpc3RhY2suZXUvLndlbGwta25vd24vY3JlZGVudGlhbHMvbGVhcl9jcmVkZW50aWFsX2VtcGxveWVlL3czYy92MyJdLCJjcmVkZW50aWFsU3RhdHVzIjp7ImlkIjoiaHR0cHM6Ly9pc3N1ZXIuZG9tZS1tYXJrZXRwbGFjZS1zYngub3JnL2JhY2tvZmZpY2UvdjEvY3JlZGVudGlhbHMvc3RhdHVzLzEjU1lDOTA4UklRUXFlVVhSMTluaDNFUSIsInN0YXR1c0xpc3RDcmVkZW50aWFsIjoiaHR0cHM6Ly9pc3N1ZXIuZG9tZS1tYXJrZXRwbGFjZS1zYngub3JnL2JhY2tvZmZpY2UvdjEvY3JlZGVudGlhbHMvc3RhdHVzLzEiLCJzdGF0dXNMaXN0SW5kZXgiOiJTWUM5MDhSSVFRcWVVWFIxOW5oM0VRIiwic3RhdHVzUHVycG9zZSI6InJldm9jYXRpb24iLCJ0eXBlIjoiUGxhaW5MaXN0RW50aXR5In0sImNyZWRlbnRpYWxTdWJqZWN0Ijp7Im1hbmRhdGUiOnsibWFuZGF0ZWUiOnsiZW1haWwiOiJhbGJhLmxvcGV6QGluMi5lcyIsImVtcGxveWVlSWQiOiIxMjM0NTY3OEEiLCJmaXJzdE5hbWUiOiJKb2huIiwiaWQiOiIxMjM0NTY3OEEiLCJsYXN0TmFtZSI6IkRvZSJ9LCJtYW5kYXRvciI6eyJjb21tb25OYW1lIjoiSm9obiBEb2UiLCJjb3VudHJ5IjoiRVMiLCJlbWFpbCI6ImFsYmEubG9wZXpAaW4yLmVzIiwiaWQiOiJkaWQ6ZWxzaTpWQVRFUy0xMTExMTExMUsiLCJvcmdhbml6YXRpb24iOiJJU0JFIEZvdW5kYXRpb24iLCJvcmdhbml6YXRpb25JZGVudGlmaWVyIjoiVkFURVMtMTExMTExMTFLIiwic2VyaWFsTnVtYmVyIjoiMTIzNDU2NzhBIn0sInBvd2VyIjpbeyJhY3Rpb24iOlsiRXhlY3V0ZSJdLCJkb21haW4iOiJET01FIiwiZnVuY3Rpb24iOiJPbmJvYXJkaW5nIiwidHlwZSI6ImRvbWFpbiJ9LHsiYWN0aW9uIjpbIkNyZWF0ZSIsIlVwZGF0ZSIsIkRlbGV0ZSJdLCJkb21haW4iOiJET01FIiwiZnVuY3Rpb24iOiJQcm9kdWN0T2ZmZXJpbmciLCJ0eXBlIjoiZG9tYWluIn1dfX0sImRlc2NyaXB0aW9uIjoiVmVyaWZpYWJsZSBDcmVkZW50aWFsIGZvciBlbXBsb3llZXMgb2YgYW4gb3JnYW5pemF0aW9uIiwiaWQiOiJ1cm46dXVpZENXWFhTQzZDS1M3RElQV1hVNkM3NkFFSFpJIiwiaXNzdWVyIjp7ImNvbW1vbk5hbWUiOiJDZXJ0QXV0aCBJZGVudGl0eSBQcm92aWRlciBmb3IgSVNCRSIsImNvdW50cnkiOiJFUyIsImlkIjoiZGlkOmVsc2k6VkFURVMtQjYwNjQ1OTAwIiwib3JnYW5pemF0aW9uIjoiSU4yIiwib3JnYW5pemF0aW9uSWRlbnRpZmllciI6IlZBVEVTLUI2MDY0NTkwMCIsInNlcmlhbE51bWJlciI6IkI0NzQ0NzU2MCJ9LCJ0eXBlIjpbIkxFQVJDcmVkZW50aWFsRW1wbG95ZWUiLCJWZXJpZmlhYmxlQ3JlZGVudGlhbCJdLCJ2YWxpZEZyb20iOiIyMDI1LTEwLTIzVDEzOjA2OjIwWiIsInZhbGlkVW50aWwiOiIyMDI2LTEwLTIzVDEzOjA2OjIwWiJ9fQ.LSIItTupThpBDG08pWbAdIk0_qaG-U6w7SpbOPyjOUn-0wybXPwl-dyv8uRiEnbkLb0Mwgch-zkQGav3ImF4UA`
)

// Basic structure for a TMF API report
type TestReport struct {
	TestName     string
	Success      bool
	Message      string
	Timestamp    time.Time
	APIOperation string
	URL          string
	StatusCode   int
}

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
	reports := []TestReport{}

	// --- Test ProductSpecification ---
	fmt.Println("--- Testing ProductSpecification ---")

	// 1. Create (POST)
	specToCreate := ProductSpecification{
		Name:            "My Test Product Specification",
		Brand:           "TestBrand",
		Description:     "A detailed description of my test product specification.",
		LifecycleStatus: "Active",
	}
	createdSpec, report := createProductSpecification(specToCreate)
	reports = append(reports, report)
	if !report.Success {
		printReports(reports)
		return
	}
	fmt.Printf("Successfully created ProductSpecification with ID: %s\n", createdSpec.ID)

	// 2. Get (GET)
	retrievedSpec, report := getProductSpecification(createdSpec.ID)
	reports = append(reports, report)
	if !report.Success {
		printReports(reports)
		return
	}
	fmt.Printf("Successfully retrieved ProductSpecification with ID: %s\n", retrievedSpec.ID)

	// 3. Update (PATCH)
	updatedSpec, report := updateProductSpecification(createdSpec.ID, ProductSpecification{Description: "An updated description."})
	reports = append(reports, report)
	if !report.Success {
		printReports(reports)
		return
	}
	fmt.Printf("Successfully patched ProductSpecification with ID: %s\n", updatedSpec.ID)

	// 4. Delete (DELETE)
	report = deleteProductSpecification(createdSpec.ID)
	reports = append(reports, report)
	if !report.Success {
		printReports(reports)
		return
	}
	fmt.Printf("Successfully deleted ProductSpecification with ID: %s\n", createdSpec.ID)

	// 5. List (GET)

	fmt.Println("--- Testing ProductSpecification List ---")

	// Create two new product specifications to test the list functionality

	specToCreate1 := ProductSpecification{

		Name: "My First List Test Product Specification",

		Brand: "TestBrand",

		Description: "First item for list test.",

		LifecycleStatus: "Active",
	}

	createdSpec1, report := createProductSpecification(specToCreate1)

	reports = append(reports, report)

	if !report.Success {

		printReports(reports)

		return

	}

	fmt.Printf("Successfully created first list ProductSpecification with ID: %s\n", createdSpec1.ID)

	specToCreate2 := ProductSpecification{

		Name: "My Second List Test Product Specification",

		Brand: "TestBrand",

		Description: "Second item for list test.",

		LifecycleStatus: "Active",
	}

	createdSpec2, report := createProductSpecification(specToCreate2)

	reports = append(reports, report)

	if !report.Success {

		printReports(reports)

		return

	}

	fmt.Printf("Successfully created second list ProductSpecification with ID: %s\n", createdSpec2.ID)

	specs, report := listProductSpecifications()

	reports = append(reports, report)

	if !report.Success {

		printReports(reports)

		return

	}

	if len(specs) < 2 {

		report.Success = false

		report.Message = fmt.Sprintf("Expected at least 2 product specifications, but got %d", len(specs))

		reports = append(reports, report)

		printReports(reports)

		return

	}

	fmt.Printf("Successfully listed ProductSpecifications, found %d items.\n", len(specs))

	// Cleanup

	report = deleteProductSpecification(createdSpec1.ID)

	reports = append(reports, report)

	if !report.Success {

		printReports(reports)

		return

	}

	fmt.Printf("Successfully deleted first list ProductSpecification with ID: %s\n", createdSpec1.ID)

	report = deleteProductSpecification(createdSpec2.ID)

	reports = append(reports, report)

	if !report.Success {

		printReports(reports)

		return

	}

	fmt.Printf("Successfully deleted second list ProductSpecification with ID: %s\n", createdSpec2.ID)

	printReports(reports)

}

func createProductSpecification(spec ProductSpecification) (ProductSpecification, TestReport) {
	report := TestReport{
		TestName:     "Create ProductSpecification",
		Timestamp:    time.Now(),
		APIOperation: "POST",
		URL:          fmt.Sprintf("%s/productSpecification", serverURL),
	}

	jsonData, err := json.Marshal(spec)
	if err != nil {
		report.Success = false
		report.Message = fmt.Sprintf("Failed to marshal JSON: %v", err)
		return ProductSpecification{}, report
	}

	req, err := http.NewRequest("POST", report.URL, bytes.NewBuffer(jsonData))
	if err != nil {
		report.Success = false
		report.Message = fmt.Sprintf("Failed to create request: %v", err)
		return ProductSpecification{}, report
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		report.Success = false
		report.Message = fmt.Sprintf("Failed to send request: %v", err)
		return ProductSpecification{}, report
	}
	defer resp.Body.Close()

	report.StatusCode = resp.StatusCode

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		report.Success = false
		report.Message = fmt.Sprintf("Failed to read response body: %v", err)
		return ProductSpecification{}, report
	}

	if resp.StatusCode != http.StatusCreated {
		report.Success = false
		report.Message = fmt.Sprintf("Expected status code 201, but got %d. Body: %s", resp.StatusCode, string(body))
		return ProductSpecification{}, report
	}

	var createdSpec ProductSpecification
	if err := json.Unmarshal(body, &createdSpec); err != nil {
		report.Success = false
		report.Message = fmt.Sprintf("Failed to unmarshal response body: %v", err)
		return ProductSpecification{}, report
	}

	report.Success = true
	report.Message = "ProductSpecification created successfully."
	return createdSpec, report
}

func getProductSpecification(id string) (ProductSpecification, TestReport) {
	report := TestReport{
		TestName:     "Get ProductSpecification",
		Timestamp:    time.Now(),
		APIOperation: "GET",
		URL:          fmt.Sprintf("%s/productSpecification/%s", serverURL, id),
	}

	req, err := http.NewRequest("GET", report.URL, nil)
	if err != nil {
		report.Success = false
		report.Message = fmt.Sprintf("Failed to create request: %v", err)
		return ProductSpecification{}, report
	}

	req.Header.Set("Authorization", "Bearer "+apiToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		report.Success = false
		report.Message = fmt.Sprintf("Failed to send request: %v", err)
		return ProductSpecification{}, report
	}
	defer resp.Body.Close()

	report.StatusCode = resp.StatusCode

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		report.Success = false
		report.Message = fmt.Sprintf("Failed to read response body: %v", err)
		return ProductSpecification{}, report
	}

	if resp.StatusCode != http.StatusOK {
		report.Success = false
		report.Message = fmt.Sprintf("Expected status code 200, but got %d. Body: %s", resp.StatusCode, string(body))
		return ProductSpecification{}, report
	}

	var retrievedSpec ProductSpecification
	if err := json.Unmarshal(body, &retrievedSpec); err != nil {
		report.Success = false
		report.Message = fmt.Sprintf("Failed to unmarshal response body: %v", err)
		return ProductSpecification{}, report
	}

	report.Success = true
	report.Message = "ProductSpecification retrieved successfully."
	return retrievedSpec, report
}

func updateProductSpecification(id string, spec ProductSpecification) (ProductSpecification, TestReport) {
	report := TestReport{
		TestName:     "Update ProductSpecification",
		Timestamp:    time.Now(),
		APIOperation: "PATCH",
		URL:          fmt.Sprintf("%s/productSpecification/%s", serverURL, id),
	}

	jsonData, err := json.Marshal(spec)
	if err != nil {
		report.Success = false
		report.Message = fmt.Sprintf("Failed to marshal JSON: %v", err)
		return ProductSpecification{}, report
	}

	req, err := http.NewRequest("PATCH", report.URL, bytes.NewBuffer(jsonData))
	if err != nil {
		report.Success = false
		report.Message = fmt.Sprintf("Failed to create request: %v", err)
		return ProductSpecification{}, report
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		report.Success = false
		report.Message = fmt.Sprintf("Failed to send request: %v", err)
		return ProductSpecification{}, report
	}
	defer resp.Body.Close()

	report.StatusCode = resp.StatusCode

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		report.Success = false
		report.Message = fmt.Sprintf("Failed to read response body: %v", err)
		return ProductSpecification{}, report
	}

	if resp.StatusCode != http.StatusOK {
		report.Success = false
		report.Message = fmt.Sprintf("Expected status code 200, but got %d. Body: %s", resp.StatusCode, string(body))
		return ProductSpecification{}, report
	}

	var updatedSpec ProductSpecification
	if err := json.Unmarshal(body, &updatedSpec); err != nil {
		report.Success = false
		report.Message = fmt.Sprintf("Failed to unmarshal response body: %v", err)
		return ProductSpecification{}, report
	}

	report.Success = true
	report.Message = "ProductSpecification updated successfully."
	return updatedSpec, report
}

func deleteProductSpecification(id string) TestReport {
	report := TestReport{
		TestName:     "Delete ProductSpecification",
		Timestamp:    time.Now(),
		APIOperation: "DELETE",
		URL:          fmt.Sprintf("%s/productSpecification/%s", serverURL, id),
	}

	req, err := http.NewRequest("DELETE", report.URL, nil)
	if err != nil {
		report.Success = false
		report.Message = fmt.Sprintf("Failed to create request: %v", err)
		return report
	}

	req.Header.Set("Authorization", "Bearer "+apiToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		report.Success = false
		report.Message = fmt.Sprintf("Failed to send request: %v", err)
		return report
	}
	defer resp.Body.Close()

	report.StatusCode = resp.StatusCode

	if resp.StatusCode != http.StatusNoContent {
		report.Success = false
		report.Message = fmt.Sprintf("Expected status code 204, but got %d", resp.StatusCode)
		return report
	}

	report.Success = true
	report.Message = "ProductSpecification deleted successfully."
	return report
}

func listProductSpecifications() ([]ProductSpecification, TestReport) {
	report := TestReport{
		TestName:     "List ProductSpecifications",
		Timestamp:    time.Now(),
		APIOperation: "GET",
		URL:          fmt.Sprintf("%s/productSpecification", serverURL),
	}

	req, err := http.NewRequest("GET", report.URL, nil)
	if err != nil {
		report.Success = false
		report.Message = fmt.Sprintf("Failed to create request: %v", err)
		return nil, report
	}

	req.Header.Set("Authorization", "Bearer "+apiToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		report.Success = false
		report.Message = fmt.Sprintf("Failed to send request: %v", err)
		return nil, report
	}
	defer resp.Body.Close()

	report.StatusCode = resp.StatusCode

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		report.Success = false
		report.Message = fmt.Sprintf("Failed to read response body: %v", err)
		return nil, report
	}

	if resp.StatusCode != http.StatusOK {
		report.Success = false
		report.Message = fmt.Sprintf("Expected status code 200, but got %d. Body: %s", resp.StatusCode, string(body))
		return nil, report
	}

	var specs []ProductSpecification
	if err := json.Unmarshal(body, &specs); err != nil {
		report.Success = false
		report.Message = fmt.Sprintf("Failed to unmarshal response body: %v", err)
		return nil, report
	}

	report.Success = true
	report.Message = "ProductSpecifications listed successfully."
	return specs, report
}

func printReports(reports []TestReport) {
	fmt.Println("\n--- Test Report ---")
	for _, r := range reports {
		status := "SUCCESS"
		if !r.Success {
			status = "FAILURE"
		}
		fmt.Printf("[%s] %s: %s\n", status, r.TestName, r.Message)
		fmt.Printf("  - Operation: %s %s\n", r.APIOperation, r.URL)
		fmt.Printf("  - Timestamp: %s\n", r.Timestamp.Format(time.RFC3339))
		if r.StatusCode > 0 {
			fmt.Printf("  - Status Code: %d\n", r.StatusCode)
		}
	}
	fmt.Println("-------------------")
}
