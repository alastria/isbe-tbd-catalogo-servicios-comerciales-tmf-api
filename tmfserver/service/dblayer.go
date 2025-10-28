package service

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/hesusruiz/isbetmf/internal/errl"
	repo "github.com/hesusruiz/isbetmf/tmfserver/repository"
	"github.com/mattn/go-sqlite3"
	"gitlab.com/greyxor/slogor"
)

// createObject creates a new TMF object.
func (svc *Service) createObject(obj *repo.TMFObject) error {
	slog.Debug("dbLayer: Creating object", slog.String("id", obj.ID), slog.String("type", obj.Type), slog.String("version", obj.Version))

	// Direct to another storage system if defined
	if svc.storage != nil {
		return svc.storage.CreateObject(obj)
	}

	// Make sure timestamps are correct
	now := time.Now()
	obj.CreatedAt = now
	obj.UpdatedAt = now

	// Execute the SQL
	_, err := svc.db.NamedExec(`INSERT INTO tmf_object (id, type, version, api_version, seller, buyer, last_update, content, created_at, updated_at)
		VALUES (:id, :type, :version, :api_version, :seller, :buyer, :last_update, :content, :created_at, :updated_at)`, obj)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if sqliteErr.Code == sqlite3.ErrConstraint && sqliteErr.ExtendedCode == sqlite3.ErrConstraintPrimaryKey {
				return &ErrObjectExists{ID: obj.ID, Type: obj.Type}
			}
		}
		err = errl.Errorf("failed to create object id=%s type=%s: %w", obj.ID, obj.Type, err)
	}
	return err
}

func (svc *Service) createLocalOrRemoteObject(req *Request, objMap repo.TMFObjectMap) *Response {

	// Set the lastUpdate attribute
	objMap.SetLastUpdate(time.Now().Format(time.RFC3339Nano))

	obj := objMap.ToTMFObject()

	if !svc.proxyEnabled {
		// Create the new object in the local database
		if err := svc.createObject(obj); err != nil {
			if errors.Is(err, &ErrObjectExists{}) {
				return ErrorResponsef(http.StatusBadRequest, "object %s already exists: %w", obj.ID, err)
			} else {
				return ErrorResponsef(http.StatusInternalServerError, "failed to create object locally: %w", err)
			}
		}

		return &Response{StatusCode: http.StatusCreated, Body: objMap}
	}

	// Proxy logic
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + req.AccessToken,
	}

	modifiedIncomingBody, err := json.Marshal(objMap)
	if err != nil {
	}
	headers["Content-Length"] = strconv.Itoa(len(modifiedIncomingBody))

	path := fmt.Sprintf("/tmf-api/%s/%s/%s", req.APIfamily, req.APIVersion, req.ResourceName)
	resp, err := svc.tmfClient.Post(path, modifiedIncomingBody, headers)
	if err != nil {
		return ErrorResponsef(http.StatusInternalServerError, "failed to proxy request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ErrorResponsef(http.StatusInternalServerError, "failed to read response body: %w", err)
	}

	if resp.StatusCode >= 300 {
		return ErrorResponsef(resp.StatusCode, "body: %s", string(body))
	}

	// Put the created objet in a map

	remoteObjectMap, err := repo.NewTMFObjectMapFromRequest(req.ResourceName, body)
	if err != nil {
		return ErrorResponsef(http.StatusInternalServerError, "failed to bind request body: %w", errl.Error(err))
	}

	id := remoteObjectMap.GetID()
	version := remoteObjectMap.GetVersion()
	lastUpdate := remoteObjectMap.GetLastUpdate()

	// It is an error if the remote server does not return a 'lastUpdate', but we fix it and log a warning
	if lastUpdate == "" {
		slog.Warn("remote server did not return lastUpdate, fixing it", slog.String("id", id))

		now := time.Now()
		lastUpdate = now.Format(time.RFC3339Nano)
		remoteObjectMap["lastUpdate"] = lastUpdate
	}

	remoteContent, err := json.Marshal(remoteObjectMap)
	if err != nil {
		return ErrorResponsef(http.StatusInternalServerError, "failed to marshal object content: %w", err)
	}

	remoteObject := repo.NewTMFObject(id, req.ResourceName, version, req.APIVersion, lastUpdate, remoteContent)

	// Create the new object in the local database
	if err := svc.createObject(remoteObject); err != nil {
		return ErrorResponsef(http.StatusInternalServerError, "failed to create object in service: %w", err)
	}

	// Set the headers of the reply to the caller, as per TMF specs
	headers = map[string]string{
		"Location": remoteObjectMap.GetHref(),
	}
	slog.Info("Object created successfully", slog.String("id", id), slog.String("resourceName", req.ResourceName), slog.String("location", remoteObjectMap.GetHref()))

	return &Response{StatusCode: http.StatusCreated, Body: remoteObjectMap}

}

// getObject retrieves a TMF object by its ID and type, returning the latest version.
func (svc *Service) getObject(id, objectType string) (*repo.TMFObject, error) {
	slog.Debug("dbLayer: Getting object", slog.String("id", id), slog.String("type", objectType))

	// Direct to another storage system if defined
	if svc.storage != nil {
		return svc.storage.GetObject(id, objectType)
	}

	var obj repo.TMFObject
	err := svc.db.Get(&obj, "SELECT * FROM tmf_object WHERE id = ? AND type = ? ORDER BY version DESC LIMIT 1", id, objectType)
	if err == sql.ErrNoRows {
		slog.Info("Service: Object not found", slog.String("id", id), slog.String("type", objectType))
		return nil, nil // Object not found
	} else if err != nil {
		err = errl.Errorf("failed to get object id=%s type=%s: %w", id, objectType, err)
	}
	return &obj, err
}

// getLocalOrRemoteObject retrieves the object from the database. If it is not found, we try to get it from the remote server (if proxy is enabled).
// If the object is not found anywhere, it returns a nil object and no error.
func (svc *Service) getLocalOrRemoteObject(req *Request) (*repo.TMFObject, error) {

	// Check if we have the object locally
	obj, err := svc.getObject(req.ID, req.ResourceName)
	if err != nil {
		return nil, errl.Errorf("failed to get object %s from local service: %w", req.ID, err)
	}

	// If proxy is not enabled, we return whatever we found, which is nil if object is not found.
	// It is not an error at this level but the caller is responsible for returning a 404 error.
	if !svc.proxyEnabled {
		return obj, nil
	}

	// If we found the object and it is not stale, return it
	if obj != nil {
		if time.Since(obj.UpdatedAt) < svc.fressness {
			slog.Debug("Object found in local database and fresh")
			return obj, nil
		}
		slog.Debug("Object found in local database but stale, retrieving from remote")
	} else {
		slog.Debug("Object not found in local database, retrieving from remote")
	}

	// The object was not found or is stale, so we have to retrieve remotely and update the local database
	remoteObj, err := svc.getRemoteObject(req)
	if err != nil {
		return nil, errl.Errorf("failed to get object %s from remote service: %w", req.ID, err)
	}

	// Store the object locally and return it to caller
	if err := svc.createObject(remoteObj); err != nil {
		slog.Error("failed to cache object", slog.Any("error", err))
		// Return the stale object or nil
		return obj, nil
	}

	slog.Info("Object retrieved from remote and cached successfully", slog.String("id", req.ID), slog.String("resourceName", req.ResourceName))
	return remoteObj, nil

}

func (svc *Service) getRemoteObject(req *Request) (*repo.TMFObject, error) {
	slog.Debug("retrieving object from remote", slog.String("id", req.ID))

	// Send the access token
	headers := map[string]string{
		"Authorization": "Bearer " + req.AccessToken,
	}

	// Build the path for the request according to TMForum specs
	path := fmt.Sprintf("/tmf-api/%s/%s/%s/%s", req.APIfamily, req.APIVersion, req.ResourceName, req.ID)

	// Send the request to the remote with our HTTP Client
	resp, err := svc.tmfClient.Get(path, headers)
	if err != nil {
		return nil, errl.Errorf("failed to proxy request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errl.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode == 404 {
		return nil, nil
	}

	if resp.StatusCode >= 300 {
		return nil, errl.Errorf("remote server returned error: %s", string(body))
	}

	// Build the object from the reply
	var incomingObjectMap map[string]any
	if err := json.Unmarshal(body, &incomingObjectMap); err != nil {
		return nil, errl.Errorf("failed to bind request body: %w", err)
	}

	id, _ := incomingObjectMap["id"].(string)
	version, _ := incomingObjectMap["version"].(string)
	lastUpdate, _ := incomingObjectMap["lastUpdate"].(string)

	// Instead of storing what we receive, we store a compact serialization of the JSON object
	incomingContent, err := json.Marshal(incomingObjectMap)
	if err != nil {
		return nil, errl.Errorf("failed to marshal object content: %w", err)
	}

	obj := repo.NewTMFObject(id, req.ResourceName, version, req.APIVersion, lastUpdate, incomingContent)

	return obj, nil
}

// updateObject updates an existing TMF object.
func (svc *Service) updateObject(obj *repo.TMFObject) error {
	slog.Debug("dbLayer: Updating object", slog.String("id", obj.ID), slog.String("type", obj.Type), slog.String("version", obj.Version))

	// Direct to another storage system if defined
	if svc.storage != nil {
		return svc.storage.UpdateObject(obj)
	}

	// Make sure timestamps are correct
	now := time.Now()
	obj.UpdatedAt = now

	// Execute the SQL
	_, err := svc.db.NamedExec(`UPDATE tmf_object SET version = :version, last_update = :last_update, content = :content, updated_at = :updated_at WHERE id = :id AND type = :type`, obj)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if sqliteErr.Code == sqlite3.ErrConstraint && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
				return &ErrObjectExists{ID: obj.ID, Type: obj.Type}
			}
		}
		err = errl.Errorf("failed to update object id=%s type=%s: %w", obj.ID, obj.Type, err)
	}
	return err
}

// deleteObject deletes a TMF object by its ID and type.
func (svc *Service) deleteObject(id, objectType string) error {
	slog.Debug("dbLayer: Deleting object", slog.String("id", id), slog.String("type", objectType))

	// Direct to another storage system if defined
	if svc.storage != nil {
		return svc.storage.DeleteObject(id, objectType)
	}

	// Execute the SQL
	_, err := svc.db.Exec("DELETE FROM tmf_object WHERE id = ? AND type = ?", id, objectType)
	if err != nil {
		err = errl.Errorf("failed to delete object id=%s type=%s: %w", id, objectType, err)
	}
	return err
}

// listObjects retrieves all TMF objects of a given type, returning only the latest version for each unique ID.
// It supports pagination, filtering, and sorting according to TMF630 guidelines.
func (svc *Service) listObjectsOld(req *Request, objectType string, apiVersion string, queryParams url.Values) ([]repo.TMFObject, int, error) {
	slog.Debug("dbLayer: Listing objects", "type", objectType, "queryParams", queryParams)

	// Direct to another storage system if defined
	if svc.storage != nil {
		return svc.storage.ListObjects(objectType, apiVersion, queryParams)
	}

	// Parse the parameters according to TM Forum specs and build the SELECT
	baseQuery, args, _, _, err := BuildSelectFromParms(objectType, queryParams)
	if err != nil {
		return nil, 0, errl.Errorf("failed to build select query: %w", err)
	}
	fmt.Printf("SQL: %s\nARGS: %v\n", baseQuery, args)

	var objs []repo.TMFObject
	var totalCount int

	// Run the SQL
	if err := svc.db.Select(&objs, baseQuery, args...); err != nil {
		err = errl.Errorf("failed to list objects for %s, params: %v: %w", objectType, queryParams, err)
		return nil, 0, err
	}

	// TODO: Implement partial field selection based on "fields" query parameter.
	// This would involve unmarshalling and then selectively marshalling the content.
	// Currently, this is done at a higher level in th eimplementation

	return objs, totalCount, err
}

func (svc *Service) listObjects(req *Request, tokenMap map[string]any, objectType string, apiVersion string, queryParams url.Values) ([]repo.TMFObject, int, error) {
	slog.Debug("dbLayer: Listing objects", "type", objectType, "queryParams", queryParams)

	// Direct to another storage system if defined
	if svc.storage != nil {
		return svc.storage.ListObjects(objectType, apiVersion, queryParams)
	}

	// Parse the parameters according to TM Forum specs and build the SELECT
	baseQuery, args, limit, offset, err := BuildSelectFromParms(objectType, queryParams)
	if err != nil {
		return nil, 0, errl.Errorf("failed to build select query: %w", err)
	}
	fmt.Printf("SQL: %s\nARGS: %v\n", baseQuery, args)

	var objs []repo.TMFObject
	var totalCount int
	var offsetCounter int

	// Run the SQL

	rows, err := svc.db.Query(baseQuery, args...)
	if err != nil {
		return nil, 0, errl.Errorf("performing query %s with args %v: %w", baseQuery, args, err)
	}
	defer rows.Close()

	const bq = `SELECT id, type, max(version) AS version, api_version, seller, buyer, last_update, content, created_at, updated_at FROM tmf_object`

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var obj repo.TMFObject
		if err := rows.Scan(&obj.ID, &obj.Type, &obj.Version, &obj.APIVersion,
			&obj.Seller, &obj.Buyer, &obj.LastUpdate, &obj.Content, &obj.CreatedAt, &obj.UpdatedAt); err != nil {

			return nil, 0, errl.Errorf("iterating over rows in query %s with args %v: %w", baseQuery, args, err)

		}

		// Check if the object can be read by the user
		authorized, err := svc.takeDecision(svc.ruleEngine, req, tokenMap, obj.GetMap())
		if !authorized {
			slog.Info("object %s not authorized: %w", obj.ID, errl.Error(err))
			continue
		}

		// We have to discard the first 'offset' objects, as specified by the user
		if offsetCounter < offset {
			offsetCounter++
			continue
		}

		// Now we add the object to the result array and check if we reached the limit as specified by the user
		objs = append(objs, obj)

		if len(objs) >= limit {
			break
		}

	}
	if err = rows.Err(); err != nil {
		return nil, 0, errl.Errorf("performing query %s with args %v: %w", baseQuery, args, err)
	}

	// ##########################################

	// TODO: Implement partial field selection based on "fields" query parameter.
	// This would involve unmarshalling and then selectively marshalling the content.
	// Currently, this is done at a higher level in th eimplementation

	return objs, totalCount, err
}

// BuildSelectFromParms creates a SELECT statement based on the query values.
// For objects with same id, selects the one with the latest version.
func BuildSelectFromParms(tmfResource string, queryValues url.Values) (query string, arguments []any, qlimit int, qoffset int, theerr error) {

	// Default values if the user did not specify them. -1 is equivalent to no values provided.
	var limit = -1
	var offset = -1

	var buf StringRenderer
	var args []any

	// The main select
	buf.Render(
		`SELECT id, type, max(version) AS version, api_version, seller, buyer, last_update, content, created_at, updated_at FROM tmf_object`,
	)

	// WHERE: normally we expect the resource name of object to be specified, but we support a query for all object types
	if len(tmfResource) > 0 {
		buf.Render(" WHERE type = ?")
		args = append(args, tmfResource)
	}

	// Build the WHERE by processing the query values specified by the user
	for key, values := range queryValues {

		if key == "sort" || key == "fields" {
			// TODO: implement processing for these parameters
			continue
		}

		selector := "[*]."
		index := strings.Index(key, "[*].")
		if index == 0 {
			err := errl.Errorf("array name in array selector is empty")
			return "", nil, 0, 0, err
		}
		if index > 0 {

			arrayName := key[:index]
			keyName := key[index+len(selector):]
			slog.Debug("array selector", "arrayName", arrayName, "keyName", keyName)

			if len(keyName) == 0 {
				err := errl.Errorf("key name in array selector is empty")
				return "", nil, 0, 0, err
			}

			vals := processValues(values)

			if len(vals) == 1 {
				buf.Render(
					" AND EXISTS (SELECT 1 FROM json_each(tmf_object.content, '$.", arrayName, "') WHERE json_extract(value, '$.", keyName, "') = ?)",
				)
			} else if len(vals) > 1 {
				buf.Render(
					" AND EXISTS (SELECT 1 FROM json_each(tmf_object.content, '$.", arrayName, "') WHERE json_extract(value, '$", keyName, "') IN ").RenderList(vals...).Render(")")
			}
			args = append(args, vals...)

			continue

		}

		switch key {
		case "limit":
			limitStr := queryValues.Get("limit")
			if limitStr != "" {
				if l, err := strconv.Atoi(limitStr); err == nil {
					limit = l
				}
			}

		case "offset":
			offsetStr := queryValues.Get("offset")
			if offsetStr != "" {
				if l, err := strconv.Atoi(offsetStr); err == nil {
					offset = l
				}
			}

		case "seller", "buyer":
			// A shortcut for DOME, to simplify life to applications (but can be also done in a TMF-compliant way).
			// Special processing to allow specifying multiple values in the form 'seller=id1,id2,id3'.
			// We also support the standard HTTP query strings like 'seller=id1,id2&seller=id3'
			vals := processValues(values)

			// Use either an equality (when one element) or an inclusion expression (when several)
			if len(vals) == 1 {
				buf.Render(" AND ", key, " = ?")
			} else if len(vals) > 1 {
				buf.Render(" AND ", key, " IN ").RenderList(vals...)
			}
			args = append(args, vals...)

		case "category.id", "productSpecification.id":
			// Simplification of the query in the category array.
			// TODO: try to generalize for all search in arrays.

			object := strings.TrimSuffix(key, ".id")

			// Special processing because TMForum allows to specify multiple values
			// in the form 'lifecycleStatus=Launched,Active'
			vals := processValues(values)

			if len(vals) == 1 {
				buf.Render(
					" AND EXISTS (SELECT 1 FROM json_each(tmf_object.content, '$.", object, "') WHERE json_extract(value, '$.id') = ?)",
				)
			} else if len(vals) > 1 {
				buf.Render(
					" AND EXISTS (SELECT 1 FROM json_each(tmf_object.content, '$.", object, "') WHERE json_extract(value, '$.id') IN ").RenderList(vals...).Render(")")
			}
			args = append(args, vals...)

		case "individualIdentification.id", "individualIdentification[*].identificationId":
			// Simplification of the query in the category array.
			object := "individualIdentification"

			// Special processing because TMForum allows to specify multiple values
			// in the form 'lifecycleStatus=Launched,Active'
			vals := processValues(values)

			if len(vals) == 1 {
				buf.Render(
					" AND EXISTS (SELECT 1 FROM json_each(tmf_object.content, '$.", object, "') WHERE json_extract(value, '$.identificationId') = ?)",
				)
			} else if len(vals) > 1 {
				buf.Render(
					" AND EXISTS (SELECT 1 FROM json_each(tmf_object.content, '$.", object, "') WHERE json_extract(value, '$.identificationId') IN ").RenderList(vals...).Render(")")
			}
			args = append(args, vals...)

		default:

			// Special processing because TMForum allows to specify multiple values
			// in the form 'lifecycleStatus=Launched,Active'
			vals := processValues(values)

			// We assume that the rest of the parameters are not in the fields of the SQL database.
			// We have to use SQLite JSON expressions to search.
			if len(vals) == 1 {
				buf.Render(" AND content->>'$.", key, "' = ?")
			} else {
				buf.Render(" AND content->>'$.", key, "' IN ").RenderList(vals...)
			}
			args = append(args, vals...)

		}
	}

	// We need to GROUP by id, so we can SELECT the record with the latest version from each group
	buf.Render(" GROUP BY id")

	// // For fairness of presenting results to customers, we want a random ordering, which is consistent and fair with the providers.
	// // Ordering by the hash of the content of the TMF object complies with the requirements, as it is consistent across paginations
	// // and nobody can predict the final ordering a-priory.
	// // For a stable catalog, the ordering is the same for all users and at any time.
	// // When a provider creates or modifies a product, it will be inserted at an unpredictable position in the catalog.
	// //
	// // TODO: we can consider a more advanced variation, where we add to the hash a random number which is
	// // generated each day or week, and keeps the same until a new one is generated.
	// // In this way, ordering is efficient, random, and changes every week (or whatever period is chosen)
	// bu.OrderBy("hash")

	// Build the query, with the statement and the arguments to be used
	sql := buf.String()

	return sql, args, limit, offset, nil
}

func BuildSelectFromParmsOld(tmfResource string, queryValues url.Values) (string, []any, error) {

	// Default values if the user did not specify them. -1 is equivalent to no values provided.
	var limit = -1
	var offset = -1

	var buf StringRenderer
	var args []any

	// The main select
	buf.Render(
		`SELECT id, type, max(version) AS version, api_version, seller, buyer, last_update, content, created_at, updated_at FROM tmf_object`,
	)

	// WHERE: normally we expect the resource name of object to be specified, but we support a query for all object types
	if len(tmfResource) > 0 {
		buf.Render(" WHERE type = ?")
		args = append(args, tmfResource)
	}

	// Build the WHERE by processing the query values specified by the user
	for key, values := range queryValues {

		if key == "sort" || key == "fields" {
			// TODO: implement processing for these parameters
			continue
		}

		selector := "[*]."
		index := strings.Index(key, "[*].")
		if index == 0 {
			err := errl.Errorf("array name in array selector is empty")
			return "", nil, err
		}
		if index > 0 {

			arrayName := key[:index]
			keyName := key[index+len(selector):]
			slog.Debug("array selector", "arrayName", arrayName, "keyName", keyName)

			if len(keyName) == 0 {
				err := errl.Errorf("key name in array selector is empty")
				return "", nil, err
			}

			vals := processValues(values)

			if len(vals) == 1 {
				buf.Render(
					" AND EXISTS (SELECT 1 FROM json_each(tmf_object.content, '$.", arrayName, "') WHERE json_extract(value, '$.", keyName, "') = ?)",
				)
			} else if len(vals) > 1 {
				buf.Render(
					" AND EXISTS (SELECT 1 FROM json_each(tmf_object.content, '$.", arrayName, "') WHERE json_extract(value, '$", keyName, "') IN ").RenderList(vals...).Render(")")
			}
			args = append(args, vals...)

			continue

		}

		switch key {
		case "limit":
			limitStr := queryValues.Get("limit")
			if limitStr != "" {
				if l, err := strconv.Atoi(limitStr); err == nil {
					limit = l
				}
			}

		case "offset":
			offsetStr := queryValues.Get("offset")
			if offsetStr != "" {
				if l, err := strconv.Atoi(offsetStr); err == nil {
					offset = l
				}
			}

		case "seller", "buyer":
			// A shortcut for DOME, to simplify life to applications (but can be also done in a TMF-compliant way).
			// Special processing to allow specifying multiple values in the form 'seller=id1,id2,id3'.
			// We also support the standard HTTP query strings like 'seller=id1,id2&seller=id3'
			vals := processValues(values)

			// Use either an equality (when one element) or an inclusion expression (when several)
			if len(vals) == 1 {
				buf.Render(" AND ", key, " = ?")
			} else if len(vals) > 1 {
				buf.Render(" AND ", key, " IN ").RenderList(vals...)
			}
			args = append(args, vals...)

		case "category.id", "productSpecification.id":
			// Simplification of the query in the category array.
			// TODO: try to generalize for all search in arrays.

			object := strings.TrimSuffix(key, ".id")

			// Special processing because TMForum allows to specify multiple values
			// in the form 'lifecycleStatus=Launched,Active'
			vals := processValues(values)

			if len(vals) == 1 {
				buf.Render(
					" AND EXISTS (SELECT 1 FROM json_each(tmf_object.content, '$.", object, "') WHERE json_extract(value, '$.id') = ?)",
				)
			} else if len(vals) > 1 {
				buf.Render(
					" AND EXISTS (SELECT 1 FROM json_each(tmf_object.content, '$.", object, "') WHERE json_extract(value, '$.id') IN ").RenderList(vals...).Render(")")
			}
			args = append(args, vals...)

		case "individualIdentification.id", "individualIdentification[*].identificationId":
			// Simplification of the query in the category array.
			object := "individualIdentification"

			// Special processing because TMForum allows to specify multiple values
			// in the form 'lifecycleStatus=Launched,Active'
			vals := processValues(values)

			if len(vals) == 1 {
				buf.Render(
					" AND EXISTS (SELECT 1 FROM json_each(tmf_object.content, '$.", object, "') WHERE json_extract(value, '$.identificationId') = ?)",
				)
			} else if len(vals) > 1 {
				buf.Render(
					" AND EXISTS (SELECT 1 FROM json_each(tmf_object.content, '$.", object, "') WHERE json_extract(value, '$.identificationId') IN ").RenderList(vals...).Render(")")
			}
			args = append(args, vals...)

		default:

			// Special processing because TMForum allows to specify multiple values
			// in the form 'lifecycleStatus=Launched,Active'
			vals := processValues(values)

			// We assume that the rest of the parameters are not in the fields of the SQL database.
			// We have to use SQLite JSON expressions to search.
			if len(vals) == 1 {
				buf.Render(" AND content->>'$.", key, "' = ?")
			} else {
				buf.Render(" AND content->>'$.", key, "' IN ").RenderList(vals...)
			}
			args = append(args, vals...)

		}
	}

	// We need to GROUP by id, so we can SELECT the record with the latest version from each group
	buf.Render(" GROUP BY id")

	// // For fairness of presenting results to customers, we want a random ordering, which is consistent and fair with the providers.
	// // Ordering by the hash of the content of the TMF object complies with the requirements, as it is consistent across paginations
	// // and nobody can predict the final ordering a-priory.
	// // For a stable catalog, the ordering is the same for all users and at any time.
	// // When a provider creates or modifies a product, it will be inserted at an unpredictable position in the catalog.
	// //
	// // TODO: we can consider a more advanced variation, where we add to the hash a random number which is
	// // generated each day or week, and keeps the same until a new one is generated.
	// // In this way, ordering is efficient, random, and changes every week (or whatever period is chosen)
	// bu.OrderBy("hash")

	// Pagination support
	buf.Render(" LIMIT ", limit, " OFFSET ", offset)

	// Build the query, with the statement and the arguments to be used
	sql := buf.String()

	return sql, args, nil
}

type StringRenderer struct {
	strings.Builder
}

func (r *StringRenderer) Render(inputs ...any) *StringRenderer {
	for _, s := range inputs {
		switch v := s.(type) {
		case string:
			r.WriteString(v)
		case []byte:
			r.Write(v)
		case int:
			r.WriteString(strconv.FormatInt(int64(v), 10))
		case byte:
			r.WriteByte(v)
		case rune:
			r.WriteRune(v)
		default:
			slog.Error("attemping to write something not a string, int, rune, []byte or byte: %T", s)
		}
	}
	return r
}

func (r *StringRenderer) Renderln(inputs ...any) *StringRenderer {
	r.Render(inputs...)
	r.Render('\n')
	return r
}

// RenderList renders the arguments as a comma separated list inside parenthesis
func (r *StringRenderer) RenderList(inputs ...any) *StringRenderer {
	r.Render("(")
	for i := range inputs {
		if i > 0 {
			r.Render(",")
		}
		r.Render("?")
	}
	r.Render(")")
	return r
}

func processValues(values []string) []any {
	// Special processing because TMForum allows to specify multiple values
	// in the form 'lifecycleStatus=Launched,Active'
	var vals []any
	// Allow several instances of 'lifecycleStatus' parameter in the query string
	for _, v := range values {
		parts := strings.Split(v, ",")
		// Allow for whitespace surrounding the elements
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
			vals = append(vals, parts[i])
		}
	}

	return vals
}

// DeleteTables drops the table and performs a VACUUM to reclaim space
func (svc *Service) DeleteTables() error {

	_, err := svc.db.Exec(repo.DeleteTMFTableSQL)
	if err != nil {
		slog.Error("deleteTables: drop table", slogor.Err(err))
		return errl.Errorf("deleteTables: drop table: %w", err)
	}

	_, err = svc.db.Exec(repo.VacuumTMFTableSQL)
	if err != nil {
		slog.Error("deleteTables: vacuum", slogor.Err(err))
		return errl.Errorf("deleteTables: vacuum: %w", err)
	}

	return nil
}
