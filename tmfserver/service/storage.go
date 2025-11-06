package service

import (
	"net/url"

	repo "github.com/hesusruiz/isbetmf/tmfserver/repository"
)

// TMFStorage abstracts persistence operations for TMF objects.
// It is used for plugging-in different persistence systems
type TMFStorage interface {
	CreateObject(obj *repo.TMFRecord) error
	GetObject(id, objectType string) (*repo.TMFRecord, error)
	UpdateObject(obj *repo.TMFRecord) error
	DeleteObject(id, objectType string) error
	ListObjects(objectType string, apiVersion string, queryParams url.Values) ([]repo.TMFRecord, int, error)
}
