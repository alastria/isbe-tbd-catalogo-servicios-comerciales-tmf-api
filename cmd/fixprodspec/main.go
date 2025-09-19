package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/hesusruiz/isbetmf/tmfclient"
)

const (
	serverURL = "https://tmf.dome-marketplace-sbx.org"
	pageSize  = 100
)

func main() {
	accessToken := ""
	if len(os.Args) >= 2 {
		accessToken = os.Args[1]
	}

	clientCfg := &tmfclient.Config{
		BaseURL: serverURL,
		Timeout: 20,
	}
	client := tmfclient.NewClient(clientCfg)

	offset := 0
	for {
		path := fmt.Sprintf("/tmf-api/productCatalogManagement/v4/productSpecification?limit=%d&offset=%d", pageSize, offset)
		log.Printf("Fetching product specifications from: %s", path)

		headers := map[string]string{}
		if accessToken != "" {
			headers["Authorization"] = "Bearer " + accessToken
		}

		resp, err := client.Get(path, headers)
		if err != nil {
			log.Fatalf("Failed to perform request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			bodyBytes, _ := io.ReadAll(resp.Body)
			log.Fatalf("Failed to get product specifications. Status: %s, Body: %s", resp.Status, string(bodyBytes))
		}

		var productSpecs []map[string]any
		if err := json.NewDecoder(resp.Body).Decode(&productSpecs); err != nil {
			log.Fatalf("Failed to decode response body: %v", err)
		}

		if len(productSpecs) == 0 {
			break
		}

		log.Printf("Retrieved %d product specifications", len(productSpecs))

		for _, ps := range productSpecs {
			processProductSpecification(client, ps, accessToken)
		}

		if len(productSpecs) < pageSize {
			break
		}

		offset += pageSize
	}

	log.Println("Finished processing all product specifications.")
}

func processProductSpecification(client *tmfclient.Client, ps map[string]any, accessToken string) {
	id, ok := ps["id"].(string)
	if !ok {
		log.Printf("Product specification without id: %v", ps)
		return
	}

	relatedParties, ok := ps["relatedParty"].([]any)
	if !ok {
		relatedParties = []any{}
	}

	var ownerID string
	var hasSeller bool
	var hasSellerOperator bool
	for _, rpAny := range relatedParties {
		rp, ok := rpAny.(map[string]any)
		if !ok {
			continue
		}
		role, _ := rp["role"].(string)
		if role == "Owner" {
			if rpID, ok := rp["id"].(string); ok && strings.HasPrefix(rpID, "urn:ngsi-ld:organization:") {
				ownerID = rpID
			}
		}
		if role == "Seller" {
			hasSeller = true
		}
		if role == "SellerOperator" {
			hasSellerOperator = true
		}
	}

	dirty := false

	if !hasSellerOperator {
		log.Printf("Product specification %s is missing SellerOperator, adding it.", id)
		newSellerOperator := map[string]any{
			"@referredType": "Organization",
			"href":          "urn:ngsi-ld:organization:df924e5d-e8c8-4ea4-aca8-edaf5acdc109",
			"id":            "urn:ngsi-ld:organization:df924e5d-e8c8-4ea4-aca8-edaf5acdc109",
			"name":          "did:elsi:VATSB-12345678J",
			"role":          "SellerOperator",
		}
		relatedParties = append(relatedParties, newSellerOperator)
		dirty = true
	}

	if !hasSeller {
		log.Printf("Product specification %s is missing Seller", id)
		if ownerID == "" {
			log.Printf("Product specification %s has no Owner with organization id, cannot add Seller", id)
		} else {
			// Get the organization to find its identificationId
			orgPath := fmt.Sprintf("/tmf-api/party/v4/organization/%s", ownerID)

			headers := map[string]string{}
			if accessToken != "" {
				headers["Authorization"] = "Bearer " + accessToken
			}

			resp, err := client.Get(orgPath, headers)
			if err != nil {
				log.Printf("Failed to get organization %s: %v", ownerID, err)
			} else {
				defer resp.Body.Close()

				if resp.StatusCode != http.StatusOK {
					log.Printf("Failed to get organization %s, status: %s", ownerID, resp.Status)
				} else {
					var org map[string]any
					if err := json.NewDecoder(resp.Body).Decode(&org); err != nil {
						log.Printf("Failed to decode organization %s: %v", ownerID, err)
					} else {
						orgIdentifications, ok := org["organizationIdentification"].([]any)
						if !ok || len(orgIdentifications) == 0 {
							log.Printf("Organization %s has no organizationIdentification", ownerID)
						} else {
							orgIdent, ok := orgIdentifications[0].(map[string]any)
							if !ok {
								log.Printf("Invalid organizationIdentification for organization %s", ownerID)
							} else {
								identificationId, ok := orgIdent["identificationId"].(string)
								if !ok {
									log.Printf("Organization %s has no identificationId in organizationIdentification", ownerID)
								} else {
									newSeller := map[string]any{
										"@referredType": "Organization",
										"href":          ownerID,
										"id":            ownerID,
										"name":          identificationId,
										"role":          "Seller",
									}
									relatedParties = append(relatedParties, newSeller)
									dirty = true
								}
							}
						}
					}
				}
			}
		}
	}

	if !dirty {
		log.Printf("Product specification %s is up to date.", id)
		return
	}

	ps["relatedParty"] = relatedParties

	patchBodyMap := make(map[string]any)
	for k, v := range ps {
		if k != "id" && k != "href" && k != "version" && k != "lastUpdate" {
			patchBodyMap[k] = v
		}
	}

	patchBody, err := json.Marshal(patchBodyMap)
	if err != nil {
		log.Printf("Failed to marshal patch for product specification %s: %v", id, err)
		return
	}

	patchURL := fmt.Sprintf("/tmf-api/productCatalogManagement/v4/productSpecification/%s", id)
	patchHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	if accessToken != "" {
		patchHeaders["Authorization"] = "Bearer " + accessToken
	}

	patchResp, err := client.Patch(patchURL, patchBody, patchHeaders)
	if err != nil {
		log.Printf("Failed to send PATCH request for product specification %s: %v", id, err)
		return
	}
	defer patchResp.Body.Close()

	if patchResp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(patchResp.Body)
		log.Printf("PATCH request for product specification %s failed. Status: %s, Body: %s", id, patchResp.Status, string(bodyBytes))
		return
	}

	log.Printf("Successfully patched product specification %s", id)
}
