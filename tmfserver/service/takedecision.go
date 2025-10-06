package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"

	"github.com/hesusruiz/isbetmf/internal/errl"
	pdp "github.com/hesusruiz/isbetmf/pdp"
	repo "github.com/hesusruiz/isbetmf/tmfserver/repository"
)

func takeDecision(
	ruleEngine *pdp.PDP,
	req *Request,
	tokenClaims map[string]any,
	objectMap repo.TMFObjectMap,
) (err error) {

	var buyerDid, buyerOperatorDid string
	var sellerDid, sellerOperatorDid string
	var userDid string

	// Pre-calculate some useful values
	if req.AuthUser.isAuthenticated {

		// The user DID must be in the format did:elsi:xxxx
		userDid = req.AuthUser.OrganizationIdentifier
		if !strings.HasPrefix(userDid, "did:elsi:") {
			userDid = "did:elsi:" + userDid
		}

		resource := strings.ToLower(req.ResourceName)
		if resource != "organization" && resource != "individual" && resource != "category" {
			// The object must have both the seller and sellerOperator identities
			sellerDid, sellerOperatorDid, err = objectMap.GetSellerInfo(req.APIVersion)
			if err != nil || sellerDid == "" || sellerOperatorDid == "" {
				err = errl.Errorf("failed to get seller and buyer info: %w", err)
				return err
			}

			// Optionally, the object may have Buyer and BuyerOperator roles defined
			buyerDid, buyerOperatorDid, _ = objectMap.GetBuyerInfo(req.APIVersion)

			// The user is the 'owner' if it is Seller, SellerOperator, Buyer or BuyerOperator
			req.AuthUser.isOwner = (userDid == sellerDid) || (userDid == sellerOperatorDid) || (userDid == buyerDid) || (userDid == buyerOperatorDid)

		}

	}

	// Read operations (GET) to public resources are allowed to all users, even unauthenticated ones.
	if req.Method == "GET" && isPublicResource(req.ResourceName) {
		// Check the user-defined rules in the PDP engine
		return callPDP(ruleEngine, req, tokenClaims, objectMap)
	}

	// GET operations to non-public resources, or any other method (POST, PUT, DELETE) to any object
	// are only allowed to authenticated users, and the user must be the owner of the object.

	if !req.AuthUser.isAuthenticated {
		return errl.Errorf("user not authenticated")
	}

	if !req.AuthUser.isOwner {
		return errl.Errorf("user not authorized: not seller, sellerOperator, buyer or buyerOperator")
	}

	err = callPDP(ruleEngine, req, tokenClaims, objectMap)
	return err

}

func callPDP(
	ruleEngine *pdp.PDP,
	req *Request,
	tokenClaims map[string]any,
	objectMap repo.TMFObjectMap,
) (err error) {

	userArgument := pdp.StarTMFMap(req.AuthUser.ToMap())
	incomingObjectArgument := pdp.StarTMFMap(objectMap)
	requestArgument := pdp.StarTMFMap(req.ToMap())
	tokenArgument := pdp.StarTMFMap(tokenClaims)

	// Assemble all data in a single "input" argument, to the style of OPA.
	// We mutate the predeclared identifier, so the policy can access the data for this request.
	// We can also service possible callbacks from the rules engine.

	input := map[string]any{
		"request": requestArgument,
		"token":   tokenArgument,
		"tmf":     incomingObjectArgument,
		"user":    userArgument,
	}

	if slog.Default().Enabled(context.Background(), slog.LevelDebug) {
		b, err := json.MarshalIndent(input, "", "  ")
		if err == nil {
			fmt.Println("PDP input:", string(b))
		}
	}

	decision := true
	if ruleEngine != nil {
		decision, err = ruleEngine.Authorize(input)

		// An error is considered a rejection, continue with the next candidate object
		if err != nil {
			return errl.Errorf("rules engine rejected request due to an error: %w", err)
		}
	}

	// The rules engine rejected the request, continue with the next candidate object
	if !decision {
		return errl.Errorf("PDP: request rejected due to policy")
	}

	// The rules engine accepted the request, add the object to the final list
	slog.Info("PDP: request authorised")
	return nil
}

// Public TMF resources are accessible by all users, even unauthenticated ones
// The public objects are the following:
var publicResources = map[string]bool{
	// TMF620 Product Catalog Management
	"category":             true,
	"catalog":              true,
	"productOffering":      true,
	"productOfferingPrice": true,
	"productSpecification": true,

	// TMF633 Service Catalog Management
	"serviceCatalog":       true,
	"serviceCategory":      true,
	"serviceCandidate":     true,
	"serviceSpecification": true,

	// TMF634 Resource Catalog Management
	"resourceCatalog":       true,
	"resourceCategory":      true,
	"resourceCandidate":     true,
	"resourceSpecification": true,

	// Organization from TMF632 Party Management. But Individual is private.
	"organization": true,

	// TMF669 Party Rola Management
	"partyRole": true,
}

func isPublicResource(resourceName string) bool {
	resourceName = strings.ToLower(resourceName)
	_, ok := publicResources[resourceName]
	return ok
}
