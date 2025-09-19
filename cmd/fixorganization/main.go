package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/hesusruiz/isbetmf/tmfclient"
)

const (
	serverURL = "https://tmf.dome-marketplace-sbx.org"
	pageSize  = 100
)

func main() {
	// if len(os.Args) < 2 {
	// 	log.Fatalf("Usage: %s <access_token>", os.Args[0])
	// }
	// accessToken := os.Args[1]
	accessToken := ""

	clientCfg := &tmfclient.Config{
		BaseURL: serverURL,
		Timeout: 20,
	}
	client := tmfclient.NewClient(clientCfg)

	offset := 0
	for {
		path := fmt.Sprintf("/tmf-api/party/v4/organization?limit=%d&offset=%d", pageSize, offset)
		log.Printf("Fetching organizations from: %s", path)

		headers := map[string]string{
			// "Authorization": "Bearer " + accessToken,
		}

		resp, err := client.Get(path, headers)
		if err != nil {
			log.Fatalf("Failed to perform request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			bodyBytes, _ := io.ReadAll(resp.Body)
			log.Fatalf("Failed to get organizations. Status: %s, Body: %s", resp.Status, string(bodyBytes))
		}

		var organizations []map[string]any
		if err := json.NewDecoder(resp.Body).Decode(&organizations); err != nil {
			log.Fatalf("Failed to decode response body: %v", err)
		}

		if len(organizations) == 0 {
			break
		}

		log.Printf("Retrieved %d organizations", len(organizations))

		for _, org := range organizations {
			processOrganization(client, org, accessToken)
		}

		if len(organizations) < pageSize {
			break
		}

		offset += pageSize
	}

	log.Println("Finished processing all organizations.")
}

func processOrganization(client *tmfclient.Client, org map[string]any, accessToken string) {
	id, ok := org["id"].(string)
	if !ok {
		log.Printf("Organization without id: %v", org)
		return
	}

	if _, ok := org["organizationIdentification"]; ok {
		if orgIdent, ok := org["organizationIdentification"].([]any); ok && len(orgIdent) > 0 {
			log.Printf("Organization %s already has organizationIdentification", id)
			return
		}
	}

	log.Printf("Organization %s is missing organizationIdentification", id)

	extRefs, ok := org["externalReference"].([]any)
	if !ok || len(extRefs) == 0 {
		log.Printf("Organization %s has no externalReference", id)
		return
	}

	var orgIdentifier string
	for _, refAny := range extRefs {
		ref, ok := refAny.(map[string]any)
		if !ok {
			continue
		}
		if refType, ok := ref["externalReferenceType"].(string); ok && refType == "idm_id" {
			if name, ok := ref["name"].(string); ok {
				orgIdentifier = name
				break
			}
		}
	}

	if orgIdentifier == "" {
		log.Printf("Could not find organization identifier in externalReference for organization %s", id)
		return
	}

	log.Printf("Found organization identifier '%s' for organization %s", orgIdentifier, id)

	did := orgIdentifier
	if !strings.HasPrefix(did, "did:elsi:") {
		did = "did:elsi:" + did
	}

	newOrgIdent := map[string]any{
		"@type":              "organizationIdentification",
		"identificationId":   did,
		"identificationType": "did:elsi",
		"issuingAuthority":   "eIDAS",
	}

	org["organizationIdentification"] = []any{newOrgIdent}

	patchBodyMap := make(map[string]any)
	for k, v := range org {
		if k != "id" && k != "href" && k != "version" {
			patchBodyMap[k] = v
		}
	}

	patchBody, err := json.Marshal(patchBodyMap)
	if err != nil {
		log.Printf("Failed to marshal patch for organization %s: %v", id, err)
		return
	}

	patchURL := fmt.Sprintf("/tmf-api/party/v4/organization/%s", id)
	headers := map[string]string{
		// "Authorization": "Bearer " + accessToken,
		"Content-Type": "application/json",
	}

	resp, err := client.Patch(patchURL, patchBody, headers)
	if err != nil {
		log.Printf("Failed to send PATCH request for organization %s: %v", id, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		log.Printf("PATCH request for organization %s failed. Status: %s, Body: %s", id, resp.Status, string(bodyBytes))
		return
	}

	log.Printf("Successfully patched organization %s", id)
}
