package types

import (
	"encoding/json"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hesusruiz/isbetmf/internal/errl"
)

type Mandate struct {
	Id       string `json:"id,omitempty"`
	Mandator struct {
		OrganizationIdentifier string `json:"organizationIdentifier,omitempty"` // OID 2.5.4.97
		CommonName             string `json:"commonName,omitempty"`             // OID 2.5.4.3
		GivenName              string `json:"givenName,omitempty"`
		Surname                string `json:"surname,omitempty"`
		EmailAddress           string `json:"emailAddress,omitempty"`
		SerialNumber           string `json:"serialNumber,omitempty"`
		Organization           string `json:"organization,omitempty"`
		Country                string `json:"country,omitempty"`
	} `json:"mandator"`
	Mandatee struct {
		Id           string `json:"id,omitempty"`
		FirstName    string `json:"firstName,omitempty"`
		LastName     string `json:"lastName,omitempty"`
		Gender       string `json:"gender,omitempty"`
		Email        string `json:"email,omitempty"`
		Mobile_phone string `json:"mobile_phone,omitempty"`
	} `json:"mandatee"`
	Power []struct {
		Id           string   `json:"id,omitempty"`
		Tmf_type     string   `json:"type,omitempty"`
		Tmf_domain   string   `json:"domain,omitempty"`
		Tmf_function string   `json:"function,omitempty"`
		Tmf_action   []string `json:"action,omitempty"`
	} `json:"power"`
}

type CredentialIssuer struct {
	Id                     string `json:"id,omitempty"`
	OrganizationIdentifier string `json:"organizationIdentifier,omitempty"`
	Organization           string `json:"organization,omitempty"`
	Country                string `json:"country,omitempty"`
	SerialNumber           string `json:"serialNumber,omitempty"`
	CommonName             string `json:"commonName,omitempty"`
}

type OnePower struct {
	Id           string   `json:"id,omitempty"`
	Tmf_type     string   `json:"type,omitempty"`
	Tmf_domain   string   `json:"domain,omitempty"`
	Tmf_function string   `json:"function,omitempty"`
	Tmf_action   []string `json:"action,omitempty"`
}

// Example of credentialStatus
//
//	"credentialStatus": {
//		"id": "https://issuer.dome-marketplace-sbx.org/backoffice/v1/credentials/status/1#SYC908RIQQqeUXR19nh3EQ",
//		"statusListCredential": "https://issuer.dome-marketplace-sbx.org/backoffice/v1/credentials/status/1",
//		"statusListIndex": "SYC908RIQQqeUXR19nh3EQ",
//		"statusPurpose": "revocation",
//		"type": "PlainListEntity"
//	},
type CredentialStatus struct {
	Id                   string `json:"id,omitempty"`
	StatusListCredential string `json:"statusListCredential,omitempty"`
	StatusListIndex      string `json:"statusListIndex,omitempty"`
	StatusPurpose        string `json:"statusPurpose,omitempty"`
	Type                 string `json:"type,omitempty"`
}

type LEARCredentialEmployee struct {
	Context []string `json:"@context,omitempty"`
	// CredentialStatus  CredentialStatus `json:"credentialStatus,omitempty"`
	Id             string   `json:"id,omitempty"`
	TypeCredential []string `json:"type,omitempty"`
	// Issuer            CredentialIssuer `json:"issuer,omitempty"`
	ValidFrom         string `json:"validFrom,omitempty"`
	ValidUntil        string `json:"validUntil,omitempty"`
	CredentialSubject struct {
		Mandate Mandate `json:"mandate,omitempty"`
	} `json:"credentialSubject"`
}

type LEARCredentialEmployeeJWTClaims struct {
	LEARCredentialEmployee
	jwt.RegisteredClaims
}

// // CreateLEARCredentialJWTtoken creates a JWT token from the given claims,
// // signed with the first private key associated to the issuer DID
// func CreateLEARCredentialJWTtoken(learCred LEARCredentialEmployee, sigMethod jwt.SigningMethod, privateKey any) (string, error) {

// 	// Prepare some fields of the LEARCredential
// 	now := time.Now()

// 	// Create claims with multiple fields populated
// 	claims := LEARCredentialEmployeeJWTClaims{
// 		learCred,
// 		jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(now.Add(24 * 365 * time.Hour)),
// 			IssuedAt:  jwt.NewNumericDate(now),
// 			NotBefore: jwt.NewNumericDate(now),
// 			Issuer:    learCred.Issuer.Id,
// 			Subject:   learCred.CredentialSubject.Mandate.Mandatee.Id,
// 			ID:        learCred.Id,
// 			Audience:  []string{"everybody"},
// 		},
// 	}

// 	// Serialize and sign the JWT. The result is a byte array with the JWT in compact form:
// 	// header.payload.signature
// 	token := jwt.NewWithClaims(sigMethod, claims)
// 	ss, err := token.SignedString(privateKey)
// 	fmt.Println(ss, err)

// 	return ss, nil

// }

func LEARCredentialFromMap(s map[string]any) (*LEARCredentialEmployee, error) {

	// Marshall to a byte array
	out, err := json.Marshal(s)
	if err != nil {
		return nil, errl.Errorf("error marshalling credential: %w", err)
	}

	// Unmarshal to a struct
	lc := &LEARCredentialEmployee{}
	err = json.Unmarshal(out, &lc)
	if err != nil {
		return nil, errl.Errorf("error unmarshalling credential: %w", err)
	}

	return lc, nil

}
