package repository

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/hesusruiz/isbetmf/config"
	"github.com/hesusruiz/isbetmf/internal/errl"
	"github.com/hesusruiz/isbetmf/types"
)

const exampleIndividualObject = `
{
  "@type": "individual",
  "contactMedium": [],
  "countryOfBirth": "",
  "externalReference": [
    {
      "externalReferenceType": "idm_id",
      "name": "did:key:zDnaeqy64MRbQfM9TNMmfBrsFnbsxGXD5kYihyxbC4PXkmnAo"
    }
  ],
  "familyName": "Alba Lopez",
  "gender": "",
  "givenName": "Alba Lopez",
  "href": "/tmf-api/party/v4/individual/urn:ngsi-ld:individual:60746ac1-f523-4a0b-8106-d9f5b77b47ad",
  "id": "urn:ngsi-ld:individual:60746ac1-f523-4a0b-8106-d9f5b77b47ad",
  "lastUpdate": "2025-10-15T09:06:27.383980245Z",
  "maritalStatus": "",
  "nationality": "",
  "partyCharacteristic": [],
  "placeOfBirth": "",
  "relatedParty": [
    {
      "@referredType": null,
      "href": "urn:ngsi-ld:organization:2bcbe859-e316-42f2-919c-f470cff9e235",
      "id": "urn:ngsi-ld:organization:2bcbe859-e316-42f2-919c-f470cff9e235",
      "name": "IN2 INGENIERIA DE LA INFORMACION SOCIEDAD LIMITADA",
      "role": "orgAdmin"
    },
    {
      "@referredType": "Organization",
      "href": "urn:ngsi-ld:organization:did:elsi:VATES-B60645900",
      "id": "urn:ngsi-ld:organization:did:elsi:VATES-B60645900",
      "name": "did:elsi:VATES-B60645900",
      "role": "Seller"
    },
    {
      "@referredType": "Organization",
      "href": "urn:ngsi-ld:organization:did:elsi:VATES-11111111K",
      "id": "urn:ngsi-ld:organization:did:elsi:VATES-11111111K",
      "name": "did:elsi:VATES-11111111K",
      "role": "SellerOperator"
    }
  ],
  "title": "",
  "version": "1.0"
}`

func TMFIndividualFromCredential(verifiableCredential map[string]any, user *Organization) (*TMFObject, error) {

	// TODO: what happens if we receive a LEARCredentialMachine?
	lc, err := types.LEARCredentialFromMap(verifiableCredential)
	if err != nil {
		return nil, errl.Errorf("error creating LEARCredentialEmployee struct: %w", err)
	}

	now := time.Now()
	lastUpdate := now.Format(time.RFC3339Nano)

	userId := lc.CredentialSubject.Mandate.Mandatee.Id
	if userId == "" {
		return nil, errl.Errorf("id field for mandatee is empty")
	}

	// The format is "urn:ngsi-ld:individual:{organizationIdentifier}:{uuid}"
	id := fmt.Sprintf("urn:ngsi-ld:individual:%s:%s", user.OrganizationIdentifier, userId)

	objectType := config.Individual
	version := "1.0"

	// Prepare contactMedium
	var contactMedium []any
	if lc.CredentialSubject.Mandate.Mandatee.Email != "" {
		contactMedium = append(contactMedium, map[string]any{
			"@type":        "EmailContactMedium",
			"preferred":    true,
			"emailAddress": user.EmailAddress,
		})
	}

	// Check if the powers are "Onboarding"
	hasOnboardingPower := false
	powers := lc.CredentialSubject.Mandate.Power
	for _, p := range powers {
		if strings.EqualFold(p.Tmf_type, "Domain") && strings.EqualFold(p.Tmf_domain, "DOME") {

			if strings.EqualFold(p.Tmf_function, "Onboarding") {

				for _, action := range p.Tmf_action {
					if strings.EqualFold(action, "execute") {
						hasOnboardingPower = true
						break
					}
				}

			}

		}
	}

	relatedParty := []any{}
	if hasOnboardingPower {
		relatedParty = append(relatedParty, map[string]any{
			"@referredType": nil,
			"href":          "urn:ngsi-ld:organization:" + user.OrganizationIdentifier,
			"id":            "urn:ngsi-ld:organization:" + user.OrganizationIdentifier,
			"name":          user.Organization,
			"role":          "orgAdmin",
		},
		)
	}

	indMap := map[string]any{
		"@type":      config.Individual,
		"id":         id,
		"href":       id,
		"version":    version,
		"lastUpdate": lastUpdate,
		"externalReference": []any{
			map[string]any{
				"externalReferenceType": "idm_id",
				"name":                  lc.CredentialSubject.Mandate.Mandatee.Id,
			},
		},
		"familyName": lc.CredentialSubject.Mandate.Mandatee.LastName,
		"givenName":  lc.CredentialSubject.Mandate.Mandatee.FirstName,
		"gender":     lc.CredentialSubject.Mandate.Mandatee.Gender,

		"contactMedium": contactMedium,

		"partyCharacteristic": []any{
			map[string]any{
				"name":  "country",
				"value": user.Country,
			},
		},

		"relatedParty": relatedParty,
	}

	theIdentification := map[string]any{
		"@type":              "individualIdentification",
		"identificationId":   lc.CredentialSubject.Mandate.Mandatee.Id,
		"identificationType": "learcredentialemployee",
		"issuingAuthority":   lc.CredentialSubject.Mandate.Mandator.OrganizationIdentifier,
	}

	indMap["individualIdentification"] = []any{theIdentification}

	content, err := json.Marshal(indMap)
	if err != nil {
		return nil, errl.Errorf("error marshalling organization: %w", err)
	}

	org := &TMFObject{
		ID:         id,
		Type:       objectType,
		Version:    version,
		APIVersion: "v4",
		LastUpdate: lastUpdate,
		Content:    content,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	return org, nil
}
