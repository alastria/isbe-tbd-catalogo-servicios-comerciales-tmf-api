// Copyright 2025 Jesus Ruiz
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"net/http"
	"testing"
	"time"

	"github.com/gavv/httpexpect/v2"
	"github.com/hesusruiz/isbetmf/config"
	repository "github.com/hesusruiz/isbetmf/tmfserver/repository"
	"github.com/hesusruiz/isbetmf/tmfserver/service"
)

const (
	serverURL = "https://tmf.mycredential.eu/tmf-api/productCatalogManagement/v4"
)

var (
	apiToken = service.FakeATold
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

func init() {
	// Generate a default configuration suitable for the environment
	// The approach is that instead of many configurable parameters, we have a set of profiles, with "hardcoded"
	// parameters for each environment, but that can be easity extended for other purposes.
	configuration, err := config.LoadConfig("mycredential", true)
	if err != nil {
		panic(err)
	}

	// Disable restarts
	configuration.RestartHour = -1

	go runNormalProcess(configuration)

	time.Sleep(1 * time.Second)

}

func TestNoAuthorization(t *testing.T) {

	// Create a new httpexpect instance.
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  serverURL,
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewCurlPrinter(t),
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	// 1. Create (POST) without authentication
	ps := repository.TMFObjectMap{
		"name":            "My Test Product Specification",
		"brand":           "TestBrand",
		"description":     "A detailed description of my test product specification.",
		"lifecycleStatus": "Active",
	}

	_ = e.POST("/productSpecification").
		WithJSON(ps).
		Expect().
		Status(http.StatusUnauthorized).
		JSON().Object()

}

func TestProductSpecificationHappy(t *testing.T) {

	// Create a new httpexpect instance.
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  serverURL,
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewCurlPrinter(t),
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	// 1. Create (POST)
	ps := repository.TMFObjectMap{
		"name":            "My Test Product Specification",
		"brand":           "TestBrand",
		"description":     "A detailed description of my test product specification.",
		"lifecycleStatus": "Active",
	}

	createdSpecObj := e.POST("/productSpecification").
		WithHeader("Authorization", "Bearer "+apiToken).
		WithJSON(ps).
		Expect().
		Status(http.StatusCreated).
		JSON().Object()

	createdSpecObj.Value("id").String().NotEmpty()
	createdSpecObj.Value("name").String().IsEqual(ps.Name())
	createdSpecObj.Value("brand").String().IsEqual(ps.GetStringField("brand"))

	createdSpecID := createdSpecObj.Value("id").String().Raw()

	// 2. Get (GET)
	e.GET("/productSpecification/{id}", createdSpecID).
		WithHeader("Authorization", "Bearer "+apiToken).
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		Value("id").String().IsEqual(createdSpecID)

	// 3. Update (PATCH)
	updatePayload := map[string]any{"description": "An updated description."}
	e.PATCH("/productSpecification/{id}", createdSpecID).
		WithHeader("Authorization", "Bearer "+apiToken).
		WithJSON(updatePayload).
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		Value("description").String().IsEqual("An updated description.")

	// 4. Delete (DELETE)
	e.DELETE("/productSpecification/{id}", createdSpecID).
		WithHeader("Authorization", "Bearer "+apiToken).
		Expect().
		Status(http.StatusNoContent)

	// 5. List (GET)

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

	// Retrieve list and assert it contains at least two items
	list := e.GET("/productSpecification").
		WithHeader("Authorization", "Bearer "+apiToken).
		Expect().
		Status(http.StatusOK).
		JSON().Array()

	list.Length().Ge(2)

	// Cleanup
	e.DELETE("/productSpecification/{id}", createdSpec1ID).WithHeader("Authorization", "Bearer "+apiToken).Expect().Status(http.StatusNoContent)
	e.DELETE("/productSpecification/{id}", createdSpec2ID).WithHeader("Authorization", "Bearer "+apiToken).Expect().Status(http.StatusNoContent)

}

func TestInvalidSeller(t *testing.T) {

	// Create a new httpexpect instance.
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  serverURL,
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewCurlPrinter(t),
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	// 1. Create (POST)
	ps := repository.TMFObjectMap{
		"name":            "My Test Product Specification",
		"brand":           "TestBrand",
		"description":     "A detailed description of my test product specification.",
		"lifecycleStatus": "Active",
	}

	ps.SetSellerInfo("pepe", "juan", "v4")

	createdSpecObj := e.POST("/productSpecification").
		WithHeader("Authorization", "Bearer "+apiToken).
		WithJSON(ps).
		Expect().
		Status(http.StatusCreated).
		JSON().Object()

	createdSpecObj.Value("id").String().NotEmpty()
	createdSpecObj.Value("name").String().IsEqual(ps.Name())
	createdSpecObj.Value("brand").String().IsEqual(ps.GetStringField("brand"))

	createdSpecObj.Value("relatedParty").Array().Length().Ge(2)

}
