package replicate

import (
	"fmt"
	"slices"
	"strings"
	"time"
)

// ValidationResult represents the result of validating a single object
type ValidationResult struct {
	ObjectID      string              `json:"object_id"`
	ObjectType    string              `json:"object_type"`
	Valid         bool                `json:"valid"`
	Errors        []ValidationError   `json:"errors,omitempty"`
	Warnings      []ValidationWarning `json:"warnings,omitempty"`
	ErrorsFixed   []ValidationError   `json:"errors_fixed,omitempty"`
	WarningsFixed []ValidationWarning `json:"warnings_fixed,omitempty"`
	Timestamp     time.Time           `json:"timestamp"`
}

// ValidationError represents a validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

// ValidationWarning represents a validation warning
type ValidationWarning struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

// Validator validates TMF objects against requirements and optionally fixes them
type Validator struct {
	config *Config
}

// NewValidator creates a new validator
func NewValidator(config *Config) *Validator {
	return &Validator{
		config: config,
	}
}

// ValidateAndFixObject validates a single TMF object and optionally fixes validation errors
func (v *Validator) ValidateAndFixObject(obj TMFObject, objectType string) ValidationResult {
	result := ValidationResult{
		ObjectID:   obj.GetID(),
		ObjectType: objectType,
		Valid:      true,
		Timestamp:  time.Now(),
	}

	// Validate required fields
	if v.config.ValidateRequiredFields {
		v.validateRequiredFields(obj, objectType, &result)
	}

	// Validate related party requirements
	if v.config.ValidateRelatedParty {
		if v.config.Version == "v4" {
			v.validateRelatedPartyV4(obj, objectType, &result)
		} else {
			v.validateRelatedPartyV5(obj, objectType, &result)
		}
	}

	// Determine overall validity (object is valid if it has zero validation errors)
	result.Valid = len(result.Errors) == 0

	return result
}

// validateRequiredFields checks if all required fields are present and optionally fixes them
func (v *Validator) validateRequiredFields(obj TMFObject, objectType string, result *ValidationResult) {
	// Required fields for all objects
	requiredFields := []string{"id", "href", "lastUpdate", "version"}

	for _, field := range requiredFields {
		if !obj.HasField(field) {
			if v.config.FixValidationErrors {
				// Attempt to fix the missing field
				if v.fixMissingRequiredField(obj, field) {
					result.ErrorsFixed = append(result.ErrorsFixed, ValidationError{
						Field:   field,
						Message: fmt.Sprintf("Required field '%s' was missing and has been fixed", field),
						Code:    "MISSING_REQUIRED_FIELD_FIXED",
					})
					continue
				}
			}

			result.Errors = append(result.Errors, ValidationError{
				Field:   field,
				Message: fmt.Sprintf("Required field '%s' is missing", field),
				Code:    "MISSING_REQUIRED_FIELD",
			})
		}
	}
}

// fixMissingRequiredField attempts to fix a missing required field
func (v *Validator) fixMissingRequiredField(obj TMFObject, field string) bool {
	switch field {
	case "lastUpdate":
		// Set to current timestamp
		obj.SetLastUpdateNow()
		return true
	case "version":
		// Set default version
		obj.SetVersion("1.0")
		return true
	}
	return false
}

// validateRelatedPartyV4 checks if required related party roles are present and optionally fixes them
func (v *Validator) validateRelatedPartyV4(obj TMFObject, objectType string, result *ValidationResult) {
	// Check if object type requires related party
	if slices.Contains(DoNotRequireRelatedParties, objectType) {
		return
	}

	// Check if related party information exists
	if !obj.HasRelatedParty() {
		if v.config.FixValidationErrors {
			// Attempt to add default related party information
			if v.fixMissingRelatedPartyV4(obj, objectType) {
				result.ErrorsFixed = append(result.ErrorsFixed, ValidationError{
					Field:   "relatedParty",
					Message: "Related party information was missing and has been fixed",
					Code:    "MISSING_RELATED_PARTY_FIXED",
				})
				return
			}
		}

		result.Errors = append(result.Errors, ValidationError{
			Field:   "relatedParty",
			Message: "Related party information is required but missing",
			Code:    "MISSING_RELATED_PARTY",
		})
		return
	}

	relatedParties := obj.GetRelatedParty()

	// Check for required roles
	requiredRoles := []string{"seller", "selleroperator", "buyer", "buyeroperator"}
	if slices.Contains(DoNotRequireBuyerInfo, objectType) {
		requiredRoles = []string{"seller", "selleroperator"}
	}

	foundRoles := make(map[string]bool)

	// Validate individual related party entries
	for i, rp := range relatedParties {
		role := strings.ToLower(TMFObject(rp).GetStringField("role"))
		if role != "" {
			foundRoles[role] = true
		}

		// Validate individual fields
		v.validateRelatedPartyEntryV4(rp, i, result)
	}

	// Check for missing required roles
	for _, requiredRole := range requiredRoles {
		if !foundRoles[requiredRole] {
			if v.config.FixValidationErrors {
				// Attempt to add missing role
				if v.fixMissingRoleV4(obj, requiredRole) {
					result.ErrorsFixed = append(result.ErrorsFixed, ValidationError{
						Field:   "relatedParty",
						Message: fmt.Sprintf("Required related party role '%s' was missing and has been fixed", requiredRole),
						Code:    "MISSING_REQUIRED_ROLE_FIXED",
					})
					continue
				}
			}

			result.Errors = append(result.Errors, ValidationError{
				Field:   "relatedParty",
				Message: fmt.Sprintf("Required related party role '%s' is missing", requiredRole),
				Code:    "MISSING_REQUIRED_ROLE",
			})
		}
	}
}

// validateRelatedPartyV5 checks if required related party roles are present and optionally fixes them
func (v *Validator) validateRelatedPartyV5(obj TMFObject, objectType string, result *ValidationResult) {
	// Check if object type requires related party
	if slices.Contains(DoNotRequireRelatedParties, objectType) {
		return
	}

	// Check if related party information exists
	if !obj.HasRelatedParty() {
		if v.config.FixValidationErrors {
			// Attempt to add default related party information
			if v.fixMissingRelatedPartyV5(obj, objectType) {
				result.ErrorsFixed = append(result.ErrorsFixed, ValidationError{
					Field:   "relatedParty",
					Message: "Related party information was missing and has been fixed",
					Code:    "MISSING_RELATED_PARTY_FIXED",
				})
				return
			}
		}

		result.Errors = append(result.Errors, ValidationError{
			Field:   "relatedParty",
			Message: "Related party information is required but missing",
			Code:    "MISSING_RELATED_PARTY",
		})
		return
	}

	relatedParties := obj.GetRelatedParty()

	// Check for required roles
	requiredRoles := RequiredRelatedPartyRoles[objectType]
	if len(requiredRoles) == 0 {
		return
	}

	foundRoles := make(map[string]bool)

	// Validate individual related party entries
	for i, rp := range relatedParties {
		role := TMFObject(rp).GetStringField("role")
		if role != "" {
			foundRoles[role] = true
		}

		// Validate individual fields
		v.validateRelatedPartyEntryV5(rp, i, result)
	}

	// Check for missing required roles
	for _, requiredRole := range requiredRoles {
		if !foundRoles[requiredRole] {
			if v.config.FixValidationErrors {
				// Attempt to add missing role
				if v.fixMissingRoleV5(obj, requiredRole) {
					result.ErrorsFixed = append(result.ErrorsFixed, ValidationError{
						Field:   "relatedParty",
						Message: fmt.Sprintf("Required related party role '%s' was missing and has been fixed", requiredRole),
						Code:    "MISSING_REQUIRED_ROLE_FIXED",
					})
					continue
				}
			}

			result.Warnings = append(result.Warnings, ValidationWarning{
				Field:   "relatedParty",
				Message: fmt.Sprintf("Required related party role '%s' is missing", requiredRole),
				Code:    "MISSING_REQUIRED_ROLE",
			})
		}
	}
}

// validateRelatedPartyEntryV4 validates a single related party entry for V4
func (v *Validator) validateRelatedPartyEntryV4(rp TMFObject, index int, result *ValidationResult) {
	role := strings.ToLower(rp.GetStringField("role"))

	// Only validate seller/buyer roles
	if role == "seller" || role == "selleroperator" || role == "buyer" || role == "buyeroperator" {
		fields := []string{"id", "href", "name", "@referredType"}
		for _, field := range fields {
			if !rp.HasField(field) {
				if v.config.FixValidationErrors {
					// Attempt to fix the missing field
					if v.fixMissingRelatedPartyFieldV4(rp, field, role) {
						result.ErrorsFixed = append(result.ErrorsFixed, ValidationError{
							Field:   fmt.Sprintf("relatedParty[%d].%s", index, field),
							Message: fmt.Sprintf("Related party %s %s was missing and has been fixed", role, field),
							Code:    "MISSING_PARTY_FIELD_FIXED",
						})
						continue
					}
				}

				result.Errors = append(result.Errors, ValidationError{
					Field:   fmt.Sprintf("relatedParty[%d].%s", index, field),
					Message: fmt.Sprintf("Related party %s %s is missing", role, field),
					Code:    "MISSING_PARTY_FIELD",
				})
			}
		}
	}
}

// validateRelatedPartyEntryV5 validates a single related party entry for V5
func (v *Validator) validateRelatedPartyEntryV5(rp TMFObject, index int, result *ValidationResult) {
	// Check role
	if !rp.HasField("role") {
		if v.config.FixValidationErrors {
			if v.fixMissingRoleFieldV5(rp) {
				result.WarningsFixed = append(result.WarningsFixed, ValidationWarning{
					Field:   fmt.Sprintf("relatedParty[%d].role", index),
					Message: "Related party role was missing and has been fixed",
					Code:    "EMPTY_ROLE_FIXED",
				})
			} else {
				result.Warnings = append(result.Warnings, ValidationWarning{
					Field:   fmt.Sprintf("relatedParty[%d].role", index),
					Message: "Related party role is empty",
					Code:    "EMPTY_ROLE",
				})
			}
		} else {
			result.Warnings = append(result.Warnings, ValidationWarning{
				Field:   fmt.Sprintf("relatedParty[%d].role", index),
				Message: "Related party role is empty",
				Code:    "EMPTY_ROLE",
			})
		}
	}

	// Check partyOrPartyRole
	partyOrPartyRole := rp.GetMapField("partyOrPartyRole")
	if partyOrPartyRole == nil {
		if v.config.FixValidationErrors {
			if v.fixMissingPartyOrPartyRoleV5(rp) {
				result.WarningsFixed = append(result.WarningsFixed, ValidationWarning{
					Field:   fmt.Sprintf("relatedParty[%d].partyOrPartyRole", index),
					Message: "Related party partyOrPartyRole was missing and has been fixed",
					Code:    "MISSING_PARTY_OR_PARTY_ROLE_FIXED",
				})
			}
		}
		return
	}

	// Check partyOrPartyRole fields
	fields := []string{"id", "href"}
	for _, field := range fields {
		if !TMFObject(partyOrPartyRole).HasField(field) {
			if v.config.FixValidationErrors {
				if v.fixMissingPartyOrPartyRoleFieldV5(partyOrPartyRole, field) {
					result.WarningsFixed = append(result.WarningsFixed, ValidationWarning{
						Field:   fmt.Sprintf("relatedParty[%d].partyOrPartyRole.%s", index, field),
						Message: fmt.Sprintf("Related party partyOrPartyRole %s was missing and has been fixed", field),
						Code:    "MISSING_PARTY_FIELD_FIXED",
					})
					continue
				}
			}

			result.Warnings = append(result.Warnings, ValidationWarning{
				Field:   fmt.Sprintf("relatedParty[%d].partyOrPartyRole.%s", index, field),
				Message: fmt.Sprintf("Related party partyOrPartyRole %s is missing", field),
				Code:    "MISSING_PARTY_FIELD",
			})
		}
	}
}

// Fixing methods for V4

func (v *Validator) fixMissingRelatedPartyV4(obj TMFObject, objectType string) bool {
	// Add default seller role
	sellerRP := TMFObject{
		"role":          "seller",
		"id":            "urn:ngsi-ld:organization:default-seller",
		"href":          "urn:ngsi-ld:organization:default-seller",
		"name":          "Default Seller",
		"@referredType": "Organization",
	}
	obj.AddRelatedParty(sellerRP)

	// Add buyer role if required
	if !slices.Contains(DoNotRequireBuyerInfo, objectType) {
		buyerRP := TMFObject{
			"role":          "buyer",
			"id":            "urn:ngsi-ld:organization:default-buyer",
			"href":          "urn:ngsi-ld:organization:default-buyer",
			"name":          "Default Buyer",
			"@referredType": "Organization",
		}
		obj.AddRelatedParty(buyerRP)
	}

	return true
}

func (v *Validator) fixMissingRoleV4(obj TMFObject, role string) bool {
	newRP := TMFObject{
		"role":          role,
		"id":            fmt.Sprintf("urn:ngsi-ld:organization:default-%s", role),
		"href":          fmt.Sprintf("urn:ngsi-ld:organization:default-%s", role),
		"name":          fmt.Sprintf("Default %s", strings.Title(role)),
		"@referredType": "Organization",
	}
	obj.AddRelatedParty(newRP)
	return true
}

func (v *Validator) fixMissingRelatedPartyFieldV4(rp TMFObject, field, role string) bool {
	switch field {
	case "id":
		rp.SetStringField("id", fmt.Sprintf("urn:ngsi-ld:organization:default-%s", role))
		return true
	case "href":
		rp.SetStringField("href", fmt.Sprintf("urn:ngsi-ld:organization:default-%s", role))
		return true
	case "name":
		rp.SetStringField("name", fmt.Sprintf("Default %s", strings.Title(role)))
		return true
	case "@referredType":
		rp.SetStringField("@referredType", "Organization")
		return true
	}
	return false
}

// Fixing methods for V5

func (v *Validator) fixMissingRelatedPartyV5(obj TMFObject, objectType string) bool {
	// Add default seller role
	sellerRP := TMFObject{
		"@type": "RelatedParty",
		"role":  "seller",
		"partyOrPartyRole": map[string]any{
			"id":    "urn:ngsi-ld:organization:default-seller",
			"href":  "urn:ngsi-ld:organization:default-seller",
			"name":  "Default Seller",
			"@type": "OrganizationRef",
		},
	}
	obj.AddRelatedParty(sellerRP)

	// Add buyer role if required
	requiredRoles := RequiredRelatedPartyRoles[objectType]
	if len(requiredRoles) > 1 { // Assuming seller is always first
		buyerRP := TMFObject{
			"@type": "RelatedParty",
			"role":  "buyer",
			"partyOrPartyRole": map[string]any{
				"id":    "urn:ngsi-ld:organization:default-buyer",
				"href":  "urn:ngsi-ld:organization:default-buyer",
				"name":  "Default Buyer",
				"@type": "OrganizationRef",
			},
		}
		obj.AddRelatedParty(buyerRP)
	}

	return true
}

func (v *Validator) fixMissingRoleV5(obj TMFObject, role string) bool {
	newRP := TMFObject{
		"@type": "RelatedParty",
		"role":  role,
		"partyOrPartyRole": map[string]any{
			"id":    fmt.Sprintf("urn:ngsi-ld:organization:default-%s", role),
			"href":  fmt.Sprintf("urn:ngsi-ld:organization:default-%s", role),
			"name":  fmt.Sprintf("Default %s", strings.Title(role)),
			"@type": "OrganizationRef",
		},
	}
	obj.AddRelatedParty(newRP)
	return true
}

func (v *Validator) fixMissingRoleFieldV5(rp TMFObject) bool {
	rp.SetStringField("role", "unknown")
	return true
}

func (v *Validator) fixMissingPartyOrPartyRoleV5(rp TMFObject) bool {
	partyOrPartyRole := map[string]any{
		"id":    "urn:ngsi-ld:organization:default",
		"href":  "urn:ngsi-ld:organization:default",
		"name":  "Default Organization",
		"@type": "OrganizationRef",
	}
	rp.SetMapField("partyOrPartyRole", partyOrPartyRole)
	return true
}

func (v *Validator) fixMissingPartyOrPartyRoleFieldV5(partyOrPartyRole map[string]any, field string) bool {
	switch field {
	case "id":
		partyOrPartyRole["id"] = "urn:ngsi-ld:organization:default"
		return true
	case "href":
		partyOrPartyRole["href"] = "urn:ngsi-ld:organization:default"
		return true
	}
	return false
}

// Constants for validation rules (copied from reporting package)
var DoNotRequireRelatedParties = []string{
	"productOfferingPrice",
	"category",
	"individual",
	"organization",
	"productCatalog",
	"customer",
	"product",
	"service",
}

var DoNotRequireBuyerInfo = []string{
	"productOfferingPrice",
	"category",
	"individual",
	"organization",
	"productCatalog",
	"customer",
	"product",
	"service",
}

var RequiredRelatedPartyRoles = map[string][]string{
	"productOffering":      {"seller", "buyer"},
	"productSpecification": {"seller", "buyer"},
	// Add more as needed
}
