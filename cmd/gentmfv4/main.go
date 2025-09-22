package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/goccy/go-yaml"
	"github.com/hesusruiz/isbetmf/internal/jpath"
	"golang.org/x/tools/imports"
)

// This is a simple tool to process the Swagger files in the "swagger" directory
// and extract the mapping of last path part to management system and the routes.
// It assumes the Swagger files are in the format used by the TMForum APIs.
// It will print the mapping and the routes to the standard output in JSON format.

//go:embed routes.hbs
var routesTemplate string

var (
	managementToUpstream = map[string]string{}
	resourceToManagement = map[string]string{}
	resourceToPath       = map[string]string{}
)

func main() {

	// Visit recursively the directories in the "swagger" directory
	// It assumes an "almost" flat structure with directories named after the management system
	// and one file inside each directory named "api.json" or similar.
	baseDir := "./www/oapiv4/oapiv4"

	// Read the directory entries
	dirEntries, err := os.ReadDir(baseDir)
	if err != nil {
		panic(err)
	}

	// Process each file in the directory
	for _, dirEntry := range dirEntries {
		if !dirEntry.IsDir() {
			// Process the file
			filePath := path.Join(baseDir, dirEntry.Name())
			if !strings.HasSuffix(filePath, ".json") {
				// Skip non-JSON files
				continue
			}
			processOneFile(filePath)
		}
	}

	if true {
		os.Exit(0)
	}

	tmpl, err := template.New("routes").Parse(routesTemplate)
	if err != nil {
		panic(err)
	}

	var b bytes.Buffer
	err = tmpl.Execute(&b, map[string]any{
		"ResourceToManagement":   resourceToManagement,
		"ResourceToStandardPath": resourceToPath,
		"ManagementToUpstream":   managementToUpstream,
	})
	if err != nil {
		panic(err)
	}

	// Adjust imports before writing the file
	out, err := imports.Process("types.go", b.Bytes(), nil)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("./types.go", out, 0644)
	if err != nil {
		panic(err)
	}

}

func processOneFile(filePath string) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	// Unmarshal the JSON content into a map
	var swagMap map[string]any
	err = yaml.Unmarshal(content, &swagMap)
	if err != nil {
		panic(err)
	}

	title := jpath.GetString(swagMap, "info.title")
	if len(title) == 0 {
		panic("info.title key not found or empty")
	}

	basePath := jpath.GetString(swagMap, "basePath")
	if len(basePath) == 0 {
		panic("basePath key not found or empty")
	}

	basePathTrimmed := strings.Trim(basePath, "/")

	basePathParts := strings.Split(basePathTrimmed, "/")
	if len(basePathParts) < 2 {
		panic("basePath does not contain at least 2 parts")
	}

	apiFamily := basePathParts[1]
	fmt.Println("API family:", apiFamily)
	managementToUpstream[apiFamily] = "TODO: set upstream host and path like http://localhost:8062"

	// Get the "paths" key from the map
	paths := jpath.GetMap(swagMap, "paths")

	localResourceNames := map[string]bool{}

	// Iterate over the keys in the "paths" map
	for thePath := range paths {
		// Check if the value is a map
		// methodsMap, ok := methods.(map[string]any)
		// if !ok {
		// 	panic("methods value is not a map")
		// }

		thePath = strings.Trim(thePath, "/")

		pathParts := strings.Split(thePath, "/")
		resourceName := pathParts[0]

		if resourceName == "importJob" || resourceName == "exportJob" {
			// We do not implement these APIs
			continue
		}

		if resourceName == "hub" || resourceName == "listener" {
			// TODO: implement specia processing for these paths
			continue
		}

		localResourceNames[strings.ToLower(resourceName)] = true
		resourceToManagement[resourceName] = apiFamily
		resourceToPath[resourceName] = path.Join("/tmf-api", basePath, resourceName)

	}

	// fmt.Println(description)
	fmt.Printf("- **%s**\n", apiFamily)

	for resourceName := range localResourceNames {
		fmt.Printf("  - %s\n", resourceName)
	}

	fmt.Println()

}
