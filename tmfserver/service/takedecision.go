package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"

	"github.com/hesusruiz/isbetmf/internal/errl"
	"github.com/hesusruiz/isbetmf/internal/jpath"
	pdp "github.com/hesusruiz/isbetmf/pdp"
	repo "github.com/hesusruiz/isbetmf/tmfserver/repository"
)

func (svc *Service) takeDecision(
	ruleEngine *pdp.PDP,
	req *Request,
	tokenClaims map[string]any,
	objectMap repo.TMFObjectMap,
) (authorized bool, err error) {

	// Evaluate the hardcoded policies, if they fail return immediately.
	// Otherwise, continue to see if the user policies allow access
	decision, reason := svc.hardcodedPolicy(req, objectMap)
	if !decision {
		return false, reason
	}

	// The caller is the owner, at least according to hardcoded policies.
	// The user policies will determine the final decision.
	req.AuthUser.IsOwner = decision

	if err := svc.callPDP(ruleEngine, req, tokenClaims, objectMap); err != nil {
		return false, errl.Errorf("user policies in PDP engine: %w", err)
	}

	return true, nil

}

func (svc *Service) hardcodedPolicy(
	req *Request,
	obj repo.TMFObjectMap,
) (decision bool, reason error) {

	// Read operations (GET) to public resources are allowed to all users, even unauthenticated ones.
	// Method GET includes actions READ and LIST
	if req.Method == "GET" && isPublicResource(req.ResourceName) {
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
		if sameOrganizations(caller.OrganizationIdentifier, svc.ServerOperatorDid) {
			return true, errl.Errorf("caller %s is server operator %s", caller.OrganizationIdentifier, svc.ServerOperatorDid)
		}

		// If the organization of the caller and object are the same, then the caller can read/write/update/delete
		objectOrganizationId := jpath.GetString(obj, "organizationIdentification.*.identificationId")
		if sameOrganizations(objectOrganizationId, caller.OrganizationIdentifier) {
			return true, errl.Errorf("caller %s is same as in object %s", caller.OrganizationIdentifier, objectOrganizationId)
		}

		return false, errl.Errorf("caller (%s) is neither the same as in object (%s) or the server operator", caller.OrganizationIdentifier, objectOrganizationId)

	case "individual":

		// TODO: revise this policy to be restrictive

		// If the caller is us (the server operator), then we can read/write/update/delete
		if sameOrganizations(caller.OrganizationIdentifier, svc.ServerOperatorDid) {
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
				if sameOrganizations(issuingAuthority, caller.OrganizationIdentifier) {
					return true, errl.Errorf("caller %s is same as mandator in Individual object %s", caller.OrganizationIdentifier, issuingAuthority)
				} else {
					return false, errl.Errorf("caller (%s) is neither the mandator in Individual object (%s) or the server operator", caller.OrganizationIdentifier, issuingAuthority)
				}
			}
		}

		return false, errl.Errorf("caller (%s) is neither the mandator in Individual object or the server operator", caller.OrganizationIdentifier)

	case "category":

		// If the caller is us (the server operator), then we can read/write/update/delete
		if sameOrganizations(caller.OrganizationIdentifier, svc.ServerOperatorDid) {
			return true, errl.Errorf("caller %s is server operator %s", caller.OrganizationIdentifier, svc.ServerOperatorDid)
		}

		return false, errl.Errorf("caller %s is not the server operator %s", caller.OrganizationIdentifier, svc.ServerOperatorDid)

	default:

		// If the caller is us (the server operator), then we can read/write/update/delete
		if sameOrganizations(caller.OrganizationIdentifier, svc.ServerOperatorDid) {
			return true, errl.Errorf("caller %s is server operator %s", caller.OrganizationIdentifier, svc.ServerOperatorDid)
		}

		// For any other objects, we require that the object includes the Seller info, and then
		// the user must be either the server operator or the seller

		// Try to retrieve the Seller info
		objSellerDid, objSellerOperatorDid, err := obj.GetSellerInfo("v4")
		if err != nil {
			return false, errl.Errorf("object (%s) does not contain seller information", obj.GetID())
		}

		// If the caller is the same as the object SellerOperator or the Seller, then is the owner
		if sameOrganizations(caller.OrganizationIdentifier, objSellerDid) || sameOrganizations(caller.OrganizationIdentifier, objSellerOperatorDid) {
			return true, errl.Errorf("caller %s is seller %s or seller operator %s", caller.OrganizationIdentifier, objSellerDid, objSellerOperatorDid)
		}

		// Try to retrieve the Buyer info, which may not exist.
		// We already checked for Seller info, so if Buyer info does not exist, caller is not the owner
		objBuyerDid, objBuyerOperatorDid, err := obj.GetBuyerInfo("v4")
		if err != nil {
			return false, errl.Errorf("object (%s) does not contain buyer information and caller (%s) is not the seller or seller operator", obj.GetID(), caller.OrganizationIdentifier)
		}

		// If the caller is the same as the object BuyerOperator or the Buyer, then is the owner
		if sameOrganizations(caller.OrganizationIdentifier, objBuyerDid) || sameOrganizations(caller.OrganizationIdentifier, objBuyerOperatorDid) {
			return true, errl.Errorf("caller %s is buyer %s or buyer operator %s", caller.OrganizationIdentifier, objBuyerDid, objBuyerOperatorDid)
		}

		return false, errl.Errorf("caller %s is not seller or buyer in object %s", caller.OrganizationIdentifier, obj.GetID())

	}

}

func (svc *Service) callPDP(
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
