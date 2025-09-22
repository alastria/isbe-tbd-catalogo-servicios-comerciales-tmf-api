package replicate

import (
	"encoding/json"
	"fmt"

	"github.com/hesusruiz/isbetmf/reporting"
	repo "github.com/hesusruiz/isbetmf/tmfserver/repository"
)

// Converter handles conversion between different TMF object representations
type Converter struct{}

// NewConverter creates a new converter
func NewConverter() *Converter {
	return &Converter{}
}

// ReportingTMFObjectToReplicateTMFObject converts a reporting.TMFObject to replicate.TMFObject
func (c *Converter) ReportingTMFObjectToReplicateTMFObject(reportingObj reporting.TMFObject) (TMFObjectMap, error) {
	// Convert the reporting object to JSON first
	jsonData, err := json.Marshal(reportingObj)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal reporting object: %w", err)
	}

	// Create replicate object from JSON
	replicateObj, err := NewTMFObjectMap(jsonData)
	if err != nil {
		return nil, fmt.Errorf("failed to create replicate object: %w", err)
	}

	// Ensure the well-known fields are properly set
	if reportingObj.ID != "" {
		replicateObj.SetID(reportingObj.ID)
	}
	if reportingObj.Href != "" {
		replicateObj.SetHref(reportingObj.Href)
	}
	if reportingObj.LastUpdate != "" {
		replicateObj.SetLastUpdate(reportingObj.LastUpdate)
	}
	if reportingObj.Version != "" {
		replicateObj.SetVersion(reportingObj.Version)
	}
	if reportingObj.Type != "" {
		replicateObj.SetType(reportingObj.Type)
	}

	// Handle related party if present
	if len(reportingObj.RelatedParty) > 0 {
		var relatedParties []any
		if err := json.Unmarshal(reportingObj.RelatedParty, &relatedParties); err == nil {
			replicateObj.SetField("relatedParty", relatedParties)
		}
	}

	return replicateObj, nil
}

// ReplicateTMFObjectToReportingTMFObject converts a replicate.TMFObject to reporting.TMFObject
func (c *Converter) ReplicateTMFObjectToReportingTMFObject(replicateObj TMFObjectMap) (reporting.TMFObject, error) {
	// Convert to JSON first
	jsonData, err := replicateObj.ToJSON()
	if err != nil {
		return reporting.TMFObject{}, fmt.Errorf("failed to marshal replicate object: %w", err)
	}

	// Create reporting object
	var reportingObj reporting.TMFObject
	if err := json.Unmarshal(jsonData, &reportingObj); err != nil {
		return reporting.TMFObject{}, fmt.Errorf("failed to unmarshal to reporting object: %w", err)
	}

	// Ensure well-known fields are properly set
	reportingObj.ID = replicateObj.GetID()
	reportingObj.Href = replicateObj.GetHref()
	reportingObj.LastUpdate = replicateObj.GetLastUpdate()
	reportingObj.Version = replicateObj.GetVersion()
	reportingObj.Type = replicateObj.GetType()

	// Handle related party
	if replicateObj.HasRelatedParty() {
		relatedParties := replicateObj.GetRelatedParty()
		if len(relatedParties) > 0 {
			relatedPartyJSON, err := json.Marshal(relatedParties)
			if err != nil {
				return reporting.TMFObject{}, fmt.Errorf("failed to marshal related party: %w", err)
			}
			reportingObj.RelatedParty = relatedPartyJSON
		}
	}

	return reportingObj, nil
}

// ReplicateTMFObjectToRepositoryTMFObject converts a replicate.TMFObject to repository.TMFObject
func (c *Converter) ReplicateTMFObjectToRepositoryTMFObject(replicateObj TMFObjectMap, objectType string) (*repo.TMFObject, error) {
	// Convert to JSON
	jsonData, err := replicateObj.ToJSON()
	if err != nil {
		return nil, fmt.Errorf("failed to marshal replicate object: %w", err)
	}

	// Create repository object
	repoObj := repo.NewTMFObject(
		replicateObj.GetID(),
		objectType,
		replicateObj.GetVersion(),
		"v4", // Default API version
		replicateObj.GetLastUpdate(),
		jsonData,
	)

	return repoObj, nil
}

// RepositoryTMFObjectToReplicateTMFObject converts a repository.TMFObject to replicate.TMFObject
func (c *Converter) RepositoryTMFObjectToReplicateTMFObject(repoObj *repo.TMFObject) (TMFObjectMap, error) {
	// Create replicate object from repository content
	replicateObj, err := NewTMFObjectMap(repoObj.Content)
	if err != nil {
		return nil, fmt.Errorf("failed to create replicate object from repository content: %w", err)
	}

	// Ensure well-known fields are properly set from repository metadata
	replicateObj.SetID(repoObj.ID)
	replicateObj.SetVersion(repoObj.Version)
	replicateObj.SetLastUpdate(repoObj.LastUpdate)

	return replicateObj, nil
}

// JSONToReplicateTMFObject converts JSON bytes to replicate.TMFObject
func (c *Converter) JSONToReplicateTMFObject(jsonData []byte) (TMFObjectMap, error) {
	return NewTMFObjectMap(jsonData)
}

// ReplicateTMFObjectToJSON converts replicate.TMFObject to JSON bytes
func (c *Converter) ReplicateTMFObjectToJSON(replicateObj TMFObjectMap) ([]byte, error) {
	return replicateObj.ToJSON()
}
