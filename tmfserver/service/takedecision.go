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

	// Pre-calculate the owner value
	req.AuthUser.IsOwner = objectMap.IsOwner(req.AuthUser)

	// Read operations (GET) to public resources are allowed to all users, even unauthenticated ones.
	if req.Method == "GET" && isPublicResource(req.ResourceName) {
		// Check the user-defined rules in the PDP engine
		return callPDP(ruleEngine, req, tokenClaims, objectMap)
	}

	// GET operations to non-public resources, or any other method (POST, PUT, DELETE) to any object
	// are only allowed to authenticated users, and the user must be the owner of the object.

	if !req.AuthUser.IsAuthenticated {
		return errl.Errorf("user not authenticated")
	}

	if !req.AuthUser.IsOwner {
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
	"productoffering":      true,
	"productofferingprice": true,
	"productspecification": true,

	// TMF633 Service Catalog Management
	"servicecatalog":       true,
	"servicecategory":      true,
	"servicecandidate":     true,
	"servicespecification": true,

	// TMF634 Resource Catalog Management
	"resourcecatalog":       true,
	"resourcecategory":      true,
	"resourcecandidate":     true,
	"resourcespecification": true,

	// Organization from TMF632 Party Management. But Individual is private.
	"organization": true,

	// TMF669 Party Rola Management
	"partyrole": true,
}

func isPublicResource(resourceName string) bool {
	resourceName = strings.ToLower(resourceName)
	_, ok := publicResources[resourceName]
	return ok
}
