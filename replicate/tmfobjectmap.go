package replicate

import (
	"encoding/json"
	"fmt"
	"maps"
	"time"
)

// TMFObjectMap represents a TMForum object as a map with utility methods
// This is designed to make fixing validation errors simple and efficient
type TMFObjectMap map[string]any

// NewTMFObjectMap creates a new TMFObject from a JSON byte slice
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
