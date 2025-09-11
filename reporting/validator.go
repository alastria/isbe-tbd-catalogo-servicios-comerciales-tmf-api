package reporting

import (
	"encoding/json"
	"fmt"
	"slices"
	"strings"
	"time"
)

// Validator validates TMF objects against requirements
type Validator struct {
	config *Config
}

// NewValidator creates a new validator
func NewValidator(config *Config) *Validator {
	return &Validator{
		config: config,
	}
}

// ValidateObject validates a single TMF object (legacy method for backward compatibility)
func (v *Validator) ValidateObject(obj TMFObject, objectType string) ValidationResult {
	return v.ValidateAndFixObject(&obj, objectType)
}

// ValidateAndFixObject validates a single TMF object and optionally fixes validation errors
func (v *Validator) ValidateAndFixObject(obj *TMFObject, objectType string) ValidationResult {
	result := ValidationResult{
		ObjectID:   obj.ID,
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
		if v.config.Version == VersionV4 {
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
func (v *Validator) validateRequiredFields(obj *TMFObject, objectType string, result *ValidationResult) {

	// This checks the fields that are required for all objects
	for _, field := range RequiredFieldsForAllObjects {
		if !v.hasField(obj, field) {
			// TODO: Implement fixing logic for missing required fields
			// If v.config.FixValidationErrors is true, attempt to fix the missing field
			// and move the error to result.ErrorsFixed if successfully fixed

			result.Errors = append(result.Errors, ValidationError{
				Field:   field,
				Message: fmt.Sprintf("Required field '%s' is missing", field),
				Code:    "MISSING_REQUIRED_FIELD",
			})
		}
	}

}

// validateRelatedPartyV5 checks if required related party roles are present and optionally fixes them
func (v *Validator) validateRelatedPartyV5(obj *TMFObject, objectType string, result *ValidationResult) {
	// We just return if the object does not require any Related Party
	if slices.Contains(DoNotRequireRelatedParties, objectType) {
		return
	}

	// Unmarshall the raw messages into RelatedPartyV5
	var relatedParties = []RelatedPartyV5{}
	json.Unmarshal(obj.RelatedParty, &relatedParties)

	// The object requires some related party. The detailed verification will be done
	if len(relatedParties) == 0 {
		// TODO: Implement fixing logic for missing related party information
		// If v.config.FixValidationErrors is true, attempt to add default related party information
		// and move the error to result.ErrorsFixed if successfully fixed

		result.Errors = append(result.Errors, ValidationError{
			Field:   "relatedParty",
			Message: "Related party information is required but missing",
			Code:    "MISSING_RELATED_PARTY",
		})
	}

	requiredRoles := RequiredRelatedPartyRoles[objectType]
	if len(requiredRoles) == 0 {
		return
	}

	if len(relatedParties) == 0 {
		result.Errors = append(result.Errors, ValidationError{
			Field:   "relatedParty",
			Message: "Related party information is required but missing",
			Code:    "MISSING_RELATED_PARTY",
		})
		return
	}

	// Check for required roles
	foundRoles := make(map[string]bool)
	for _, rp := range relatedParties {
		if rp.Role != "" {
			foundRoles[rp.Role] = true
		}
	}

	for _, requiredRole := range requiredRoles {
		if !foundRoles[requiredRole] {
			// TODO: Implement fixing logic for missing required roles
			// If v.config.FixValidationErrors is true, attempt to add default role information
			// and move the warning to result.WarningsFixed if successfully fixed

			result.Warnings = append(result.Warnings, ValidationWarning{
				Field:   "relatedParty",
				Message: fmt.Sprintf("Required related party role '%s' is missing", requiredRole),
				Code:    "MISSING_REQUIRED_ROLE",
			})
		}
	}

	// Validate individual related party entries
	for i, rp := range relatedParties {
		v.validateRelatedPartyEntry(rp, i, result)
	}
}

// validateRelatedPartyEntry validates a single related party entry and optionally fixes issues
func (v *Validator) validateRelatedPartyEntry(rp RelatedPartyV5, index int, result *ValidationResult) {
	if rp.Role == "" {
		// TODO: Implement fixing logic for empty role
		// If v.config.FixValidationErrors is true, attempt to set a default role
		// and move the warning to result.WarningsFixed if successfully fixed

		result.Warnings = append(result.Warnings, ValidationWarning{
			Field:   fmt.Sprintf("relatedParty[%d].role", index),
			Message: "Related party role is empty",
			Code:    "EMPTY_ROLE",
		})
	}

	if rp.PartyOrPartyRole.ID == "" {
		// TODO: Implement fixing logic for missing party ID
		// If v.config.FixValidationErrors is true, attempt to generate or set a default ID
		// and move the warning to result.WarningsFixed if successfully fixed

		result.Warnings = append(result.Warnings, ValidationWarning{
			Field:   fmt.Sprintf("relatedParty[%d].partyOrPartyRole.id", index),
			Message: "Related party ID is missing",
			Code:    "MISSING_PARTY_ID",
		})
	}

	if rp.PartyOrPartyRole.Href == "" {
		// TODO: Implement fixing logic for missing party href
		// If v.config.FixValidationErrors is true, attempt to generate or set a default href
		// and move the warning to result.WarningsFixed if successfully fixed

		result.Warnings = append(result.Warnings, ValidationWarning{
			Field:   fmt.Sprintf("relatedParty[%d].partyOrPartyRole.href", index),
			Message: "Related party href is missing",
			Code:    "MISSING_PARTY_HREF",
		})
	}
}

// hasField checks if an object has a specific field
func (v *Validator) hasField(obj *TMFObject, field string) bool {
	switch field {
	case "id":
		return obj.ID != ""
	case "href":
		return obj.Href != ""
	case "lastUpdate":
		return obj.LastUpdate != ""
	case "version":
		return obj.Version != ""
	case "relatedParty":
		return len(obj.RelatedParty) > 0
	default:
		// Check additional fields
		_, exists := obj.AdditionalFields[field]
		return exists
	}
}

// ValidateObjects validates multiple TMF objects
func (v *Validator) ValidateObjects(objects []TMFObject, objectType string) []ValidationResult {
	results := make([]ValidationResult, len(objects))
	for i, obj := range objects {
		results[i] = v.ValidateObject(obj, objectType)
	}
	return results
}

func (v *Validator) validateRelatedPartyV4(obj *TMFObject, objectType string, result *ValidationResult) {
	if obj.ID == "urn:ngsi-ld:applied-customer-billing-rate:a886304d-d699-4adf-b93e-dcdcd54474f1" {
		fmt.Println("urn:ngsi-ld:applied-customer-billing-rate:a886304d-d699-4adf-b93e-dcdcd54474f1")
	}

	// We just return if the object does not require any Related Party
	if slices.Contains(DoNotRequireRelatedParties, objectType) {
		return
	}

	// Unmarshall the raw messages into RelatedPartyV4
	var relatedParties = []RelatedPartyV4{}
	json.Unmarshal(obj.RelatedParty, &relatedParties)

	// The object requires some related party.
	if len(relatedParties) == 0 {
		// TODO: Implement fixing logic for missing related party information (V4)
		// If v.config.FixValidationErrors is true, attempt to add default related party information
		// and move the error to result.ErrorsFixed if successfully fixed

		result.Errors = append(result.Errors, ValidationError{
			Field:   "relatedParty",
			Message: "Related party information is required but missing",
			Code:    "MISSING_RELATED_PARTY",
		})
		return
	}

	// Check for required roles
	foundRoles := make(map[string]bool)
	// Validate individual related party entries
	for _, rp := range relatedParties {
		// Normalize the role to lowercase for comparison
		rp.Role = strings.ToLower(rp.Role)

		// Process each type of role found
		if rp.Role == "seller" || rp.Role == "selleroperator" || rp.Role == "buyer" || rp.Role == "buyeroperator" {
			foundRoles[rp.Role] = true
			if rp.ID == "" {
				// TODO: Implement fixing logic for missing party ID (V4)
				// If v.config.FixValidationErrors is true, attempt to generate or set a default ID
				// and move the error to result.ErrorsFixed if successfully fixed

				result.Errors = append(result.Errors, ValidationError{
					Field:   fmt.Sprintf("relatedParty %s.id", rp.Role),
					Message: "Related party ID is missing",
					Code:    "MISSING_PARTY_ID",
				})
			}

			if rp.Href == "" {
				// TODO: Implement fixing logic for missing party href (V4)
				// If v.config.FixValidationErrors is true, attempt to generate or set a default href
				// and move the error to result.ErrorsFixed if successfully fixed

				result.Errors = append(result.Errors, ValidationError{
					Field:   fmt.Sprintf("relatedParty %s.href", rp.Role),
					Message: "Related party href is missing",
					Code:    "MISSING_PARTY_HREF",
				})
			}

			if rp.Name == "" {
				// TODO: Implement fixing logic for missing party name (V4)
				// If v.config.FixValidationErrors is true, attempt to generate or set a default name
				// and move the error to result.ErrorsFixed if successfully fixed

				result.Errors = append(result.Errors, ValidationError{
					Field:   fmt.Sprintf("relatedParty %s.name", rp.Role),
					Message: "Related party name is missing",
					Code:    "MISSING_PARTY_NAME",
				})
			}

			if rp.ReferredType == "" {
				// TODO: Implement fixing logic for missing party referred type (V4)
				// If v.config.FixValidationErrors is true, attempt to generate or set a default referred type
				// and move the error to result.ErrorsFixed if successfully fixed

				result.Errors = append(result.Errors, ValidationError{
					Field:   fmt.Sprintf("relatedParty %s.referredType", rp.Role),
					Message: "Related party referred type is missing",
					Code:    "MISSING_PARTY_REFERRED_TYPE",
				})
			}

		}
	}

	// Set the required roles depending on the type of object
	requiredRoles := []string{"seller", "selleroperator", "buyer", "buyeroperator"}
	if slices.Contains(DoNotRequireBuyerInfo, objectType) {
		requiredRoles = []string{"seller", "selleroperator"}
	}

	// Check if we found all the required roles
	for _, requiredRole := range requiredRoles {
		if !foundRoles[requiredRole] {
			// TODO: Implement fixing logic for missing required roles (V4)
			// If v.config.FixValidationErrors is true, attempt to add default role information
			// and move the error to result.ErrorsFixed if successfully fixed

			result.Errors = append(result.Errors, ValidationError{
				Field:   "relatedParty",
				Message: fmt.Sprintf("Required related party role '%s' is missing", requiredRole),
				Code:    "MISSING_REQUIRED_ROLE",
			})
		}
	}

}
