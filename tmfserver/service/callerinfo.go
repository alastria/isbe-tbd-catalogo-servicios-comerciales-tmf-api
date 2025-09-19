package service

import (
	"errors"
	"log/slog"
	"strings"

	"github.com/hesusruiz/isbetmf/internal/errl"
	"github.com/hesusruiz/isbetmf/internal/jpath"
	"github.com/hesusruiz/isbetmf/tmfserver/repository"
	"gitlab.com/greyxor/slogor"
)

const AllowFakeClaims = true
const VerifyTokenSignature = false

// processAccessToken retrieves the Access Token from the request, verifies it if it exists and
// creates a map ready to be passed to the rules engine.
//
// The access token may not exist, but if it does then it must be valid.
// For convenience of the policies, some calculated fields are created and returned in the 'user' object.
func (svc *Service) processAccessToken(r *Request) (tokenClaims map[string]any, err error) {

	var authUser *AuthUser

	// This is to support testing
	verify := true

	if len(r.AccessToken) == 0 {
		// The user did not provide an access token.
		// Normally this is forbidden, but for testing we can provide a fake one, and do not verify signature
		if AllowFakeClaims {
			slog.Debug("PDP: using fake claims for testing")
			r.AccessToken = FakeAT
			verify = false
		}
	}

	// An empty token is not considered an error, and the caller should enforce its existence
	if len(r.AccessToken) == 0 {
		return nil, nil
	}

	// It is an error to send an invaild token with the request, so we have to verify it.

	// Verify the token and extract the claims.
	// A verification error stops processing.

	if !VerifyTokenSignature {
		verify = false
	}

	tokenClaims, authUser, err = ParseJWT(svc, r.AccessToken, verify)
	if err != nil {
		slog.Error("invalid access token", slogor.Err(err), "token", r.AccessToken)
		return nil, errl.Errorf("invalid access token: %w", err)
	}

	verifiableCredential := jpath.GetMap(tokenClaims, "vc")

	if len(verifiableCredential) > 0 {
		authUser.isAuthenticated = true

		powers := jpath.GetList(verifiableCredential, "credentialSubject.mandate.power")
		for _, p := range powers {

			// This is to support old version of the Verifiable Credential
			ptype := jpath.GetString(p, "type")
			pdomain := jpath.GetString(p, "domain")
			pfunction := jpath.GetString(p, "function")
			paction := jpath.GetString(p, "action")

			// Check fields without regards to case
			if strings.EqualFold(ptype, "Domain") &&
				strings.EqualFold(pdomain, "DOME") &&
				strings.EqualFold(pfunction, "Onboarding") &&
				strings.EqualFold(paction, "execute") {

				authUser.isLEAR = true

			}

			// And this for the new version of the Verifiable Credential
			ptype = jpath.GetString(p, "tmf_type")
			pdomain = jpath.GetString(p, "tmf_domain")
			pfunction = jpath.GetString(p, "tmf_function")
			paction = jpath.GetString(p, "tmf_action")

			if strings.EqualFold(ptype, "Domain") &&
				strings.EqualFold(pdomain, "DOME") &&
				strings.EqualFold(pfunction, "Onboarding") &&
				strings.EqualFold(paction, "execute") {

				authUser.isLEAR = true
			}

		}

	} else {

		// There is not a Verifiable Credential inside the token
		return nil, errl.Errorf("access token without verifiable credential: %s", r.AccessToken)

	}

	// Update the request with the authenticated user info
	r.AuthUser = authUser

	// If the user has an organization identifier, create a new organization object.
	// If it is created, we just receive an error which we ignore
	// This is needed to be able to create a new organization object in the DOME Marketplace.
	if len(authUser.OrganizationIdentifier) > 0 {

		org := &repository.Organization{
			CommonName:             authUser.CommonName,
			Country:                authUser.Country,
			EmailAddress:           authUser.EmailAddress,
			Organization:           authUser.Organization,
			OrganizationIdentifier: authUser.OrganizationIdentifier,
			SerialNumber:           authUser.SerialNumber,
		}

		obj, _ := repository.TMFOrganizationFromToken(tokenClaims, org)

		if err := svc.createObject(obj); err != nil {
			if errors.Is(err, &ErrObjectExists{}) {
				slog.Debug("organization already exists", "organizationIdentifier", authUser.OrganizationIdentifier)
			} else {
				err = errl.Error(err)
				return nil, err
			}
		}
	}

	return tokenClaims, nil

}
