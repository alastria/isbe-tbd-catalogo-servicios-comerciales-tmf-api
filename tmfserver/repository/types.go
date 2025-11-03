package repository

import (
	"encoding/json"
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

func (vr *ValidationResult) String() string {
	var buf strings.Builder
	buf.WriteString("ValidationError{")
	buf.WriteString("ObjectID=")
	buf.WriteString(vr.ObjectID)

	out, _ := json.Marshal(vr.Errors)

	buf.WriteString(string(out))
	buf.WriteString("}")
	return buf.String()
}

var RequiredFieldsForAllObjects = []string{
	"id", "href", "lastUpdate",
}

var RecommendedFieldsForAllObjects = []string{
	"name", "version",
}

var DoNotRequireRelatedParties = []string{
	"category",
	"individual",
	"organization",
}

var DoNotRequireBuyerInfo = []string{
	"catalog",
	"productoffering",
	"productspecification",
	"productofferingprice",
	"resourcespecification",
	"servicespecification",
}
