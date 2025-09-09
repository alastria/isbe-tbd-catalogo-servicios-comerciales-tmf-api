package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"

	"github.com/hesusruiz/isbetmf/internal/errl"
	pdp "github.com/hesusruiz/isbetmf/pdp"
)

func takeDecision(
	ruleEngine *pdp.PDP,
	r *Request,
	tokenClaims map[string]any,
	incomingObjectMap map[string]any,
	existingObjectMap map[string]any,
) (err error) {

	// Some rules are hardcoded because they are always enforced
	// The rest is delegated to the policy engine

	// The object must have both the seller and sellerOperator identities
	sellerDid, sellerOperatorDid, err := getSellerAndBuyerInfo(incomingObjectMap, r.APIVersion)
	if err != nil || sellerDid == "" || sellerOperatorDid == "" {
		err = errl.Errorf("failed to get seller and buyer info: %w", err)
		return err
	}

	userDid := r.AuthUser.OrganizationIdentifier
	if !strings.HasPrefix(userDid, "did:elsi:") {
		userDid = "did:elsi:" + userDid
	}

	// The user is the 'owner' of the object if it is the seller or seller operator
	r.AuthUser.isOwner = (userDid == sellerDid) || (userDid == sellerOperatorDid)

	userArgument := pdp.StarTMFMap(r.AuthUser.ToMap())
	incomingObjectArgument := pdp.StarTMFMap(incomingObjectMap)
	requestArgument := pdp.StarTMFMap(r.ToMap())
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
