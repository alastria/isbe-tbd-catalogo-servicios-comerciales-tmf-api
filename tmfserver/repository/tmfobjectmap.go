package repository

import (
	"encoding/json"
	"fmt"
	"maps"
	"strings"
	"time"

	"github.com/hesusruiz/isbetmf/config"
	"github.com/hesusruiz/isbetmf/internal/errl"
	"github.com/hesusruiz/isbetmf/internal/jpath"
	"golang.org/x/exp/slog"
)

// TMFObjectMap represents a TMForum object as a map with utility methods
// This is designed to make fixing validation errors simple and efficient
type TMFObjectMap map[string]any

// NewTMFObject creates a new TMFObject from a JSON byte slice
func NewTMFObjectMap(data []byte) (TMFObjectMap, error) {
	var obj TMFObjectMap
	err := json.Unmarshal(data, &obj)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal TMF object: %w", err)
	}
	return obj, nil
}

// NewTMFObjectFromMap creates a new TMFObject from an existing map
func NewTMFObjectFromMap(data map[string]any) TMFObjectMap {
	obj := TMFObjectMap(maps.Clone(data))
	return obj
}

// ToJSON converts the TMFObject to JSON bytes
func (obj TMFObjectMap) ToJSON() ([]byte, error) {
	return json.Marshal(obj)
}

// ToMap converts the TMFObject to a regular map[string]any
func (obj TMFObjectMap) ToMap() map[string]any {
	return maps.Clone(obj)
}

// Utility methods for well-known top-level attributes

// GetID returns the object ID
func (obj TMFObjectMap) GetID() string {
	if id, ok := obj["id"].(string); ok {
		return id
	}
	return ""
}

// SetID sets the object ID
func (obj TMFObjectMap) SetID(id string) {
	obj["id"] = id
}

// GetHref returns the object href
func (obj TMFObjectMap) GetHref() string {
	if href, ok := obj["href"].(string); ok {
		return href
	}
	return ""
}

// SetHref sets the object href
func (obj TMFObjectMap) SetHref(href string) {
	obj["href"] = href
}

// GetVersion returns the object version
func (obj TMFObjectMap) GetVersion() string {
	if version, ok := obj["version"].(string); ok {
		return version
	}
	return ""
}

// SetVersion sets the object version
func (obj TMFObjectMap) SetVersion(version string) {
	obj["version"] = version
}

// GetLastUpdate returns the object lastUpdate
func (obj TMFObjectMap) GetLastUpdate() string {
	if lastUpdate, ok := obj["lastUpdate"].(string); ok {
		return lastUpdate
	}
	return ""
}

// SetLastUpdate sets the object lastUpdate
func (obj TMFObjectMap) SetLastUpdate(lastUpdate string) {
	obj["lastUpdate"] = lastUpdate
}

// SetLastUpdateNow sets the object lastUpdate to current timestamp in RFC3339 format
func (obj TMFObjectMap) SetLastUpdateNow() {
	obj["lastUpdate"] = time.Now().Format(time.RFC3339)
}

// GetType returns the object @type
func (obj TMFObjectMap) GetType() string {
	if objType, ok := obj["@type"].(string); ok {
		return objType
	}
	return ""
}

// SetType sets the object @type
func (obj TMFObjectMap) SetType(objType string) {
	obj["@type"] = objType
}

// GetLifecycleStatus returns the object lifecycleStatus
func (obj TMFObjectMap) GetLifecycleStatus() string {
	if lifecycleStatus, ok := obj["lifecycleStatus"].(string); ok {
		return lifecycleStatus
	}
	return ""
}

// SetLifecycleStatus sets the object lifecycleStatus
func (obj TMFObjectMap) SetLifecycleStatus(lifecycleStatus string) {
	obj["lifecycleStatus"] = lifecycleStatus
}

// GetName returns the object name
func (obj TMFObjectMap) GetName() string {
	if name, ok := obj["name"].(string); ok {
		return name
	}
	return ""
}

// SetName sets the object name
func (obj TMFObjectMap) SetName(name string) {
	obj["name"] = name
}

// GetDescription returns the object description
func (obj TMFObjectMap) GetDescription() string {
	if description, ok := obj["description"].(string); ok {
		return description
	}
	return ""
}

// SetDescription sets the object description
func (obj TMFObjectMap) SetDescription(description string) {
	obj["description"] = description
}

// RelatedParty methods

// GetRelatedParty returns the relatedParty array
func (obj TMFObjectMap) GetRelatedParty() []map[string]any {
	if relatedParty, ok := obj["relatedParty"].([]any); ok {
		result := make([]map[string]any, 0, len(relatedParty))
		for _, item := range relatedParty {
			if rp, ok := item.(map[string]any); ok {
				result = append(result, rp)
			}
		}
		return result
	}
	return nil
}

// SetRelatedParty sets the relatedParty array
func (obj TMFObjectMap) SetRelatedParty(relatedParty []map[string]any) {
	// Convert []map[string]any to []any for JSON serialization
	items := make([]any, len(relatedParty))
	for i, rp := range relatedParty {
		items[i] = rp
	}
	obj["relatedParty"] = items
}

// AddRelatedParty adds a related party entry
func (obj TMFObjectMap) AddRelatedParty(relatedParty map[string]any) {
	current := obj.GetRelatedParty()
	if current == nil {
		current = make([]map[string]any, 0)
	}
	current = append(current, relatedParty)
	obj.SetRelatedParty(current)
}

// HasRelatedParty returns true if the object has related party information
func (obj TMFObjectMap) HasRelatedParty() bool {
	relatedParty := obj.GetRelatedParty()
	return len(relatedParty) > 0
}

// Validation helper methods

// HasField checks if a specific field exists and is not empty
func (obj TMFObjectMap) HasField(field string) bool {
	value, exists := obj[field]
	if !exists {
		return false
	}

	switch v := value.(type) {
	case string:
		return v != ""
	case []any:
		return len(v) > 0
	case map[string]any:
		return len(v) > 0
	default:
		return true // Other types are considered present if they exist
	}
}

// SetField sets a field value
func (obj TMFObjectMap) SetField(field string, value any) {
	obj[field] = value
}

// GetField returns a field value
func (obj TMFObjectMap) GetField(field string) any {
	return obj[field]
}

// GetStringField returns a field value as string
func (obj TMFObjectMap) GetStringField(field string) string {
	if value, ok := obj[field].(string); ok {
		return value
	}
	return ""
}

// SetStringField sets a field value as string
func (obj TMFObjectMap) SetStringField(field string, value string) {
	obj[field] = value
}

// GetArrayField returns a field value as []any
func (obj TMFObjectMap) GetArrayField(field string) []any {
	if value, ok := obj[field].([]any); ok {
		return value
	}
	return nil
}

// SetArrayField sets a field value as []any
func (obj TMFObjectMap) SetArrayField(field string, value []any) {
	obj[field] = value
}

// GetMapField returns a field value as map[string]any
func (obj TMFObjectMap) GetMapField(field string) map[string]any {
	if value, ok := obj[field].(map[string]any); ok {
		return value
	}
	return nil
}

// SetMapField sets a field value as map[string]any
func (obj TMFObjectMap) SetMapField(field string, value map[string]any) {
	obj[field] = value
}

// Clone creates a deep copy of the TMFObject
func (obj TMFObjectMap) Clone() TMFObjectMap {
	// Use JSON marshaling/unmarshaling for deep copy
	data, err := obj.ToJSON()
	if err != nil {
		// If JSON fails, create a shallow copy
		return NewTMFObjectFromMap(obj.ToMap())
	}

	clone, err := NewTMFObjectMap(data)
	if err != nil {
		// If unmarshaling fails, create a shallow copy
		return NewTMFObjectFromMap(obj.ToMap())
	}

	return clone
}

// IsEmpty returns true if the object is empty
func (obj TMFObjectMap) IsEmpty() bool {
	return len(obj) == 0
}

// Keys returns all the keys in the object
func (obj TMFObjectMap) Keys() []string {
	keys := make([]string, 0, len(obj))
	for k := range obj {
		keys = append(keys, k)
	}
	return keys
}

// DeleteField removes a field from the object
func (obj TMFObjectMap) DeleteField(field string) {
	delete(obj, field)
}

// String returns a string representation of the object (for debugging)
func (obj TMFObjectMap) String() string {
	data, err := obj.ToJSON()
	if err != nil {
		return fmt.Sprintf("TMFObject{error: %v}", err)
	}
	return string(data)
}

// setSellerAndBuyerInfo adds the required fields to the incoming object argument
// It calls the appropriate version-specific function based on the API version
func (tmfObjectMap TMFObjectMap) SetSellerAndBuyerInfo(organizationIdentifier string, apiVersion string) (err error) {
	switch apiVersion {
	case "v4":
		return setSellerAndBuyerInfoV4(tmfObjectMap, organizationIdentifier)
	case "v5":
		return setSellerAndBuyerInfoV5(tmfObjectMap, organizationIdentifier)
	default:
		// Default to V5 for backward compatibility
		return setSellerAndBuyerInfoV4(tmfObjectMap, organizationIdentifier)
	}
}

// setSellerAndBuyerInfoV4 adds the required fields to the incoming object argument for V4 API
// Specifically, the Seller and SellerOperator roles are added to the relatedParty list
func setSellerAndBuyerInfoV4(tmfObjectMap map[string]any, organizationIdentifier string) (err error) {

	// Normalize the organization identifier to the DID format
	if !strings.HasPrefix(organizationIdentifier, "did:elsi:") {
		organizationIdentifier = "did:elsi:" + organizationIdentifier
	}

	// Look for the "Seller", "SellerOperator", "Buyer" and "BuyerOperator" roles
	relatedParties := jpath.GetList(tmfObjectMap, "relatedParty")

	// Build the two entries for V4 format
	sellerEntry := map[string]any{
		"role":          "Seller",
		"id":            "urn:ngsi-ld:organization:" + organizationIdentifier,
		"href":          "urn:ngsi-ld:organization:" + organizationIdentifier,
		"name":          organizationIdentifier,
		"@referredType": "Organization",
	}
	sellerOperator := map[string]any{
		"role":          "SellerOperator",
		"id":            "urn:ngsi-ld:organization:" + config.ServerOperatorDid,
		"href":          "urn:ngsi-ld:organization:" + config.ServerOperatorDid,
		"name":          config.ServerOperatorDid,
		"@referredType": "Organization",
	}

	if len(relatedParties) == 0 {
		slog.Debug("setSellerAndBuyerInfoV4: no relatedParty, adding seller and sellerOperator")
		tmfObjectMap["relatedParty"] = []any{sellerEntry, sellerOperator}
		return nil
	}

	foundSeller := false
	foundSellerOperator := false

	newRelatedParties := []any{}

	for _, rp := range relatedParties {

		// Convert entry to a map
		rpMap, _ := rp.(map[string]any)
		if len(rpMap) == 0 {
			return errl.Errorf("invalid relatedParty entry")
		}

		rpRole, _ := rpMap["role"].(string)
		rpRole = strings.ToLower(rpRole)

		if rpRole != "seller" && rpRole != "selleroperator" {
			newRelatedParties = append(newRelatedParties, rp)
			// Go to next entry
			continue
		}

		if rpRole == "seller" {
			// Overwrite the entry, because we can not allow the user to create fake info
			newRelatedParties = append(newRelatedParties, sellerEntry)
			foundSeller = true
			continue
		}
		if rpRole == "selleroperator" {
			// Overwrite the entry, because we can not allow the user to create fake info
			newRelatedParties = append(newRelatedParties, sellerOperator)
			foundSellerOperator = true
			continue
		}

	}

	if !foundSeller {
		// Add the seller if it is not already in the list
		slog.Debug("setSellerAndBuyerInfoV4: adding seller", "organizationIdentifier", organizationIdentifier)
		newRelatedParties = append(newRelatedParties, sellerEntry)
	}

	if !foundSellerOperator {
		// Add the seller operator if it is not already in the list
		slog.Debug("setSellerAndBuyerInfoV4: adding seller operator", "organizationIdentifier", organizationIdentifier)
		newRelatedParties = append(newRelatedParties, sellerOperator)
	}

	tmfObjectMap["relatedParty"] = newRelatedParties

	return nil

}

// setSellerAndBuyerInfoV5 adds the required fields to the incoming object argument
// Specifically, the Seller and SellerOperator roles are added to the relatedParty list
func setSellerAndBuyerInfoV5(tmfObjectMap map[string]any, organizationIdentifier string) (err error) {

	// Normalize all organization identifiers to the DID format
	if !strings.HasPrefix(organizationIdentifier, "did:elsi:") {
		organizationIdentifier = "did:elsi:" + organizationIdentifier
	}

	// Look for the "Seller", "SellerOperator", "Buyer" and "BuyerOperator" roles
	relatedParties := jpath.GetList(tmfObjectMap, "relatedParty")

	// Build the two entries
	sellerEntry := map[string]any{
		"role":  "Seller",
		"@type": "RelatedPartyRefOrPartyRoleRef",
		"partyOrPartyRole": map[string]any{
			"@type":         "PartyRef",
			"href":          "urn:ngsi-ld:organization:" + organizationIdentifier,
			"id":            "urn:ngsi-ld:organization:" + organizationIdentifier,
			"name":          organizationIdentifier,
			"@referredType": "Organization",
		},
	}
	sellerOperator := map[string]any{
		"role":  "SellerOperator",
		"@type": "RelatedPartyRefOrPartyRoleRef",
		"partyOrPartyRole": map[string]any{
			"@type":         "PartyRef",
			"href":          "urn:ngsi-ld:organization:" + config.ServerOperatorDid,
			"id":            "urn:ngsi-ld:organization:" + config.ServerOperatorDid,
			"name":          config.ServerOperatorDid,
			"@referredType": "Organization",
		},
	}

	if len(relatedParties) == 0 {
		slog.Debug("setSellerAndBuyerInfo: no relatedParty, adding seller and sellerOperator")
		tmfObjectMap["relatedParty"] = []any{sellerEntry, sellerOperator}
		return nil
	}

	foundSeller := false
	foundSellerOperator := false

	newRelatedParties := []any{}

	for _, rp := range relatedParties {

		// Convert entry to a map
		rpMap, _ := rp.(map[string]any)
		if len(rpMap) == 0 {
			return errl.Errorf("invalid relatedParty entry")
		}

		rpRole, _ := rpMap["role"].(string)
		rpRole = strings.ToLower(rpRole)

		if rpRole != "seller" && rpRole != "selleroperator" {
			newRelatedParties = append(newRelatedParties, rp)
			// Go to next entry
			continue
		}

		if rpRole == "seller" {
			// Overwrite the entry, because we can not allow the user to create fake info
			newRelatedParties = append(newRelatedParties, sellerEntry)
			foundSeller = true
			continue
		}
		if rpRole == "selleroperator" {
			// Overwrite the entry, because we can not allow the user to create fake info
			newRelatedParties = append(newRelatedParties, sellerOperator)
			foundSellerOperator = true
			continue
		}

	}

	if !foundSeller {
		// Add the seller if it is not already in the list
		slog.Debug("setSellerAndBuyerInfo: adding seller", "organizationIdentifier", organizationIdentifier)
		newRelatedParties = append(newRelatedParties, sellerEntry)
	}

	if !foundSellerOperator {
		// Add the seller operator if it is not already in the list
		slog.Debug("setSellerAndBuyerInfo: adding seller operator", "organizationIdentifier", organizationIdentifier)
		newRelatedParties = append(newRelatedParties, sellerOperator)
	}

	tmfObjectMap["relatedParty"] = newRelatedParties

	return nil

}

func (tmfObjectMap TMFObjectMap) GetSellerAndBuyerInfo(apiVersion string) (sellerDid string, sellerOperatorDid string, err error) {
	switch apiVersion {
	case "v4":
		return getSellerAndBuyerInfoV4(tmfObjectMap)
	case "v5":
		return getSellerAndBuyerInfoV5(tmfObjectMap)
	default:
		// Default to V5 for backward compatibility
		return getSellerAndBuyerInfoV4(tmfObjectMap)
	}
}

func getSellerAndBuyerInfoV4(tmfObjectMap map[string]any) (sellerDid string, sellerOperatorDid string, err error) {
	// In V4, relatedParty is a list of maps with fields like "role", "id", "href", "name", "@referredType"
	// We need to extract the "name" for the "Seller" and "SellerOperator" roles

	relatedParties := jpath.GetList(tmfObjectMap, "relatedParty")

	if len(relatedParties) == 0 {
		err = errl.Errorf("no relatedParty")
		return
	}

	for _, rp := range relatedParties {
		rpMap, _ := rp.(map[string]any)
		if len(rpMap) == 0 {
			return "", "", errl.Errorf("invalid relatedParty entry")
		}

		rpRole, _ := rpMap["role"].(string)
		rpRole = strings.ToLower(rpRole)

		if rpRole != "seller" && rpRole != "selleroperator" {
			continue
		}

		if rpRole == "seller" {
			sellerDid, _ = rpMap["name"].(string)
			continue
		}
		if rpRole == "selleroperator" {
			sellerOperatorDid, _ = rpMap["name"].(string)
			continue
		}
	}

	if sellerDid == "" && sellerOperatorDid == "" {
		err = errl.Errorf("no seller or seller operator")
		return
	}

	if sellerDid == "" {
		err = errl.Errorf("no seller")
		return
	}
	if sellerOperatorDid == "" {
		err = errl.Errorf("no seller operator")
		return
	}

	return

}

func getSellerAndBuyerInfoV5(tmfObjectMap map[string]any) (sellerDid string, sellerOperatorDid string, err error) {

	// Look for the "Seller", "SellerOperator", "Buyer" and "BuyerOperator" roles
	relatedParties := jpath.GetList(tmfObjectMap, "relatedParty")

	if len(relatedParties) == 0 {
		err = errl.Errorf("no relatedParty")
		return
	}

	for _, rp := range relatedParties {

		// Convert entry to a map
		rpMap, _ := rp.(map[string]any)
		if len(rpMap) == 0 {
			return "", "", errl.Errorf("invalid relatedParty entry")
		}

		rpRole, _ := rpMap["role"].(string)
		rpRole = strings.ToLower(rpRole)

		if rpRole != "seller" && rpRole != "selleroperator" {
			// Go to next entry
			continue
		}

		if rpRole == "seller" {
			party, _ := rpMap["partyOrPartyRole"].(map[string]any)
			sellerDid, _ = party["name"].(string)
			continue
		}
		if rpRole == "selleroperator" {
			party, _ := rpMap["partyOrPartyRole"].(map[string]any)
			sellerOperatorDid, _ = party["name"].(string)
			continue
		}

	}

	if sellerDid == "" && sellerOperatorDid == "" {
		err = errl.Errorf("no seller or seller operator")
		return
	}

	if sellerDid == "" {
		err = errl.Errorf("no seller")
		return
	}
	if sellerOperatorDid == "" {
		err = errl.Errorf("no seller operator")
		return
	}

	return

}
