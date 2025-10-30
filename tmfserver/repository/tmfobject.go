package repository

import (
	"strings"
	"time"

	"github.com/hesusruiz/isbetmf/internal/errl"
)

// TMFObject represents a generic TMForum object.
// It is used to store and retrieve objects from the database.
type TMFObject struct {
	ID          string           `db:"id"`
	Type        string           `db:"type"`
	Version     string           `db:"version"`
	APIVersion  string           `db:"api_version"`
	Seller      string           `db:"seller"`
	Buyer       string           `db:"buyer"`
	LastUpdate  string           `db:"last_update"`
	Random      int              `db:"random"`
	Content     []byte           `db:"content"`
	ContentMap  TMFObjectMap     `db:"-"`
	Validations ValidationResult `db:"-"`
	CreatedAt   time.Time        `db:"created_at"`
	UpdatedAt   time.Time        `db:"updated_at"`
}

// NewTMFObject creates a new TMFObject.
func NewTMFObject(id, objectType, version, apiVersion, lastUpdate string, content []byte) *TMFObject {
	now := time.Now()

	o := &TMFObject{
		ID:         id,
		Type:       objectType,
		Version:    version,
		APIVersion: apiVersion,
		LastUpdate: lastUpdate,
		Content:    content,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	return o

}

func (o *TMFObject) SetBuyerID(buyerID string) {
	if !strings.HasPrefix(buyerID, "did:elsi:") {
		buyerID = "did:elsi:" + buyerID
	}
	o.Buyer = buyerID
}

func (o *TMFObject) SetSellerID(sellerID string) {
	if !strings.HasPrefix(sellerID, "did:elsi:") {
		sellerID = "did:elsi:" + sellerID
	}
	o.Seller = sellerID
}

// Validate validates the TMFObject.
// It returns an error if the object is invalid, and detailed validations are in o.Validations
func (o *TMFObject) Validate() error {
	// We need to unmarshal the JSON to a Map
	// Doing this we get the validations performed and the result in o.Validations
	_, err := o.ToMap()
	return err
}

// ToMap converts the TMFObject to a map[string]any.
func (o *TMFObject) ToMap() (TMFObjectMap, error) {
	if o == nil {
		return nil, errl.Errorf("object is nil")
	}
	if o.ContentMap != nil {
		return o.ContentMap, nil
	}

	omap, err := NewTMFObjectMap(o.Content)
	if err != nil {
		o.Validations.ObjectID = o.ID
		o.Validations.ObjectType = o.Type
		o.Validations.Valid = false
		o.Validations.Timestamp = time.Now()
		o.Validations.Errors = append(o.Validations.Errors, ValidationError{
			Field:   "object",
			Message: errl.Error(err).Error(),
			Code:    "INVALID_OBJECT",
		})

		err = errl.Errorf("failed to unmarshal object content: %w", err)
		return nil, err
	}

	o.Validations = omap.Validate(o.Type)
	if len(o.Validations.Errors) > 0 {
		return nil, errl.Errorf("invalid object")
	}

	o.ContentMap = omap

	return o.ContentMap, nil
}
