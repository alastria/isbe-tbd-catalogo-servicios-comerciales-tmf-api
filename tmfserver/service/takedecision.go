package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"

	"github.com/hesusruiz/isbetmf/config"
	"github.com/hesusruiz/isbetmf/internal/errl"
	"github.com/hesusruiz/isbetmf/internal/jpath"
	pdp "github.com/hesusruiz/isbetmf/pdp"
	repo "github.com/hesusruiz/isbetmf/tmfserver/repository"
)

// takeDecision evaluates access authorization for a TMF request against both hardcoded and user-defined policies.
// It first checks hardcoded policies, and if they pass, proceeds to evaluate user policies in the PDP engine.
//
// Parameters:
//   - ruleEngine: PDP engine instance for evaluating user policies
//   - req: The request details containing auth and operation information
//   - tokenClaims: JWT claims from the authentication token
//   - objectMap: Map of TMF objects involved in the request
//
// Returns:
//   - authorized: boolean indicating if access is authorized
//   - err: error if any occurred during policy evaluation
//
// The function works in two stages:
// 1. Evaluates hardcoded policies first - if these fail, denies access immediately
// 2. If hardcoded policies pass, evaluates user-defined policies through the PDP engine
func (svc *Service) takeDecision(
	ruleEngine *pdp.PDP,
	req *Request,
	tokenClaims map[string]any,
	obj *repo.TMFRecord,
) (authorized bool, err error) {

	objectMap, err := obj.ToTMFObjectMap()
	if err != nil {
		return false, errl.Errorf("failed to convert object to map: %w", err)
	}

	// Evaluate the hardcoded policies, if they fail return immediately.
	// Otherwise, continue to see if the user policies allow access
	decision, reason := svc.hardcodedPolicies(req, objectMap)
	if !decision {
		return false, reason
	}

	// The caller is the owner, at least according to hardcoded policies.
	// The user policies will determine the final decision.
	req.AuthUser.IsOwner = decision

	if err := svc.userPolicies(ruleEngine, req, tokenClaims, objectMap); err != nil {
		return false, errl.Errorf("user policies in PDP engine: %w", err)
	}

	return true, nil

}

// hardcodedPolicies determines if a request is authorized based on predefined access control rules.
// It evaluates the request against different policies depending on the object type and user authentication.
//
// The following rules are applied:
// - GET requests to public resources are allowed for all users (authenticated or not)
// - All other operations require authenticated users
// - For Organization objects:
//   - Server operator has full access
//   - Organization members have full access to their own organization
//
// - For Individual objects:
//   - Server operator has full access
//   - Organizations listed as mandators in LEARCredential have full access
//
// - For Category objects:
//   - Only server operator has full access
//
// - For all other objects:
//   - Server operator has full access
//   - Seller and SellerOperator have full access
//   - Buyer and BuyerOperator have full access if buyer info exists
//
// Parameters:
//   - req: Contains request details including method, resource name and authenticated user info
//   - obj: The TMF object map being accessed
//
// Returns:
//   - decision: true if access is granted, false otherwise
//   - reason: error containing the reason for the decision
func (svc *Service) hardcodedPolicies(req *Request, obj repo.TMFObjectMap) (decision bool, reason error) {

	// Read operations (GET) to public resources are allowed to all users, even unauthenticated ones.
	// Method GET includes actions READ and LIST
	if req.Method == "GET" && config.IsPublicResource(req.ResourceName) {
		return true, errl.Errorf("GET request to a public resource %s", req.ResourceName)
	}

	// GET operations to non-public resources, or any other method (POST, PUT, DELETE) to any object
	// are only allowed to authenticated users.
	if !req.AuthUser.IsAuthenticated {
		return false, errl.Errorf("user not authenticated")
	}

	// Ownership of an object depends on the type of object
	objType, _ := obj["@type"].(string)
	objType = strings.ToLower(objType)

	caller := req.AuthUser

	switch objType {
	case "organization":

		// If the caller is us (the server operator), then we can read/write/update/delete
		if repo.SameOrganizations(caller.OrganizationIdentifier, svc.ServerOperatorDid) {
			return true, errl.Errorf("caller %s is server operator %s", caller.OrganizationIdentifier, svc.ServerOperatorDid)
		}

		// If the organization of the caller and object are the same, then the caller can read/write/update/delete
		objectOrganizationId := jpath.GetString(obj, "organizationIdentification.*.identificationId")
		if repo.SameOrganizations(objectOrganizationId, caller.OrganizationIdentifier) {
			return true, errl.Errorf("caller %s is same as in object %s", caller.OrganizationIdentifier, objectOrganizationId)
		}

		return false, errl.Errorf("caller (%s) is neither the same as in object (%s) or the server operator", caller.OrganizationIdentifier, objectOrganizationId)

	case "individual":

		// TODO: revise this policy to be restrictive

		// If the caller is us (the server operator), then we can read/write/update/delete
		if repo.SameOrganizations(caller.OrganizationIdentifier, svc.ServerOperatorDid) {
			return true, errl.Errorf("caller %s is server operator %s", caller.OrganizationIdentifier, svc.ServerOperatorDid)
		}

		// If the caller is the Organization that is the mandator is the LEARCredential of the employee
		// then the caller can read/write/update/delete
		individualIdentificationArray := jpath.GetList(obj, "individualIdentification")

		// Look for an entry with 'identificationType=learcredentialemployee'
		for _, individualIdentification := range individualIdentificationArray {
			individualIdentificationMap, _ := individualIdentification.(map[string]any)
			if individualIdentificationMap["identificationType"] == "learcredentialemployee" {
				// The 'issuingAuthority' must be equal to the caller organizationIdentifier
				issuingAuthority := individualIdentificationMap["issuingAuthority"].(string)
				if repo.SameOrganizations(issuingAuthority, caller.OrganizationIdentifier) {
					return true, errl.Errorf("caller %s is same as mandator in Individual object %s", caller.OrganizationIdentifier, issuingAuthority)
				} else {
					return false, errl.Errorf("caller (%s) is neither the mandator in Individual object (%s) or the server operator", caller.OrganizationIdentifier, issuingAuthority)
				}
			}
		}

		return false, errl.Errorf("caller (%s) is neither the mandator in Individual object or the server operator", caller.OrganizationIdentifier)

	case "category":

		// If the caller is us (the server operator), then we can read/write/update/delete
		if repo.SameOrganizations(caller.OrganizationIdentifier, svc.ServerOperatorDid) {
			return true, errl.Errorf("caller %s is server operator %s", caller.OrganizationIdentifier, svc.ServerOperatorDid)
		}

		return false, errl.Errorf("caller %s is not the server operator %s", caller.OrganizationIdentifier, svc.ServerOperatorDid)

	default:

		// If the caller is us (the server operator), then we can read/write/update/delete
		if repo.SameOrganizations(caller.OrganizationIdentifier, svc.ServerOperatorDid) {
			return true, errl.Errorf("caller %s is server operator %s", caller.OrganizationIdentifier, svc.ServerOperatorDid)
		}

		// For any other objects, we require that the object includes the Seller info, and then
		// the user must be either the server operator or the seller

		// Try to retrieve the Seller info
		objSellerDid, objSellerOperatorDid, err := obj.GetSellerInfo("v4")
		if err != nil {
			return false, errl.Errorf("object (%s) does not contain seller information", obj.ID())
		}

		// If the caller is the same as the object SellerOperator or the Seller, then is the owner
		if repo.SameOrganizations(caller.OrganizationIdentifier, objSellerDid) || repo.SameOrganizations(caller.OrganizationIdentifier, objSellerOperatorDid) {
			return true, errl.Errorf("caller %s is seller %s or seller operator %s", caller.OrganizationIdentifier, objSellerDid, objSellerOperatorDid)
		}

		// Try to retrieve the Buyer info, which may not exist.
		// We already checked for Seller info, so if Buyer info does not exist, caller is not the owner
		objBuyerDid, objBuyerOperatorDid, err := obj.GetBuyerInfo("v4")
		if err != nil {
			return false, errl.Errorf("object (%s) does not contain buyer information and caller (%s) is not the seller or seller operator", obj.ID(), caller.OrganizationIdentifier)
		}

		// If the caller is the same as the object BuyerOperator or the Buyer, then is the owner
		if repo.SameOrganizations(caller.OrganizationIdentifier, objBuyerDid) || repo.SameOrganizations(caller.OrganizationIdentifier, objBuyerOperatorDid) {
			return true, errl.Errorf("caller %s is buyer %s or buyer operator %s", caller.OrganizationIdentifier, objBuyerDid, objBuyerOperatorDid)
		}

		return false, errl.Errorf("caller %s is not seller or buyer in object %s", caller.OrganizationIdentifier, obj.ID())

	}

}

// userPolicies constructs the policy input for a PDP (policy decision point)
// from the provided request, token claims, TMF object map and authenticated user,
// then evaluates the assembled input against the rules engine to determine
// whether the request is authorized.
//
// The assembled input contains the following top-level keys:
//   - "request": a dereferenced map representation of req (req.ToMap())
//   - "token": a dereferenced map representation of tokenClaims
//   - "tmf": a dereferenced map representation of the incoming TMF object map
//   - "user": a dereferenced map representation of the authenticated user
//
// The function will pretty-print the assembled input as JSON when the default
// slog logger is at debug level to aid debugging.
//
// Decision & error semantics:
//   - If ruleEngine is nil the function defaults to allowing the request (returns nil).
//   - If ruleEngine is present, ruleEngine.Authorize(input) is invoked.
//   - Any error returned by the PDP is treated as a rejection and returned to the caller.
//   - If the PDP returns false the request is considered denied and a non-nil error is returned.
//   - If the PDP returns true the request is authorized and the function returns nil.
//
// Side effects:
//   - Logs an informational message on authorization and prints debug JSON when enabled.
//   - Constructs a single "input" object (OPA-style) so policies can access request,
//     token, tmf and user together and support callbacks.
//
// Return value:
//   - nil when the request is authorized.
//   - non-nil error when the rules engine rejects the request or an internal error occurs.
func (svc *Service) userPolicies(
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
