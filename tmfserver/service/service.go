package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
	"github.com/hesusruiz/isbetmf/tmfclient"
	"github.com/hesusruiz/isbetmf/tmfserver/notifications"
	"github.com/hesusruiz/isbetmf/tmfserver/repository"
	repo "github.com/hesusruiz/isbetmf/tmfserver/repository"
	"github.com/hesusruiz/isbetmf/types"
	"github.com/jmoiron/sqlx"
)

// The DOME implementation has some non-conformances with the TMF specifications, and this is to bypass them.
var DOMEHacks = true

// Request represents a generic HTTP request. Handlers must convert to this representation.
// In this way, we support easily any HTTP framework (currently Fiber), but also other
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
	AuthUser     types.AuthUser
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

	// Pluggable storage backend. When nil, falls back to built-in SQLite via db
	storage TMFStorage

	// The rules engine implemented using Starlark
	ruleEngine *pdp.PDP

	// The Verifier server which signs the Access Tokens,
	// and the PDP retrieves the JWKS from it to verify the signatures.
	verifierServer string

	// The OpenID configuration of the Verifier Server
	oid *OpenIDConfig

	// Notifications manager
	notif *notifications.Manager

	// TMF Client for proxying requests
	tmfClient *tmfclient.Client

	// Fressness for local objects when proxy enabled
	fressness time.Duration

	// Flag to enable/disable proxy functionality
	proxyEnabled bool

	// The paging service to help process remote TMForum objects
	paging *ClientWithPaging

	// Information about us (the server operator)
	ServerOperatorOrganizationIdentifier string
	ServerOperatorDid                    string
	ServerOperatorName                   string
	ServerOperatorCountry                string
}

// NewService creates a new service.
func NewService(cnf *config.Config, db *sqlx.DB, ruleEngine *pdp.PDP) *Service {
	svc := &Service{
		db:             db,
		ruleEngine:     ruleEngine,
		verifierServer: cnf.VerifierServer,
		proxyEnabled:   cnf.ProxyEnabled,
	}

	svc.ServerOperatorOrganizationIdentifier = cnf.ServerOperatorOrganizationIdentifier
	svc.ServerOperatorDid = cnf.ServerOperatorDid
	svc.ServerOperatorName = cnf.ServerOperatorName
	svc.ServerOperatorCountry = cnf.ServerOperatorCountry

	if svc.proxyEnabled {
		clientCfg := &tmfclient.Config{
			BaseURL: cnf.RemoteTMFServer,
			Timeout: 20,
		}

		svc.tmfClient = tmfclient.NewClient(clientCfg)
		svc.fressness = 10 * time.Minute
	}

	err := svc.initializeService()
	if err != nil {
		panic(err)
	}

	// Create the paging service
	pagingConfig := DefaultPagingConfig()
	pagingConfig.PageSize = 10
	paging := NewClientWithPaging(pagingConfig)
	svc.paging = paging

	// // Test the initial paging
	// counter := 0
	// _, err = svc.paging.GetAllObjectsOfType(context.Background(), "productOffering", func(objType string, obj repo.TMFObjectMap) (repo.TMFObjectMap, bool, error) {
	// 	counter++
	// 	fmt.Println("==>", counter, obj.GetID())
	// 	return obj, true, nil
	// })

	// Initialize notifications with in-memory store and HTTP delivery
	store := notifications.NewMemoryStore()
	deliver := notifications.NewHTTPDelivery(5 * time.Second)
	svc.notif = notifications.NewManager(store, deliver)

	return svc
}

func (svc *Service) initializeService() error {

	// Create the server operator identity, in case it is not yet in the database
	org := &repository.Organization{
		CommonName:             svc.ServerOperatorName,
		Country:                svc.ServerOperatorCountry,
		EmailAddress:           "",
		Organization:           svc.ServerOperatorName,
		OrganizationIdentifier: svc.ServerOperatorOrganizationIdentifier,
	}

	obj, _ := repository.TMFOrganizationFromToken(nil, org)

	if err := svc.upsertObject(obj); err != nil {
		if errors.Is(err, &ErrObjectExists{}) {
			slog.Debug("server operator organization already exists", "organizationIdentifier", svc.ServerOperatorOrganizationIdentifier)
		} else {
			err = errl.Errorf("error creating server operator organization: %w", err)
			panic(err)
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
		return ErrorResponsef(http.StatusUnauthorized, "invalid access token: %w", err)
	}

	// Parse incoming body
	var body map[string]any
	if err := json.Unmarshal(req.Body, &body); err != nil {
		return ErrorResponsef(http.StatusBadRequest, "failed to bind request body: %w", err)
	}

	callback, _ := body["callback"].(string)
	if callback == "" {
		return ErrorResponsef(http.StatusBadRequest, "callback is required")
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
		return ErrorResponsef(http.StatusInternalServerError, "failed to create subscription: %w", err)
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
		return ErrorResponsef(http.StatusUnauthorized, "invalid access token: %w", err)
	}

	if req.ID == "" {
		return ErrorResponsef(http.StatusBadRequest, "id is required")
	}

	if err := svc.notif.DeleteSubscription(req.APIfamily, req.ID); err != nil {
		return ErrorResponsef(http.StatusNotFound, "failed to delete subscription: %w", err)
	}

	return &Response{StatusCode: http.StatusNoContent}
}

// CreateGenericObject creates a new TMF object using generalized parameters.
//
// The function performs the following checks and operations:
//
// 1.  **Authentication**:
//   - It requires an authenticated user, obtained by processing the Access Token.
//   - If the user is not authenticated, it returns a 401 Unauthorized error.
//
// 2.  **Incoming TMF Object Checks**:
//   - It parses the request body to get the incoming TMF object.
//   - It checks for mandatory name fields based on the resource type (`givenName` and `familyName` for `Individual`, `tradingName` for `Organization`, and `name` for others).
//   - If an `id` is provided in the incoming object, it ensures that a `version` is also present.
//   - If no `id` is provided, a new one is generated in the format "urn:ngsi-ld:{resource-in-kebab-case}:{uuid}".
//   - It sets a default `version` to "1.0" if not provided.
//   - It overwrites the `href` field to ensure it is correct.
//   - It adds a default `lifecycleStatus` if the resource type requires it and it's not present.
//   - It sets the `@type` field if it's not specified.
//   - It sets the `seller` and `buyer` information, overwriting any existing values.
//
// 3.  **Authorization**:
//   - It calls the Policy Decision Point (PDP) to check if the user is authorized to create the object.
//   - If the user is not authorized, it returns a 403 Forbidden error.
//
// 4.  **Object Creation**:
//   - If proxy mode is enabled, it forwards the request to the remote TMF server.
//   - It checks the remote server's response. If the remote server does not provide a `lastUpdate` field, it adds one and logs a warning.
//   - It stores the object returned by the remote server in the local database.
//   - If proxy mode is disabled, it creates the object directly in the local database.
//   - It sets the `lastUpdate` field to the current time.
//
// 5.  **Response and Notification**:
//   - It returns a 201 Created response with the created object in the body.
//   - It sends a "CreateEvent" notification to subscribed listeners.
func (svc *Service) CreateGenericObject(req *Request) *Response {
	var err error
	slog.Debug("CreateGenericObject called", slog.String("apiFamily", req.APIfamily), slog.String("resourceName", req.ResourceName))

	var tokenMap map[string]any

	// ************************************************************************************************
	// Authentication: we require the user to be authenticated
	// ************************************************************************************************
	{
		// Process the AccessToken to extract caller info from its claims in the payload
		if tokenMap, err = svc.processAccessToken(req); err != nil {
			return ErrorResponsef(http.StatusUnauthorized, "invalid access token: %w", err)
		}

		// This operation can not be done without authentication
		if len(tokenMap) == 0 {
			return ErrorResponsef(http.StatusUnauthorized, "user not authenticated")
		}

	}

	var incomingObjectMap repo.TMFObjectMap
	var id string
	var version string

	// ************************************************************************************************
	// Parse the request body, which contains the TMForum object being created.
	// We perform some formal verifications and add default values if needed:
	// - Check for compulsory 'name' field (or equivalent fields for Individual or Organization)
	// - If the incoming object specifies an 'id', this is only possible if it creates a new version.
	//     That means the incoming object must include a 'version' property.
	// - Create a new 'id' if the user did not specify it
	// - Overwrite href in case it is specified by the caller.
	// - Set a proper lifecycleStatus, if it was not specified by the caller
	// - Set the '@type' based on the 'resource' name, if not specified by the user
	// - Add Seller and Buyer info. We overwrite whatever is in the incoming object.
	//     At this moment we set the Seller to the caller, and SellerOperator to us (the server operator).
	//     This is good when we are a single server, but if we replicate with other marketplaces, we need
	//     to have a trusted list of marketplaces. It the request comes from one of them, we trust it for
	//     the content of the Seller object, as long as the SellerOperator field is one of the marketplaces.
	//     In other words: if the organizationIdentifier of the caller is the same as the SellerOperator atribute,
	//     and it is different from our identifier, then we do not check the Seller attribute.
	// TODO: revise the Seller and SellerOperator logic when the trusted list of marketplaces is implemented
	// ************************************************************************************************
	{

		incomingObjectMap, err = repo.NewTMFObjectMapFromRequest(req.ResourceName, req.Body)
		if err != nil {
			return ErrorResponsef(http.StatusBadRequest, "failed to bind request body: %w", errl.Error(err))
		}

		// Check for the different required name fields depending on the object type
		if strings.EqualFold(req.ResourceName, "Individual") {
			// Individual requires givenName and familyName
			if givenName, _ := incomingObjectMap["givenName"].(string); givenName == "" {
				return ErrorResponsef(http.StatusBadRequest, "givenName is required in individual object: %s", incomingObjectMap)
			}
			if familyName, _ := incomingObjectMap["familyName"].(string); familyName == "" {
				return ErrorResponsef(http.StatusBadRequest, "familyName is required in individual object: %s", incomingObjectMap)
			}
		} else if strings.EqualFold(req.ResourceName, "Organization") {
			// Organization requires tradingName
			if tradingName, _ := incomingObjectMap["tradingName"].(string); tradingName == "" {
				return ErrorResponsef(http.StatusBadRequest, "tradingName is required in organization object: %s", incomingObjectMap)
			}
		} else if name, _ := incomingObjectMap["name"].(string); name == "" {
			// Other objects require name
			return ErrorResponsef(http.StatusBadRequest, "name is required: %s", incomingObjectMap)
		}

		id = incomingObjectMap.GetID()
		version = incomingObjectMap.GetVersion()

		// If the incoming object specifies an 'id', this is only possible if it creates a new version.
		// That means the incoming object must have a 'version' property.
		if id != "" && version == "" {
			return ErrorResponsef(http.StatusBadRequest, "id specified but version is missing: %s", incomingObjectMap)
		}

		// Create a new 'id' if the user did not specify it
		if id == "" {
			// DOME does not support using an 'id' in the incoming object.
			if !svc.proxyEnabled || !DOMEHacks {
				// If the incoming object does not have an 'id', we generate a new one
				// The format is "urn:ngsi-ld:{resource-in-kebab-case}:{uuid}"
				id = fmt.Sprintf("urn:ngsi-ld:%s:%s:%s", ToKebabCase(req.ResourceName), req.AuthUser.OrganizationIdentifier, uuid.NewString())
				incomingObjectMap.SetID(id)
				slog.Debug("Generated new ID for object", "id", id)
			}

			// Set default 'version' if not provided by the user
			if version == "" {
				version = "1.0"
				incomingObjectMap.SetVersion(version)
				slog.Debug("Set default version", slog.String("version", version))
			}

		}

		// Overwrite href in case it is specified by the caller.
		// We could be more strict and reject the request, but this is more permissive.
		if !svc.proxyEnabled || !DOMEHacks {
			incomingObjectMap.SetHref(fmt.Sprintf("/tmf-api/%s/%s/%s/%s", req.APIfamily, req.APIVersion, req.ResourceName, id))
			slog.Debug("Set href", slog.String("href", incomingObjectMap["href"].(string)))
		}

		if incomingObjectMap.GetLastUpdate() != "" {
			// Set the lastUpdate property. We overwrite whatever the user sets.
			incomingObjectMap.SetLastUpdate(time.Now().Format(time.RFC3339Nano))
		}

		// If the object requires a lifecycleStatus, add it if not specified by the caller
		if baseStatus, ok := LifecycleStatusMandatory[req.ResourceName]; ok {
			if lifecycleStatus := incomingObjectMap.GetLifecycleStatus(); lifecycleStatus == "" {
				slog.Debug("adding lifecycleStatus", slog.String("status", baseStatus))
				incomingObjectMap.SetLifecycleStatus(baseStatus)
			}
		}

		// Set the @type if not specified by the user
		if resourceType := incomingObjectMap.GetType(); resourceType == "" {
			resourceType = strings.ToUpper(req.ResourceName[0:1]) + req.ResourceName[1:]
			incomingObjectMap.SetType(resourceType)
		}

		// Add Seller info. We overwrite whatever is in the incoming object
		err = incomingObjectMap.SetSellerInfo(svc.ServerOperatorDid, req.AuthUser.OrganizationIdentifier, req.APIVersion)
		if err != nil {
			return ErrorResponsef(http.StatusInternalServerError, "failed to add Seller and Buyer info: %w", err)
		}

	}

	// ************************************************************************************************
	// Before performing the action, check if the user can perform the operation on the object,
	// based on the rules defined by the user in the policy engine.
	// ************************************************************************************************

	if authorized, err := svc.takeDecision(svc.ruleEngine, req, tokenMap, incomingObjectMap); !authorized {
		return ErrorResponsef(http.StatusForbidden,
			"user %s is not authorized, object: %s, error: %w",
			req.AuthUser.OrganizationIdentifier,
			incomingObjectMap,
			err,
		)
	} else {
		slog.Debug("caller is authorized to create object", "reason", err)
	}

	// ************************************************************************************************
	// Now we can proceed, creating an object in the database or the remote server in proxy mode
	// ************************************************************************************************

	respSt := svc.createLocalOrRemoteObject(req, incomingObjectMap)

	if respSt.StatusCode == http.StatusCreated {
		// Send TMForum notification
		eventType := toEventType(req.ResourceName, "CreateEvent")
		eventPayload := buildEventPayload(req, eventType, respSt.Body)
		svc.notif.PublishEvent(req.APIfamily, eventType, eventPayload)
	}

	// We respond with the created object with all the properties, including the default ones.
	// TODO: not critical, but we could send only the mandatory attributes and those requested by the caller.
	// Sending the whole object is compliant with the TMForum specification,
	// but sending only the mandatory attributes is more efficient.
	return respSt
}

// GetGenericObject retrieves a TMF object using generalized parameters.
//
// The function performs the following checks and operations:
//
// 1.  **Authentication**:
//   - It processes the Access Token to get the caller's information. Public (unauthenticated) access is allowed.
//
// 2.  **Object Retrieval**:
//   - It first tries to retrieve the object from the local database.
//   - If the object is not found and proxy mode is enabled, it tries to fetch it from the remote TMF server.
//   - If the object is found remotely, it is cached in the local database.
//   - If the object is not found locally or remotely, it returns a 404 Not Found error.
//
// 3.  **Authorization**:
//   - It calls the Policy Decision Point (PDP) to check if the user is authorized to read the object.
//   - If the user is not authorized, it returns a 403 Forbidden error.
//
// 4.  **Response**:
//   - It handles partial field selection using the "fields" query parameter.
//   - It returns a 200 OK response with the (potentially filtered) object in the body.
func (svc *Service) GetGenericObject(req *Request) *Response {
	slog.Debug("GetGenericObject called", slog.String("id", req.ID), slog.String("resourceName", req.ResourceName))

	// Authentication: process the AccessToken to extract caller info from its claims in the payload
	// Anonymous access is allowed, so the tokenMap may be nil
	tokenMap, err := svc.processAccessToken(req)
	if err != nil {
		return ErrorResponsef(http.StatusUnauthorized, "invalid access token: %w", errl.Error(err))
	}

	// Retrieve the object from the database. If it is not found, we try to get it from the remote server (if proxy is enabled).
	// If the object is not found, we return a 404 error.
	existingObject, err := svc.getLocalOrRemoteObject(req)
	if err != nil {
		// This is an unexpected error, so we return a server error
		return ErrorResponsef(http.StatusInternalServerError, "failed to get object from service: %w", errl.Error(err))
	}

	if existingObject == nil {
		return ErrorResponsef(http.StatusNotFound, "object not found")
	}

	existingObjectMap, err := existingObject.ToMap()
	if err != nil {
		return ErrorResponsef(http.StatusInternalServerError, "failed to unmarshal existing object content: %w", errl.Error(err))
	}

	// ************************************************************************************************
	// Before performing the action, check if the user can perform the operation on the object,
	// based on the rules defined by the user in the policy engine.
	// ************************************************************************************************

	if authorized, err := svc.takeDecision(svc.ruleEngine, req, tokenMap, existingObjectMap); !authorized {
		return ErrorResponsef(http.StatusForbidden,
			"user %s is not authorized, object: %s, error: %w",
			req.AuthUser.OrganizationIdentifier,
			existingObjectMap,
			err,
		)
	} else {
		slog.Debug("caller is authorized to read object", "reason", err)
	}

	// ************************************************************************************************
	// Now we can proceed.
	// ************************************************************************************************

	// Handle partial field selection
	fieldsParam := req.QueryParams.Get("fields")
	if fieldsParam != "" {
		var fields []string
		if fieldsParam != "none" {
			fields = strings.Split(fieldsParam, ",")
		}

		// Create a set of fields for quick lookup
		fieldSet := make(map[string]bool)
		for _, f := range fields {
			fieldSet[strings.TrimSpace(f)] = true
		}

		// Always include id, href, lastUpdate, version, @type and lifecycleStatus
		fieldSet["id"] = true
		fieldSet["href"] = true
		fieldSet["lastUpdate"] = true
		fieldSet["version"] = true
		fieldSet["@type"] = true
		fieldSet["lifecycleStatus"] = true

		filteredItem := make(map[string]any)
		for key, value := range existingObjectMap {
			if fieldSet[key] {
				filteredItem[key] = value
			}
		}
		existingObjectMap = filteredItem
	}

	slog.Info("Object retrieved successfully", slog.String("id", req.ID), slog.String("resourceName", req.ResourceName))
	return &Response{StatusCode: http.StatusOK, Body: existingObjectMap}
}

// UpdateGenericObject updates an existing TMF object using generalized parameters.
//
// The function performs the following checks and operations:
//
// 1.  **Authentication**:
//   - It requires an authenticated user, obtained by processing the Access Token.
//   - If the user is not authenticated, it returns a 401 Unauthorized error.
//
// 2.  **Incoming TMF Object Checks**:
//   - It parses the request body to get the incoming TMF object patch.
//   - It ensures that if an `id` is present in the body, it matches the `id` in the URL.
//   - It sets the `seller` and `buyer` information, overwriting any existing values.
//   - It sets the `lastUpdate` field to the current time.
//
// 3.  **Existing TMF Object Checks**:
//   - It retrieves the existing object from the local database or the remote server (if in proxy mode).
//   - If the object is not found, it returns a 404 Not Found error.
//   - It checks that the object is managed by the current server operator by verifying the `sellerOperator` DID.
//   - It checks that the authenticated user is the owner of the object by verifying the `seller` DID.
//   - It checks the `lifecycleStatus` of the existing object. If it is "Launched", "Retired", or "Obsolete", it requires the incoming object to have a higher `version` number.
//
// 4.  **Authorization**:
//   - The ownership and operator checks act as an authorization mechanism.
//
// 5.  **Object Update**:
//   - If proxy mode is enabled, it forwards the PATCH request to the remote TMF server.
//   - The response from the remote server is then stored in the local database.
//   - If proxy mode is disabled, it merges the incoming patch with the existing object using JSON Merge Patch (RFC 7396).
//   - The updated object is then saved to the local database.
//
// 6.  **Response and Notification**:
//   - It returns a 200 OK response with the updated object in the body.
//   - It sends an "AttributeValueChangeEvent" notification to subscribed listeners.
func (svc *Service) UpdateGenericObject(req *Request) *Response {
	var err error
	slog.Debug("UpdateGenericObject called", slog.String("id", req.ID), slog.String("resourceName", req.ResourceName))

	// ************************************************************************************************
	// Authentication: we require the user to be authenticated
	// ************************************************************************************************

	var tokenMap map[string]any
	{
		// Process the AccessToken to extract caller info from its claims in the payload
		if tokenMap, err = svc.processAccessToken(req); err != nil {
			return ErrorResponsef(http.StatusUnauthorized, "invalid access token: %w", errl.Error(err))
		}

		// This operation can not be done without authentication
		if len(tokenMap) == 0 {
			return ErrorResponsef(http.StatusUnauthorized, "user not authenticated")
		}

	}

	var incomingObjMap repo.TMFObjectMap

	// ************************************************************************************************
	// Parse the request body, which contains the TMForum object being updated.
	// We perform some formal verifications and add default values if needed.
	// ************************************************************************************************
	{

		incomingObjMap, err = repo.NewTMFObjectMapFromRequest(req.ResourceName, req.Body)
		if err != nil {
			return ErrorResponsef(http.StatusBadRequest, "failed to bind request body: %w", errl.Error(err))
		}

		// An update operation specifies the ID of the object to update in the URL.
		// To facilitate life to applications, we allow the 'id' to be also in the body.
		// But if the 'id' is present in the body, ensure it matches the 'id' in the URL
		if !DOMEHacks {
			id, _ := incomingObjMap["id"].(string)
			if id != "" && id != req.ID {
				err := errl.Errorf("ID in body must match ID in URL")
				return ErrorResponsef(http.StatusBadRequest,
					"invalid object, request id: %s, id in body: %s, error: %w",
					req.ID,
					id,
					err,
				)
			}
		}

		// TODO: check if we could set some information for the Individual and Organization objects

		// Set the lastUpdate property. We overwrite whatever the user set.
		now := time.Now()
		lastUpdate := now.Format(time.RFC3339Nano)
		incomingObjMap["lastUpdate"] = lastUpdate

	}

	// ************************************************************************************************
	// Retrieve existing object locally or from proxy
	// ************************************************************************************************

	var existingObj *repo.TMFObject

	// Retrieve existing object, locally or remotely
	existingObj, err = svc.getLocalOrRemoteObject(req)
	if err != nil {
		return ErrorResponsef(http.StatusInternalServerError, "failed to get existing object %s for update: %w", req.ID, err)
	}

	if existingObj == nil {
		return ErrorResponsef(http.StatusNotFound, "object %s not found", req.ID)
	}

	existingObjectMap, err := existingObj.ToMap()
	if err != nil {
		return ErrorResponsef(http.StatusInternalServerError, "failed to unmarshal existing object content: %w", err)
	}

	// ************************************************************************************************
	// Before performing the action, check if the user can perform the operation on the object,
	// based on the rules defined by the user in the policy engine.
	// ************************************************************************************************

	if authorized, err := svc.takeDecision(svc.ruleEngine, req, tokenMap, existingObjectMap); !authorized {
		return ErrorResponsef(http.StatusForbidden,
			"user %s is not authorized, object: %s, error: %w",
			req.AuthUser.OrganizationIdentifier,
			existingObjectMap,
			err,
		)
	} else {
		slog.Debug("caller is authorized to update object", "reason", err)
	}

	// ************************************************************************************************
	// Now we can proceed.
	// ************************************************************************************************

	if svc.proxyEnabled {

		// Proxy logic
		headers := map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Bearer " + req.AccessToken,
		}
		path := fmt.Sprintf("/tmf-api/%s/%s/%s/%s", req.APIfamily, req.APIVersion, req.ResourceName, req.ID)
		resp, err := svc.tmfClient.Patch(path, req.Body, headers)
		if err != nil {
			return ErrorResponsef(http.StatusInternalServerError, "failed to proxy request: %w", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return ErrorResponsef(http.StatusInternalServerError, "failed to read response body: %w", err)
		}

		if resp.StatusCode >= 300 {
			return ErrorResponsef(resp.StatusCode, "remote server returned error: %s", string(body))
		}

		var remoteObjectMap map[string]any
		if err := json.Unmarshal(body, &remoteObjectMap); err != nil {
			return ErrorResponsef(http.StatusBadRequest, "failed to bind request body: %w", err)
		}

		// Set the existingObjectMap to the remoteObjectMap, so we store the object as it is in the remote server
		existingObjectMap = remoteObjectMap

	} else {

		// Merge incomingObjMap into existing object using RFC7396 (JSON Merge Patch)

		// RFC7396 merge implementation: modify target in place
		mergeRFC7396(existingObjectMap, incomingObjMap)

	}

	existingVersion, _ := existingObjectMap["version"].(string)
	existingLastUpdate, _ := existingObjectMap["lastUpdate"].(string)

	existingObjectContent, err := json.Marshal(existingObjectMap)
	if err != nil {
		return ErrorResponsef(http.StatusInternalServerError, "failed to marshal object content for update: %w", err)
	}

	existingObject := &repo.TMFObject{
		ID:         req.ID,
		Type:       req.ResourceName,
		Version:    existingVersion,
		APIVersion: req.APIVersion,
		LastUpdate: existingLastUpdate,
		Content:    existingObjectContent,
		CreatedAt:  existingObj.CreatedAt,
		UpdatedAt:  time.Now(),
	}

	if err := svc.updateObject(existingObject); err != nil {
		return ErrorResponsef(http.StatusInternalServerError, "failed to update object in service: %w", err)
	}

	slog.Info("Object updated successfully", slog.String("id", req.ID), slog.String("resourceName", req.ResourceName))

	// Send TMForum notification (AttributeValueChangeEvent)
	eventType := toEventType(req.ResourceName, "AttributeValueChangeEvent")
	eventPayload := buildEventPayload(req, eventType, existingObjectMap)
	svc.notif.PublishEvent(req.APIfamily, eventType, eventPayload)

	return &Response{StatusCode: http.StatusOK, Body: existingObjectMap}
}

// DeleteGenericObject deletes a TMF object using generalized parameters.
//
// The function performs the following checks and operations:
//
// 1.  **Authentication**:
//   - It requires an authenticated user, obtained by processing the Access Token.
//   - If the user is not authenticated, it returns a 401 Unauthorized error.
//
// 2.  **Existing TMF Object Checks**:
//   - It retrieves the existing object from the local database or the remote server (if in proxy mode).
//   - If the object is not found, it returns a 404 Not Found error.
//   - It checks that the object is managed by the current server operator by verifying the `sellerOperator` DID.
//   - It checks that the authenticated user is the owner of the object by verifying the `seller` DID.
//
// 3.  **Authorization**:
//   - The ownership and operator checks act as an authorization mechanism.
//
// 4.  **Object Deletion**:
//   - If proxy mode is enabled, it forwards the DELETE request to the remote TMF server.
//   - It deletes the object from the local database.
//
// 5.  **Response and Notification**:
//   - It returns a 204 No Content response on successful deletion.
//   - It sends a "DeleteEvent" notification to subscribed listeners.
func (svc *Service) DeleteGenericObject(req *Request) *Response {
	var err error
	slog.Debug("DeleteGenericObject called", slog.String("id", req.ID), slog.String("resourceName", req.ResourceName))

	// ************************************************************************************************
	// Authentication: we require the user to be authenticated
	// ************************************************************************************************

	var tokenMap map[string]any
	{
		// Process the AccessToken to extract caller info from its claims in the payload
		if tokenMap, err = svc.processAccessToken(req); err != nil {
			return ErrorResponsef(http.StatusUnauthorized, "invalid access token: %w", errl.Error(err))
		}

		// This operation can not be done without authentication
		if len(tokenMap) == 0 {
			return ErrorResponsef(http.StatusUnauthorized, "user not authenticated")
		}

	}

	// ************************************************************************************************
	// Retrieve existing object from database
	// ************************************************************************************************

	var existingObject *repo.TMFObject

	// Retrieve existing object, locally or remotely
	existingObject, err = svc.getLocalOrRemoteObject(req)
	if err != nil {
		// TODO: check the return code from remote server and reply accordingly
		return ErrorResponsef(http.StatusBadRequest, "failed to get existing object for update: %w", err)
	}

	// If nothing to delete, return 404
	if existingObject == nil {
		return ErrorResponsef(http.StatusNotFound, "object %s not found", req.ID)
	}

	// Convert to a type-safe map representation to facilitate manipulation
	existingObjectMap, err := existingObject.ToMap()
	if err != nil {
		return ErrorResponsef(http.StatusInternalServerError, "failed to unmarshal existing object content: %w", err)
	}

	// ************************************************************************************************
	// Before performing the action, check if the user can perform the operation on the object,
	// based on the rules defined by the user in the policy engine.
	// ************************************************************************************************

	if authorized, err := svc.takeDecision(svc.ruleEngine, req, tokenMap, existingObjectMap); !authorized {
		return ErrorResponsef(http.StatusForbidden,
			"user %s is not authorized, object: %s, error: %w",
			req.AuthUser.OrganizationIdentifier,
			existingObjectMap,
			err,
		)
	} else {
		slog.Debug("caller is authorized to delete object", "reason", err)
	}

	// ##########################################################
	// Delete the object
	// ##########################################################

	// Delete the object in the remote server, if the proxy is enabled
	if svc.proxyEnabled {
		// Send the authentication header
		headers := map[string]string{
			"Authorization": "Bearer " + req.AccessToken,
		}
		path := fmt.Sprintf("/tmf-api/%s/%s/%s/%s", req.APIfamily, req.APIVersion, req.ResourceName, req.ID)
		resp, err := svc.tmfClient.Delete(path, headers)
		if err != nil {
			return ErrorResponsef(http.StatusInternalServerError, "failed to proxy request: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode >= 300 {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return ErrorResponsef(http.StatusInternalServerError, "failed to read response body: %w", err)
			}
			return &Response{
				StatusCode: resp.StatusCode,
				Body:       body,
			}
		}

		slog.Info("Object deleted from remote server successfully", slog.String("id", req.ID), slog.String("resourceName", req.ResourceName))

	}

	// Delete the object in the local database
	if err := svc.deleteObject(req.ID, req.ResourceName); err != nil {
		return ErrorResponsef(http.StatusInternalServerError, "failed to delete object %s from service: %w", req.ID, err)
	}

	slog.Info("Object deleted successfully from local database", slog.String("id", req.ID), slog.String("resourceName", req.ResourceName))

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
	// Anonymous access is allowed, so the tokenMap may be nil
	tokenMap, err := svc.processAccessToken(req)
	if err != nil {
		return ErrorResponsef(http.StatusUnauthorized, "invalid access token: %w", errl.Error(err))
	}

	// Process the query values: extract 'limit' and 'offset' and leave the rest of query parameters
	var limit = -1
	var offset = -1

	for key, values := range req.QueryParams {

		if key == "limit" {
			limitStr := values[0]
			if limitStr != "" {
				if l, err := strconv.Atoi(limitStr); err == nil {
					limit = l
				}
				req.QueryParams.Del(key)
			}
			continue
		}

		if key == "offset" {
			offsetStr := values[0]
			if offsetStr != "" {
				if l, err := strconv.Atoi(offsetStr); err == nil {
					offset = l
				}
				req.QueryParams.Del(key)
			}
			continue
		}

	}

	if limit == 0 {
		var objs []repo.TMFObject
		return &Response{StatusCode: http.StatusOK, Body: objs}
	}

	if svc.proxyEnabled {

		// Will send the authorization header
		headers := map[string]string{
			"Authorization": "Bearer " + req.AccessToken,
		}

		var objs []repo.TMFObjectMap
		var offsetCounter int
		invalidObjects := 0

		// The page size
		pageSize := 100
		pageOffset := 0

		for {

			// Retrieve the page from the remote server
			path := fmt.Sprintf("/tmf-api/%s/%s/%s?limit=%d&offset=%d", req.APIfamily, req.APIVersion, req.ResourceName, pageSize, pageOffset)
			slog.Debug("sending request to remote", "path", path)

			resp, err := svc.tmfClient.Get(path, headers)
			if err != nil {
				return ErrorResponsef(http.StatusInternalServerError, "failed to proxy request: %w", err)
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return ErrorResponsef(http.StatusInternalServerError, "failed to read response body: %w", err)
			}

			if resp.StatusCode >= 300 {
				return &Response{
					StatusCode: resp.StatusCode,
					Body:       body,
				}
			}

			responseData := make([]map[string]any, 0)
			if err := json.Unmarshal(body, &responseData); err != nil {
				return ErrorResponsef(http.StatusInternalServerError, "failed to unmarshal object content for listing: %w", err)
			}

			// Process the list of response objects to filter depending on the user
			for _, obj := range responseData {
				objectMap := repo.TMFObjectMap(obj)
				validations := objectMap.Validate(req.ResourceName)
				if len(validations.Errors) > 0 {
					invalidObjects++
					fmt.Println(validations.String())
					continue
				}

				theObject := objectMap.ToTMFObject(req.ResourceName)

				// Store locally for further usage
				if err := svc.upsertObject(theObject); err != nil {
					if errors.Is(err, &ErrObjectExists{}) {
						slog.Debug("object already exists", "id", theObject.ID)
					} else {
						invalidObjects++
						err = errl.Errorf("error saving object %s in local database: %w", theObject.ID, err)
						slog.Error("error saving object in local database", "error", err)
						continue
					}
				}

				// Check if the object can be read by the user
				authorized, err := svc.takeDecision(svc.ruleEngine, req, tokenMap, objectMap)
				if !authorized {
					invalidObjects++
					slog.Info("object %s not authorized: %w", objectMap.GetID(), errl.Error(err))
					continue
				}

				// We have to discard the first 'offset' objects, as specified by the user
				if offsetCounter < offset {
					offsetCounter++
					continue
				}

				// Now we add the object to the result array and check if we reached the limit as specified by the user
				objs = append(objs, objectMap)

				if limit >= 0 && len(objs) >= limit {
					break
				}

			}

			// Exit from the outer loop if we already have the number of objects that the user requested
			if limit >= 0 && len(objs) >= limit {
				break
			}

			// Exit the loop if the number of retrieved objects is less than the page size requested.
			// That means that there are no more objects
			if len(responseData) < pageSize {
				break
			}

			// Retrieve next page
			pageOffset += pageSize

		}

		responseHeaders := make(map[string]string)
		responseHeaders["X-Total-Count"] = strconv.Itoa(len(objs))

		slog.Info("Remote objects listed", slog.Int("valid", len(objs)), slog.Int("invalid", invalidObjects), slog.String("resourceName", req.ResourceName))
		return &Response{StatusCode: http.StatusOK, Headers: responseHeaders, Body: objs}

	} else {

		var objs []repo.TMFObject
		var totalCount int

		objs, totalCount, err = svc.listObjects(req, tokenMap, req.ResourceName, req.APIVersion, req.QueryParams)
		if err != nil {
			slog.Error("failed to list objects from service", slog.String("error", err.Error()))
			return ErrorResponsef(http.StatusInternalServerError, "failed to list objects from service: %w", err)
		}

		responseHeaders := make(map[string]string)
		responseHeaders["X-Total-Count"] = strconv.Itoa(totalCount)

		responseData := make([]map[string]any, 0)

		for _, obj := range objs {
			item, err := obj.ToMap()
			if err != nil {
				slog.Error("failed to unmarshal object content for listing", slog.String("error", err.Error()))
				return ErrorResponsef(http.StatusInternalServerError, "failed to unmarshal object content for listing: %w", err)
			}
			responseData = append(responseData, item)
		}

		// Handle partial field selection (filtering)
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
		return &Response{StatusCode: http.StatusOK, Headers: responseHeaders, Body: responseData}
	}

}

func (svc *Service) ListPublicRemoteObjects(req *Request, limit, offset int) ([]repo.TMFObjectMap, error) {

	counter := 0
	objects, err := svc.paging.GetAllObjectsOfType(context.Background(), req.ResourceName, func(objType string, obj repo.TMFObjectMap) (repo.TMFObjectMap, bool, error) {
		counter++
		fmt.Println("==>", counter, obj.GetID())
		return obj, true, nil
	})

	if err != nil {
		return nil, errl.Errorf("retrieving remote objects")
	}

	return objects, nil

}

// func (svc *Service) ListGenericObjectsOld(req *Request) *Response {
// 	slog.Debug("ListGenericObjects called", slog.String("resourceName", req.ResourceName))

// 	// Authentication: process the AccessToken to extract caller info from its claims in the payload
// 	// Anonymous access is allowed, so the tokenMap may be nil
// 	tokenMap, err := svc.processAccessToken(req)
// 	if err != nil {
// 		return ErrorResponsef(http.StatusUnauthorized, "invalid access token: %w", errl.Error(err))
// 	}

// 	// Process the query values: extract 'limit' and 'offset' and leave the rest of query parameters
// 	var limit = -1
// 	var offset = -1

// 	for key, values := range req.QueryParams {

// 		if key == "limit" {
// 			limitStr := values[0]
// 			if limitStr != "" {
// 				if l, err := strconv.Atoi(limitStr); err == nil {
// 					limit = l
// 				}
// 				req.QueryParams.Del(key)
// 			}
// 			continue
// 		}

// 		if key == "offset" {
// 			offsetStr := values[0]
// 			if offsetStr != "" {
// 				if l, err := strconv.Atoi(offsetStr); err == nil {
// 					offset = l
// 				}
// 				req.QueryParams.Del(key)
// 			}
// 			continue
// 		}

// 	}

// 	if limit == 0 {
// 		var objs []repo.TMFObject
// 		return &Response{StatusCode: http.StatusOK, Body: objs}
// 	}

// 	if svc.proxyEnabled {

// 		// Will send the authorization header
// 		headers := map[string]string{
// 			"Authorization": "Bearer " + req.AccessToken,
// 		}

// 		var objs []repo.TMFObjectMap
// 		var offsetCounter int

// 		// The page size
// 		pageSize := 100
// 		pageOffset := 0

// 		for {

// 			// Retrieve the page from the remote server
// 			path := fmt.Sprintf("/tmf-api/%s/%s/%s?limit=%d&offset=%d", req.APIfamily, req.APIVersion, req.ResourceName, pageSize, pageOffset)
// 			slog.Debug("sending request to remote", "path", path)

// 			resp, err := svc.tmfClient.Get(path, headers)
// 			if err != nil {
// 				return ErrorResponsef(http.StatusInternalServerError, "failed to proxy request: %w", err)
// 			}
// 			defer resp.Body.Close()

// 			body, err := io.ReadAll(resp.Body)
// 			if err != nil {
// 				return ErrorResponsef(http.StatusInternalServerError, "failed to read response body: %w", err)
// 			}

// 			if resp.StatusCode >= 300 {
// 				return &Response{
// 					StatusCode: resp.StatusCode,
// 					Body:       body,
// 				}
// 			}

// 			responseData := make([]map[string]any, 0)
// 			if err := json.Unmarshal(body, &responseData); err != nil {
// 				return ErrorResponsef(http.StatusInternalServerError, "failed to unmarshal object content for listing: %w", err)
// 			}

// 			// Process the list of response objects to filter depending on the user
// 			for _, obj := range responseData {
// 				objectMap := repo.TMFObjectMap(obj)

// 				// TODO: Perform verifications

// 				// Check if the object can be read by the user
// 				authorized, err := svc.takeDecision(svc.ruleEngine, req, tokenMap, objectMap)
// 				if !authorized {
// 					slog.Info("object %s not authorized: %w", objectMap.GetID(), errl.Error(err))
// 					continue
// 				}

// 				// We have to discard the first 'offset' objects, as specified by the user
// 				if offsetCounter < offset {
// 					offsetCounter++
// 					continue
// 				}

// 				// Now we add the object to the result array and check if we reached the limit as specified by the user
// 				objs = append(objs, objectMap)

// 				if len(objs) >= limit {
// 					break
// 				}

// 			}

// 			// Exit from the outer loop if we already have the number of objects that the user requested
// 			if len(objs) >= limit {
// 				break
// 			}

// 			// Exit the loop if the number of retrieved objects is less than the page size requested.
// 			// That means that there are no more objects
// 			if len(responseData) < pageSize {
// 				break
// 			}

// 			// Retrieve next page
// 			pageOffset += pageSize

// 		}

// 		responseHeaders := make(map[string]string)
// 		responseHeaders["X-Total-Count"] = strconv.Itoa(len(objs))

// 		slog.Info("Remote objects listed successfully", slog.Int("count", len(objs)), slog.String("resourceName", req.ResourceName))
// 		return &Response{StatusCode: http.StatusOK, Headers: responseHeaders, Body: objs}

// 	} else {

// 		var objs []repo.TMFObject
// 		var totalCount int

// 		objs, totalCount, err = svc.listObjects(req, tokenMap, req.ResourceName, req.APIVersion, req.QueryParams)
// 		if err != nil {
// 			slog.Error("failed to list objects from service", slog.String("error", err.Error()))
// 			return ErrorResponsef(http.StatusInternalServerError, "failed to list objects from service: %w", err)
// 		}

// 		responseHeaders := make(map[string]string)
// 		responseHeaders["X-Total-Count"] = strconv.Itoa(totalCount)

// 		responseData := make([]map[string]any, 0)

// 		for _, obj := range objs {
// 			item, err := obj.ToMap()
// 			if err != nil {
// 				slog.Error("failed to unmarshal object content for listing", slog.String("error", err.Error()))
// 				return ErrorResponsef(http.StatusInternalServerError, "failed to unmarshal object content for listing: %w", err)
// 			}
// 			responseData = append(responseData, item)
// 		}

// 		// Handle partial field selection (filtering)
// 		fieldsParam := req.QueryParams.Get("fields")
// 		if fieldsParam != "" {
// 			var fields []string
// 			if fieldsParam == "none" {
// 				fields = []string{"id", "href", "lastUpdate", "version"}
// 			} else {
// 				fields = strings.Split(fieldsParam, ",")
// 			}

// 			// Create a set of fields for quick lookup
// 			fieldSet := make(map[string]bool)
// 			for _, f := range fields {
// 				fieldSet[strings.TrimSpace(f)] = true
// 			}

// 			// Always include id, href, lastUpdate, version and @type
// 			fieldSet["id"] = true
// 			fieldSet["href"] = true
// 			fieldSet["lastUpdate"] = true
// 			fieldSet["version"] = true
// 			fieldSet["@type"] = true

// 			var filteredResponseData []map[string]any
// 			for _, item := range responseData {
// 				filteredItem := make(map[string]any)
// 				for key, value := range item {
// 					if fieldSet[key] {
// 						filteredItem[key] = value
// 					}
// 				}
// 				filteredResponseData = append(filteredResponseData, filteredItem)
// 			}
// 			responseData = filteredResponseData
// 		}

// 		slog.Info("Objects listed successfully", slog.Int("count", len(responseData)), slog.String("resourceName", req.ResourceName))
// 		return &Response{StatusCode: http.StatusOK, Headers: responseHeaders, Body: responseData}
// 	}

// }

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
	_, code, reason := getHTTPStatusInfo(statusCode)

	message := ""
	if err != nil {
		wrappedErr := errl.Error2(err)
		message = fmt.Sprintf("%s %s: %s", code, reason, wrappedErr.Error())
	} else {
		message = fmt.Sprintf("%s %s", code, reason)
	}

	if statusCode >= 500 {
		slog.Error("server error", slog.String("error", message))
	} else {
		slog.Warn("client error", slog.String("error", message))
	}

	apiErr := NewApiError(code, reason, message, "", "")
	return &Response{StatusCode: statusCode, Body: apiErr}
}

// ErrorResponsef creates a standardized error response with formatted error message using only the HTTP status code
func ErrorResponsef(statusCode int, format string, args ...any) *Response {
	_, code, reason := getHTTPStatusInfo(statusCode)

	wrappedErr := errl.Errorf2(format, args...)
	message := fmt.Sprintf("%s %s: %s", code, reason, wrappedErr.Error())

	if statusCode >= 500 {
		slog.Error("server error", slog.String("error", message))
	} else {
		slog.Warn("client error", slog.String("error", message))
	}

	apiErr := NewApiError(code, reason, message, "", "")
	return &Response{StatusCode: statusCode, Body: apiErr}
}

// LifecycleStatusMandatory defines which TMF resources require the lifecycleStatus field as mandatory.
// The map key is the resource name, and the value is the default lifecycleStatus value to set if missing.
// The keys are in lowercase to facilitate case-independent lookup.
var LifecycleStatusMandatory = map[string]string{
	// TMF620
	"catalog":              "In Study",
	"category":             "In Study",
	"productOffering":      "In Study",
	"productOfferingPrice": "In Study",
	"productSpecification": "In Study",

	// TMF633
	"serviceCandidate":     "In Study",
	"serviceCatalog":       "In Study",
	"serviceCategory":      "In Study",
	"serviceSpecification": "In Study",

	// TMF634
	"LogicalResourceSpecification":  "In Study",
	"PhysicalResourceSpecification": "In Study",
	"ResourceCandidate":             "In Study",
	"ResourceCatalog":               "In Study",
	"ResourceCategory":              "In Study",
	"ResourceFunctionSpecification": "In Study",
	"ResourceSpecification":         "In Study",

	// TMF635
	"UsageSpecification": "In Study",

	// TMF651
	"AgreementSpecification": "In Study",
}

func sameOrganizations(did, orgID string) bool {

	did = strings.TrimPrefix(did, "did:elsi:")
	orgID = strings.TrimPrefix(orgID, "did:elsi:")

	return did == orgID
}
