package repository

import (
	"encoding/json"
	"time"

	"github.com/hesusruiz/isbetmf/internal/errl"
	"golang.org/x/exp/slog"
)

// TMFObject represents a generic TMForum object.
// It is used to store and retrieve objects from the database.
type TMFObject struct {
	ID         string    `db:"id"`
	Type       string    `db:"type"`
	Version    string    `db:"version"`
	APIVersion string    `db:"api_version"`
	LastUpdate string    `db:"last_update"`
	Content    []byte    `db:"content"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

// NewTMFObject creates a new TMFObject.
func NewTMFObject(id, objectType, version, apiVersion, lastUpdate string, content []byte) *TMFObject {
	now := time.Now()
	return &TMFObject{
		ID:         id,
		Type:       objectType,
		Version:    version,
		APIVersion: apiVersion,
		LastUpdate: lastUpdate,
		Content:    content,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}

// ToMapNoErr converts the TMFObject to a map[string]any.
func (o *TMFObject) ToMapNoErr() map[string]any {
	var data map[string]any

	if o == nil {
		return nil
	}

	err := json.Unmarshal(o.Content, &data)
	if err != nil {
		err = errl.Errorf("failed to unmarshal object content: %w", err)
		slog.Error("Failed to unmarshal object content", slog.Any("error", err))
		return nil
	}
	return data
}

// ToMap converts the TMFObject to a map[string]any.
func (o *TMFObject) ToMap() (map[string]any, error) {
	var data map[string]any

	if o == nil {
		return nil, errl.Errorf("object is nil")
	}

	err := json.Unmarshal(o.Content, &data)
	if err != nil {
		return nil, errl.Errorf("failed to unmarshal object content: %w", err)
	}
	return data, nil
}
