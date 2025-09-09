package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"log/slog"

	"github.com/google/uuid"
	"github.com/hesusruiz/isbetmf/config"
	"github.com/hesusruiz/isbetmf/internal/errl"
	pdp "github.com/hesusruiz/isbetmf/pdp"
	"github.com/hesusruiz/isbetmf/tmfserver/notifications"
	"github.com/hesusruiz/isbetmf/tmfserver/repository"
	repo "github.com/hesusruiz/isbetmf/tmfserver/repository"
	"github.com/jmoiron/sqlx"
)

// AuthUser represents the authenticated user's information from the JWT mandator object.
type AuthUser struct {
	CommonName             string `json:"commonName"`
	Country                string `json:"country"`
	EmailAddress           string `json:"emailAddress"`
	Organization           string `json:"organization"`
	OrganizationIdentifier string `json:"organizationIdentifier"`
	SerialNumber           string `json:"serialNumber"`
	isAuthenticated        bool
	isLEAR                 bool
	isOwner                bool
}

func (u *AuthUser) ToMap() map[string]any {
	return map[string]any{
		"commonName":             u.CommonName,
		"country":                u.Country,
		"emailAddress":           u.EmailAddress,
		"organization":           u.Organization,
		"organizationIdentifier": u.OrganizationIdentifier,
		"serialNumber":           u.SerialNumber,
		"isAuthenticated":        u.isAuthenticated,
		"isLEAR":                 u.isLEAR,
		"isOwner":                u.isOwner,
	}
}

// Request represents a generic HTTP request. Handlers must convert to this representation.
// In this way, we support easily any HTTP framework (currently Fiber and Echo), but also other
// future channels like JSON-RPC or even non-HTTP channels like GRPC.
type Request struct {
	Method       string
	Action       HttpAction
	APIfamily    string
	APIVersion   string
	ResourceName string
	ID           string
	QueryParams  url.Values
	Body         []byte
	AuthUser     *AuthUser
	AccessToken  string
}

func (r *Request) ToMap() map[string]any {
	return map[string]any{
		"method":   r.Method,
		"action":   r.Action,
		"api":      r.APIfamily,
		"version":  r.APIVersion,
		"resource": r.ResourceName,
		"id":       r.ID,
	}
}

type HttpAction string

const (
	READ   HttpAction = "READ"
	CREATE HttpAction = "CREATE"
	UPDATE HttpAction = "UPDATE"
	DELETE HttpAction = "DELETE"
	LIST   HttpAction = "LIST"
)

// These are more friendly names for the writers of policy rules and can be used interchangeably
var HttpActions = map[string]HttpAction{
	"GET":    READ,
	"POST":   CREATE,
	"PATCH":  UPDATE,
	"DELETE": DELETE,
	"LIST":   LIST,
}

// Response represents a generic HTTP response.
type Response struct {
	StatusCode int
	Headers    map[string]string
	Body       any
}

// Service is the service for the API.
type Service struct {

	// The SQL layer on top of the actual storage engine
	db *sqlx.DB

	// Pluggable storage backend (optional). When nil, falls back to built-in SQLite via db
	storage Storage

	// The rules engine in Starlark
	ruleEngine *pdp.PDP

	// The Verifier server which signs the Access Tokens,
	// and the PDP retrieves the JWKS from it to verify the signatures.
	verifierServer string

	// The OpenID configuration of the Verifier Server
	oid *OpenIDConfig

	// Notifications manager
	notif *notifications.Manager
}

// NewService creates a new service.
func NewService(db *sqlx.DB, ruleEngine *pdp.PDP, verifierServer string) *Service {
	svc := &Service{
		db:             db,
		ruleEngine:     ruleEngine,
		verifierServer: verifierServer,
	}

	err := svc.initializeService()
	if err != nil {
		panic(err)
	}

	// Initialize notifications with in-memory store and HTTP delivery
	store := notifications.NewMemoryStore()
	deliver := notifications.NewHTTPDelivery(5 * time.Second)
	svc.notif = notifications.NewManager(store, deliver)

	return svc
}

func (svc *Service) initializeService() error {

	// Create the server operator identity, in case it is not yet in the database
	org := &repository.Organization{
		CommonName:             config.ServerOperatorName,
		Country:                config.ServerOperatorCountry,
		EmailAddress:           "",
		Organization:           config.ServerOperatorName,
		OrganizationIdentifier: config.ServerOperatorOrganizationIdentifier,
	}

	obj, _ := repository.TMFOrganizationFromToken(nil, org)

	if err := svc.createObject(obj); err != nil {
		if errors.Is(err, &ErrObjectExists{}) {
			slog.Debug("server operator organization already exists", "organizationIdentifier", config.ServerOperatorOrganizationIdentifier)
		} else {
			err = errl.Error(err)
			panic("error creatingserver operator organization")
		}
	}

	// Retrieve the OpenId configuration of the Verifier server
	oid, err := NewOpenIDConfig(svc.verifierServer)
	if err != nil {
		return errl.Errorf("failed to retrieve OpenID configuration: %w", err)
	}
	svc.oid = oid

	return nil

}

// CreateHubSubscription creates a new notification subscription (hub) for an API family.
func (svc *Service) CreateHubSubscription(req *Request) *Response {
	// Authenticate like write operations
	_, err := svc.processAccessToken(req)
	if err != nil {
		return ErrorResponsef(err, http.StatusUnauthorized, "invalid access token: %w")
	}

	// Parse incoming body
	var body map[string]any
	if err := json.Unmarshal(req.Body, &body); err != nil {
		return ErrorResponsef(err, http.StatusBadRequest, "failed to bind request body: %w")
	}

	callback, _ := body["callback"].(string)
	if callback == "" {
		return ErrorResponsef(nil, http.StatusBadRequest, "callback is required")
	}

	var eventTypes []string
	if raw, ok := body["eventTypes"].([]any); ok {
		for _, v := range raw {
			if s, ok := v.(string); ok {
				eventTypes = append(eventTypes, s)
			}
		}
	}

	headers := make(map[string]string)
	if hmap, ok := body["headers"].(map[string]any); ok {
		for k, v := range hmap {
			if sv, ok := v.(string); ok {
				headers[strings.ToLower(k)] = sv
			}
		}
	}

	query, _ := body["query"].(string)

	// Build subscription
	id := uuid.NewString()
	sub := &notifications.Subscription{
		ID:         id,
		APIFamily:  req.APIfamily,
		Callback:   callback,
		EventTypes: eventTypes,
		Headers:    headers,
		Query:      query,
	}

	_, err = svc.notif.CreateSubscription(req.APIfamily, sub)
	if err != nil {
		return ErrorResponsef(err, http.StatusInternalServerError, "failed to create subscription: %w")
	}

	resp := map[string]any{
		"id":         sub.ID,
		"callback":   sub.Callback,
		"eventTypes": sub.EventTypes,
		"query":      sub.Query,
		"href":       fmt.Sprintf("/tmf-api/%s/%s/hub/%s", req.APIfamily, req.APIVersion, sub.ID),
	}
	if token, ok := sub.Headers["x-auth-token"]; ok && token != "" {
		resp["headers"] = map[string]any{"x-auth-token": token}
	}

	return &Response{StatusCode: http.StatusCreated, Body: resp}
}

// DeleteHubSubscription deletes a subscription by id for an API family.
func (svc *Service) DeleteHubSubscription(req *Request) *Response {
	// Authenticate like write operations
	_, err := svc.processAccessToken(req)
	if err != nil {
		return ErrorResponsef(err, http.StatusUnauthorized, "invalid access token: %w")
	}

	if req.ID == "" {
		return ErrorResponsef(nil, http.StatusBadRequest, "id is required")
	}

	if err := svc.notif.DeleteSubscription(req.APIfamily, req.ID); err != nil {
		return ErrorResponsef(err, http.StatusNotFound, "failed to delete subscription: %w")
	}

	return &Response{StatusCode: http.StatusNoContent}
}

// CreateGenericObject creates a new TMF object using generalized parameters.
func (svc *Service) CreateGenericObject(req *Request) *Response {
	var err error
	slog.Debug("CreateGenericObject called", slog.String("apiFamily", req.APIfamily), slog.String("resourceName", req.ResourceName))

	// ************************************************************************************************
	// Authentication: we require the user to be authenticated
	// ************************************************************************************************

	var token map[string]any
	{
		// Process the AccessToken to extract caller info from its claims in the payload
		if token, err = svc.processAccessToken(req); err != nil {
			return ErrorResponsef(err, http.StatusUnauthorized, "invalid access token: %w")
		}

		// This operation can not be done without authentication
		if len(token) == 0 {
			return ErrorResponsef(nil, http.StatusUnauthorized, "user not authenticated")
		}

	}

	// ************************************************************************************************
	// Parse the request body, which contains the TMForum object being created.
	// We perform some formal verifications and add default values if needed.
	// ************************************************************************************************

	var incomingObjectMap map[string]any
	var id string
	var version string
	var lastUpdate string
	var href string
	{

		// Parse the request body
		if err := json.Unmarshal(req.Body, &incomingObjectMap); err != nil {
			return ErrorResponsef(err, http.StatusBadRequest, "failed to bind request body: %w")
		}

		id, _ = incomingObjectMap["id"].(string)
		version, _ = incomingObjectMap["version"].(string)

		// Create a new 'id' if the user did not specify it
		if id == "" {
			// If the incoming object does not have an 'id', we generate a new one
			// The format is "urn:ngsi-ld:{resource-in-kebab-case}:{uuid}"
			id = fmt.Sprintf("urn:ngsi-ld:%s:%s", ToKebabCase(req.ResourceName), uuid.NewString())
			incomingObjectMap["id"] = id
			slog.Debug("Generated new ID for object", "id", id)
		}

		// Create a new 'href' if the user did not specify it
		href, _ = incomingObjectMap["href"].(string)
		if href == "" {
			// Add href to the object using the correct API version
			incomingObjectMap["href"] = fmt.Sprintf("/tmf-api/%s/%s/%s/%s", req.APIfamily, req.APIVersion, req.ResourceName, id)
			slog.Debug("Set href", slog.String("href", incomingObjectMap["href"].(string)))
		}

		// Check and process '@type' field
		if typeVal, typeOk := incomingObjectMap["@type"].(string); typeOk {
			if !strings.EqualFold(typeVal, req.ResourceName) {
				return ErrorResponsef(nil, http.StatusBadRequest, "@type mismatch: expected %s, got %s", req.ResourceName, typeVal)
			}
		} else {
			// If @type is not specified, add it
			incomingObjectMap["@type"] = req.ResourceName
			slog.Debug("Added missing @type field", slog.String("type", req.ResourceName))
		}

		// Set default 'version' if not provided by the user
		if version == "" {
			version = "1.0"
			incomingObjectMap["version"] = version // Update data map for content marshaling
			slog.Debug("Set default version", slog.String("version", version))
		}

		// Set the lastUpdate property. We overwrite whatever the user set.
		now := time.Now()
		lastUpdate = now.Format(time.RFC3339Nano)
		incomingObjectMap["lastUpdate"] = lastUpdate

		// Add Seller and Buyer info. We overwrite whatever is in the incoming object, if any
		err = setSellerAndBuyerInfo(incomingObjectMap, req.AuthUser.OrganizationIdentifier, req.APIVersion)
		if err != nil {
			return ErrorResponsef(err, http.StatusInternalServerError, "failed to add Seller and Buyer info: %w")
		}

	}

	// ************************************************************************************************
	// Before performing the action, check if the user can perform the operation on the object,
	// based on the rules defined by the user in the policy engine.
	// ************************************************************************************************

	err = takeDecision(svc.ruleEngine, req, token, incomingObjectMap, nil)
	if err != nil {
		return ErrorResponsef(err, http.StatusForbidden, "user not authorized: %w")
	}

	// ************************************************************************************************
	// Now we can proceed, creating an object in the database.
	// ************************************************************************************************

	incomingContent, err := json.Marshal(incomingObjectMap)
	if err != nil {
		return ErrorResponsef(err, http.StatusInternalServerError, "failed to marshal object content: %w")
	}

	incomingObject := repo.NewTMFObject(id, req.ResourceName, version, req.APIVersion, lastUpdate, incomingContent)

	if err := svc.createObject(incomingObject); err != nil {
		return ErrorResponsef(err, http.StatusInternalServerError, "failed to create object in service: %w")
	}

	headers := make(map[string]string)
	headers["Location"] = incomingObjectMap["href"].(string)
	slog.Info("Object created successfully", slog.String("id", id), slog.String("resourceName", req.ResourceName), slog.String("location", incomingObjectMap["href"].(string)))

	// Send TMForum notification
	eventType := toEventType(req.ResourceName, "CreateEvent")
	eventPayload := buildEventPayload(req, eventType, incomingObjectMap)
	svc.notif.PublishEvent(req.APIfamily, eventType, eventPayload)

	return &Response{
		StatusCode: http.StatusCreated,
		Headers:    headers,
		Body:       incomingObjectMap,
	}
}

// GetGenericObject retrieves a TMF object using generalized parameters.
func (svc *Service) GetGenericObject(req *Request) *Response {
	slog.Debug("GetGenericObject called", slog.String("id", req.ID), slog.String("resourceName", req.ResourceName))

	// Authentication: process the AccessToken to extract caller info from its claims in the payload
	token, err := svc.processAccessToken(req)
	if err != nil {
		return ErrorResponsef(err, http.StatusUnauthorized, "invalid access token: %w")
	}

	// Retrieve the object from the database. If it is not found, we return a 404 error.
	obj, err := svc.getObject(req.ID, req.ResourceName)
	if err != nil {
		return ErrorResponsef(err, http.StatusInternalServerError, "failed to get object from service: %w")
	}

	if obj == nil {
		return ErrorResponsef(nil, http.StatusNotFound, "object not found")
	}

	// ************************************************************************************************
	// Before performing the action, check if the user can perform the operation on the object,
	// based on the rules defined by the user in the policy engine.
	// ************************************************************************************************

	err = takeDecision(svc.ruleEngine, req, token, obj.ToMap(), obj.ToMap())
	if err != nil {
		return ErrorResponse(err, http.StatusForbidden)
	}

	// ************************************************************************************************
	// Now we can proceed.
	// ************************************************************************************************

	var responseData map[string]any
	err = json.Unmarshal(obj.Content, &responseData)
	if err != nil {
		return ErrorResponsef(err, http.StatusInternalServerError, "failed to unmarshal object content: %w")
	}

	// Handle partial field selection
	fieldsParam := req.QueryParams.Get("fields")
	if fieldsParam != "" {
		var fields []string
		if fieldsParam == "none" {
			fields = []string{"id", "href", "lastUpdate", "version"}
		} else {
			fields = strings.Split(fieldsParam, ",")
		}

		// Create a set of fields for quick lookup
		fieldSet := make(map[string]bool)
		for _, f := range fields {
			fieldSet[strings.TrimSpace(f)] = true
		}

		// Always include id, href, lastUpdate, version and @type
		fieldSet["id"] = true
		fieldSet["href"] = true
		fieldSet["lastUpdate"] = true
		fieldSet["version"] = true
		fieldSet["@type"] = true

		filteredItem := make(map[string]any)
		for key, value := range responseData {
			if fieldSet[key] {
				filteredItem[key] = value
			}
		}
		responseData = filteredItem
	}

	slog.Info("Object retrieved successfully", slog.String("id", req.ID), slog.String("resourceName", req.ResourceName))
	return &Response{StatusCode: http.StatusOK, Body: responseData}
}

// UpdateGenericObject updates an existing TMF object using generalized parameters.
func (svc *Service) UpdateGenericObject(req *Request) *Response {
	var err error
	slog.Debug("UpdateGenericObject called", slog.String("id", req.ID), slog.String("resourceName", req.ResourceName))

	// ************************************************************************************************
	// Authentication: we require the user to be authenticated
	// ************************************************************************************************

	var token map[string]any
	{
		// Process the AccessToken to extract caller info from its claims in the payload
		if token, err = svc.processAccessToken(req); err != nil {
			return ErrorResponsef(err, http.StatusUnauthorized, "invalid access token: %w")
		}

		// This operation can not be done without authentication
		if len(token) == 0 {
			return ErrorResponsef(nil, http.StatusUnauthorized, "user not authenticated")
		}

	}

	// ************************************************************************************************
	// Parse the request body, which contains the TMForum object being updated.
	// We perform some formal verifications and add default values if needed.
	// ************************************************************************************************

	var incomingObjMap map[string]any
	var incomingVersion string
	var lastUpdate string
	var sellerDid string
	var sellerOperatorDid string
	{
		// Parse the request body, which contains the TMForum object being created
		var incomingObjMap map[string]any
		if err := json.Unmarshal(req.Body, &incomingObjMap); err != nil {
			return ErrorResponsef(err, http.StatusBadRequest, "failed to bind request body: %w")
		}

		// An update operation specifies the ID of the object to update in the URL.
		// To facilitate life to applications, we allow the ID to be also in the body.
		// But if the ID is present in the body, ensure it matches the ID in the URL
		id, _ := incomingObjMap["id"].(string)
		if id != "" && id != req.ID {
			return ErrorResponsef(nil, http.StatusBadRequest, "ID in body must match ID in URL")
		}

		// version must be specified for update operations
		incomingVersion, _ = incomingObjMap["version"].(string)
		if incomingVersion == "" {
			return ErrorResponsef(nil, http.StatusBadRequest, "version field is required for update operations")
		}
		// TODO: should we do the same for href?

		// Check and process '@type' field
		if typeVal, typeOk := incomingObjMap["@type"].(string); typeOk {
			if !strings.EqualFold(typeVal, req.ResourceName) {
				return ErrorResponsef(nil, http.StatusBadRequest, "@type field in body must match resource name in URL")
			}
		} else {
			// If @type is not specified, add it
			incomingObjMap["@type"] = req.ResourceName
			slog.Debug("Added missing @type field to update request", slog.String("type", req.ResourceName))
		}

		// Set the lastUpdate property. We overwrite whatever the user set.
		now := time.Now()
		lastUpdate := now.Format(time.RFC3339Nano)
		incomingObjMap["lastUpdate"] = lastUpdate

		// Add Seller and Buyer info. We overwrite whatever is in the incoming object, if any
		err = setSellerAndBuyerInfo(incomingObjMap, req.AuthUser.OrganizationIdentifier, req.APIVersion)
		if err != nil {
			return ErrorResponsef(err, http.StatusInternalServerError, "failed to add Seller and Buyer info: %w")
		}

		sellerDid, sellerOperatorDid, err = getSellerAndBuyerInfo(incomingObjMap, req.APIVersion)

	}

	// ************************************************************************************************
	// Retrieve existing object from database to preserve CreatedAt
	// ************************************************************************************************

	var existingObj *repo.TMFObject
	var existingSellerDid string
	var existingSellerOperatorDid string
	{
		// Retrieve existing object
		existingObj, err = svc.getObject(req.ID, req.ResourceName)
		if err != nil {
			return ErrorResponsef(err, http.StatusInternalServerError, "failed to get existing object for update: %w")
		}

		if existingObj == nil {
			return ErrorResponsef(nil, http.StatusNotFound, "object not found")
		}

		// incomingVersion must be lexicographically greater than existingVersion
		if incomingVersion <= existingObj.Version {
			return ErrorResponsef(nil, http.StatusBadRequest, "incoming version must be greater than existing version")
		}

		existingSellerDid, existingSellerOperatorDid, err = getSellerAndBuyerInfo(existingObj.ToMap(), req.APIVersion)
		if err != nil {
			return ErrorResponsef(err, http.StatusInternalServerError, "failed to get seller and buyer info: %w")
		}

	}

	// ************************************************************************************************
	// Before performing the action, check if the user can perform the operation on the object,
	// based on the rules defined by the user in the policy engine.
	// ************************************************************************************************

	// Seller and seller operator must match
	if existingSellerDid != sellerDid || existingSellerOperatorDid != sellerOperatorDid {
		return ErrorResponsef(nil, http.StatusForbidden, "seller or seller operator mismatch")
	}

	// ************************************************************************************************
	// Now we can proceed.
	// ************************************************************************************************

	// Merge incomingObjMap into existing object using RFC7396 (JSON Merge Patch)
	var existingMap map[string]any
	if err := json.Unmarshal(existingObj.Content, &existingMap); err != nil {
		return ErrorResponsef(err, http.StatusInternalServerError, "failed to unmarshal existing object content for merge: %w")
	}

	// RFC7396 merge implementation: modify target in place
	mergeRFC7396(existingMap, incomingObjMap)

	// update incomingObjMap to the merged result so response/notification contains the final content
	incomingObjMap = existingMap

	incomingContent, err := json.Marshal(incomingObjMap)
	if err != nil {
		return ErrorResponsef(err, http.StatusInternalServerError, "failed to marshal object content for update: %w")
	}

	obj := &repo.TMFObject{
		ID:         req.ID,
		Type:       req.ResourceName,
		Version:    incomingVersion,
		APIVersion: req.APIVersion,
		LastUpdate: lastUpdate,
		Content:    incomingContent,
		CreatedAt:  existingObj.CreatedAt,
		UpdatedAt:  time.Now(),
	}

	if err := svc.updateObject(obj); err != nil {
		return ErrorResponsef(err, http.StatusInternalServerError, "failed to update object in service: %w")
	}

	slog.Info("Object updated successfully", slog.String("id", req.ID), slog.String("resourceName", req.ResourceName))

	// Send TMForum notification (AttributeValueChangeEvent)
	eventType := toEventType(req.ResourceName, "AttributeValueChangeEvent")
	eventPayload := buildEventPayload(req, eventType, incomingObjMap)
	svc.notif.PublishEvent(req.APIfamily, eventType, eventPayload)

	return &Response{StatusCode: http.StatusOK, Body: incomingObjMap}
}

// DeleteGenericObject deletes a TMF object using generalized parameters.
func (svc *Service) DeleteGenericObject(req *Request) *Response {
	slog.Debug("DeleteGenericObject called", slog.String("id", req.ID), slog.String("resourceName", req.ResourceName))

	// Authentication: process the AccessToken to extract caller info from its claims in the payload
	token, err := svc.processAccessToken(req)
	if err != nil {
		return ErrorResponsef(err, http.StatusUnauthorized, "invalid access token: %w")
	}

	// Deleting an object can not be done without authentication
	if len(token) == 0 {
		return ErrorResponsef(nil, http.StatusUnauthorized, "user not authenticated")
	}

	// ************************************************************************************************
	// Retrieve existing object from database
	// ************************************************************************************************

	var existingObj *repo.TMFObject
	var existingSellerDid string
	var existingSellerOperatorDid string
	{
		// Retrieve existing object
		existingObj, err = svc.getObject(req.ID, req.ResourceName)
		if err != nil {
			return ErrorResponsef(err, http.StatusInternalServerError, "failed to get existing object for update: %w")
		}

		if existingObj == nil {
			return ErrorResponsef(nil, http.StatusNotFound, "object not found")
		}

		existingSellerDid, existingSellerOperatorDid, err = getSellerAndBuyerInfo(existingObj.ToMap(), req.APIVersion)
		if err != nil {
			return ErrorResponsef(err, http.StatusInternalServerError, "failed to get seller and buyer info: %w")
		}

	}

	// The user must be either the Seller or the SellerOperator to delete the object
	if existingSellerDid != req.AuthUser.OrganizationIdentifier && existingSellerOperatorDid != req.AuthUser.OrganizationIdentifier {
		return ErrorResponsef(nil, http.StatusForbidden, "user must be either the Seller or the SellerOperator")
	}

	// Delete the object in the database
	if err := svc.deleteObject(req.ID, req.ResourceName); err != nil {
		return ErrorResponsef(err, http.StatusInternalServerError, "failed to delete object from service: %w")
	}

	slog.Info("Object deleted successfully", slog.String("id", req.ID), slog.String("resourceName", req.ResourceName))

	// Send TMForum notification
	eventType := toEventType(req.ResourceName, "DeleteEvent")
	minimal := map[string]any{
		"id":    req.ID,
		"@type": req.ResourceName,
		"href":  fmt.Sprintf("/tmf-api/%s/%s/%s/%s", req.APIfamily, req.APIVersion, req.ResourceName, req.ID),
	}
	eventPayload := buildEventPayload(req, eventType, minimal)
	svc.notif.PublishEvent(req.APIfamily, eventType, eventPayload)

	return &Response{StatusCode: http.StatusNoContent}
}

// ListGenericObjects retrieves all TMF objects of a given type using generalized parameters.
func (svc *Service) ListGenericObjects(req *Request) *Response {
	slog.Debug("ListGenericObjects called", slog.String("resourceName", req.ResourceName))

	// Authentication: process the AccessToken to extract caller info from its claims in the payload
	_, err := svc.processAccessToken(req)
	if err != nil {
		return ErrorResponsef(err, http.StatusUnauthorized, "invalid access token: %w")
	}

	objs, totalCount, err := svc.listObjects(req.ResourceName, req.QueryParams)
	if err != nil {
		return ErrorResponsef(err, http.StatusInternalServerError, "failed to list objects from service: %w")
	}

	headers := make(map[string]string)
	headers["X-Total-Count"] = strconv.Itoa(totalCount)

	// var responseData []map[string]any

	responseData := make([]map[string]any, 0)

	for _, obj := range objs {
		var item map[string]any
		err := json.Unmarshal(obj.Content, &item)
		if err != nil {
			return ErrorResponsef(err, http.StatusInternalServerError, "failed to unmarshal object content for listing: %w")
		}
		responseData = append(responseData, item)
	}

	// Handle partial field selection
	fieldsParam := req.QueryParams.Get("fields")
	if fieldsParam != "" {
		var fields []string
		if fieldsParam == "none" {
			fields = []string{"id", "href", "lastUpdate", "version"}
		} else {
			fields = strings.Split(fieldsParam, ",")
		}

		// Create a set of fields for quick lookup
		fieldSet := make(map[string]bool)
		for _, f := range fields {
			fieldSet[strings.TrimSpace(f)] = true
		}

		// Always include id, href, lastUpdate, version and @type
		fieldSet["id"] = true
		fieldSet["href"] = true
		fieldSet["lastUpdate"] = true
		fieldSet["version"] = true
		fieldSet["@type"] = true

		var filteredResponseData []map[string]any
		for _, item := range responseData {
			filteredItem := make(map[string]any)
			for key, value := range item {
				if fieldSet[key] {
					filteredItem[key] = value
				}
			}
			filteredResponseData = append(filteredResponseData, filteredItem)
		}
		responseData = filteredResponseData
	}

	slog.Info("Objects listed successfully", slog.Int("count", len(responseData)), slog.String("resourceName", req.ResourceName))
	return &Response{StatusCode: http.StatusOK, Headers: headers, Body: responseData}
}

// mergeRFC7396 implements JSON Merge Patch (RFC 7396) to merge a patch object into a target object.
// The function modifies the target object in place according to RFC 7396 rules:
// - If the patch value is null, the member is removed from the target
// - If both values are objects, they are merged recursively
// - For arrays or scalar values, the target value is replaced
func mergeRFC7396(target, patch map[string]any) {
	for k, v := range patch {
		// If the patch value is nil -> remove the member from the target
		if v == nil {
			delete(target, k)
			continue
		}

		// If both are objects, merge recursively
		vMap, vIsMap := v.(map[string]any)
		if vIsMap {
			if existingChild, ok := target[k]; ok {
				if existingChildMap, ok2 := existingChild.(map[string]any); ok2 {
					mergeRFC7396(existingChildMap, vMap)
					target[k] = existingChildMap
					continue
				}
			}
			// Otherwise, replace with the incoming object
			target[k] = vMap
			continue
		}

		// For arrays or scalar values, replace
		target[k] = v
	}
}

// ToKebabCase converts a camelCase string to kebab-case.
// For example: "productOffering" becomes "product-offering".
func ToKebabCase(s string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(s, "${1}-${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}-${2}")
	return strings.ToLower(snake)
}

// toEventType composes the TMF event type name from a resource name and a suffix.
func toEventType(resourceName, suffix string) string {
	// Convert resourceName to PascalCase
	parts := strings.Split(resourceName, "-")
	if len(parts) > 1 {
		for i, p := range parts {
			if len(p) == 0 {
				continue
			}
			parts[i] = strings.ToUpper(p[:1]) + strings.ToLower(p[1:])
		}
		return strings.Join(parts, "") + suffix
	}
	if len(resourceName) == 0 {
		return suffix
	}
	return strings.ToUpper(resourceName[:1]) + resourceName[1:] + suffix
}

// buildEventPayload builds a generic TMF event envelope.
func buildEventPayload(req *Request, eventType string, resource any) map[string]any {
	return map[string]any{
		"eventId":      uuid.NewString(),
		"eventTime":    time.Now().Format(time.RFC3339Nano),
		"eventType":    eventType,
		"apiFamily":    req.APIfamily,
		"resourceName": req.ResourceName,
		"resourceId":   req.ID,
		"resourcePath": fmt.Sprintf("/tmf-api/%s/%s/%s", req.APIfamily, req.APIVersion, req.ResourceName),
		"event": map[string]any{
			"resource": resource,
		},
	}
}

// getHTTPStatusInfo returns the standard HTTP status code, code string, and status message
func getHTTPStatusInfo(statusCode int) (int, string, string) {
	switch statusCode {
	case http.StatusBadRequest:
		return statusCode, "400", "Bad Request"
	case http.StatusUnauthorized:
		return statusCode, "401", "Unauthorized"
	case http.StatusForbidden:
		return statusCode, "403", "Forbidden"
	case http.StatusNotFound:
		return statusCode, "404", "Not Found"
	case http.StatusMethodNotAllowed:
		return statusCode, "405", "Method Not Allowed"
	case http.StatusConflict:
		return statusCode, "409", "Conflict"
	case http.StatusInternalServerError:
		return statusCode, "500", "Internal Server Error"
	case http.StatusNotImplemented:
		return statusCode, "501", "Not Implemented"
	case http.StatusBadGateway:
		return statusCode, "502", "Bad Gateway"
	case http.StatusServiceUnavailable:
		return statusCode, "503", "Service Unavailable"
	default:
		// For any other status code, use the numeric value as both code and status
		return statusCode, fmt.Sprintf("%d", statusCode), http.StatusText(statusCode)
	}
}

// ErrorResponse creates a standardized error response using only the HTTP status code
func ErrorResponse(err error, statusCode int) *Response {
	_, code, status := getHTTPStatusInfo(statusCode)
	wrappedErr := errl.Error2(err)
	apiErr := NewApiError(code, status, wrappedErr.Error(), fmt.Sprintf("%d", statusCode), "")
	return &Response{StatusCode: statusCode, Body: apiErr}
}

// ErrorResponsef creates a standardized error response with formatted error message using only the HTTP status code
func ErrorResponsef(err error, statusCode int, format string, args ...any) *Response {
	_, code, status := getHTTPStatusInfo(statusCode)
	wrappedErr := errl.Errorf2(format, append(args, err)...)
	apiErr := NewApiError(code, status, wrappedErr.Error(), fmt.Sprintf("%d", statusCode), "")
	return &Response{StatusCode: statusCode, Body: apiErr}
}
