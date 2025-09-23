package service

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/url"
	"strconv"
	"strings"

	"github.com/hesusruiz/isbetmf/internal/errl"
	repo "github.com/hesusruiz/isbetmf/tmfserver/repository"
	sqlb "github.com/huandu/go-sqlbuilder"
	"github.com/mattn/go-sqlite3"
)

// createObject creates a new TMF object.
func (svc *Service) createObject(obj *repo.TMFObject) error {
	slog.Debug("dbLayer: Creating object", slog.String("id", obj.ID), slog.String("type", obj.Type), slog.String("version", obj.Version))
	if svc.storage != nil {
		return svc.storage.CreateObject(obj)
	}
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

// getObject retrieves a TMF object by its ID and type, returning the latest version.
func (svc *Service) getObject(id, objectType string) (*repo.TMFObject, error) {
	slog.Debug("dbLayer: Getting object", slog.String("id", id), slog.String("type", objectType))
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

// updateObject updates an existing TMF object.
func (svc *Service) updateObject(obj *repo.TMFObject) error {
	slog.Debug("dbLayer: Updating object", slog.String("id", obj.ID), slog.String("type", obj.Type), slog.String("version", obj.Version))
	if svc.storage != nil {
		return svc.storage.UpdateObject(obj)
	}
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
	if svc.storage != nil {
		return svc.storage.DeleteObject(id, objectType)
	}
	_, err := svc.db.Exec("DELETE FROM tmf_object WHERE id = ? AND type = ?", id, objectType)
	if err != nil {
		err = errl.Errorf("failed to delete object id=%s type=%s: %w", id, objectType, err)
	}
	return err
}

// listObjects retrieves all TMF objects of a given type, returning only the latest version for each unique ID.
// It supports pagination, filtering, and sorting according to TMF630 guidelines.
func (svc *Service) listObjectsOld(objectType string, apiVersion string, queryParams url.Values) ([]repo.TMFObject, int, error) {
	slog.Debug("dbLayer: Listing objects", "type", objectType, "queryParams", queryParams)
	q, a := BuildSelectFromParms(objectType, queryParams)
	fmt.Printf("SQL: %s\nARGS: %v\n", q, a)

	if svc.storage != nil {
		return svc.storage.ListObjects(objectType, apiVersion, queryParams)
	}
	var objs []repo.TMFObject
	var totalCount int

	// Base query to get the latest version for each unique ID and type
	baseQuery := `
		SELECT t1.*
		FROM tmf_object t1
		INNER JOIN (
			SELECT id, type, MAX(version) AS max_version
			FROM tmf_object
			WHERE type = ? AND api_version = ?
			GROUP BY id, type
		) AS t2
		ON t1.id = t2.id AND t1.type = t2.type AND t1.version = t2.max_version
		WHERE t1.type = ? AND t1.api_version = ?
	`
	countQuery := `
		SELECT COUNT(DISTINCT t1.id)
		FROM tmf_object t1
		INNER JOIN (
			SELECT id, type, MAX(version) AS max_version
			FROM tmf_object
			WHERE type = ? AND api_version = ?
			GROUP BY id, type
		) AS t2
		ON t1.id = t2.id AND t1.type = t2.type AND t1.version = t2.max_version
		WHERE t1.type = ? AND t1.api_version = ?
	`

	args := []any{objectType, apiVersion, objectType, apiVersion}
	countArgs := []any{objectType, apiVersion, objectType, apiVersion}

	// Add filters
	filterClauses := []string{}
	for key, values := range queryParams {
		// TMF630 reserved words for query parameters
		if key == "limit" || key == "offset" || key == "sort" || key == "fields" {
			continue
		}
		// Assuming simple equality filter for now
		filterClauses = append(filterClauses, fmt.Sprintf("json_extract(t1.content, '$.%s') = ?", key))
		args = append(args, values[0])
		countArgs = append(countArgs, values[0])
	}

	if len(filterClauses) > 0 {
		filterString := " AND " + strings.Join(filterClauses, " AND ")
		baseQuery += filterString
		countQuery += filterString
	}

	// Get total count before pagination
	err := svc.db.Get(&totalCount, countQuery, countArgs...)
	if err != nil {
		err = errl.Errorf("failed to get total count for %s, params: %v: %w", objectType, queryParams, err)
		return nil, 0, err
	}

	// Add sorting
	sortParam := queryParams.Get("sort")
	if sortParam != "" {
		sortFields := strings.Split(sortParam, ",")
		orderByClauses := []string{}
		for _, field := range sortFields {
			direction := "ASC"
			if strings.HasPrefix(field, "-") {
				direction = "DESC"
				field = field[1:]
			}
			orderByClauses = append(orderByClauses, fmt.Sprintf("json_extract(t1.content, '$.%s') %s", field, direction))
		}
		baseQuery += " ORDER BY " + strings.Join(orderByClauses, ", ")
	}

	// Add pagination
	limitStr := queryParams.Get("limit")
	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err == nil && limit > 0 {
			baseQuery += fmt.Sprintf(" LIMIT %d", limit)
		}
	}

	offsetStr := queryParams.Get("offset")
	if offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err == nil && offset >= 0 {
			baseQuery += fmt.Sprintf(" OFFSET %d", offset)
		}
	}

	err = svc.db.Select(&objs, baseQuery, args...)
	if err != nil {
		err = errl.Errorf("failed to list objects for %s, params: %v: %w", objectType, queryParams, err)
		return nil, 0, err
	}

	// TODO: Implement partial field selection based on "fields" query parameter.
	// This would involve unmarshalling and then selectively marshalling the content.
	// Currently, this is done at a higher level in th eimplementation

	return objs, totalCount, err
}

func (svc *Service) listObjects(objectType string, apiVersion string, queryParams url.Values) ([]repo.TMFObject, int, error) {
	slog.Debug("dbLayer: Listing objects", "type", objectType, "queryParams", queryParams)
	baseQuery, args := BuildSelectFromParms(objectType, queryParams)
	fmt.Printf("SQL: %s\nARGS: %v\n", baseQuery, args)

	if svc.storage != nil {
		return svc.storage.ListObjects(objectType, apiVersion, queryParams)
	}
	var objs []repo.TMFObject
	var totalCount int

	err := svc.db.Select(&objs, baseQuery, args...)
	if err != nil {
		err = errl.Errorf("failed to list objects for %s, params: %v: %w", objectType, queryParams, err)
		return nil, 0, err
	}

	// TODO: Implement partial field selection based on "fields" query parameter.
	// This would involve unmarshalling and then selectively marshalling the content.
	// Currently, this is done at a higher level in th eimplementation

	return objs, totalCount, err
}

// BuildSelectFromParms creates a SELECT statement based on the query values.
// For objects with same id, selects the one with the latest version.
func BuildSelectFromParms(tmfResource string, queryValues url.Values) (string, []any) {

	// Default values if the user did not specify them. -1 is equivalent to no values provided.
	var limit = -1
	var offset = -1

	bu := sqlb.SQLite.NewSelectBuilder()

	// SELECT: for each object with a given id, select the latest version.
	// We use the 'max(version)' function, and will GROUP by id.
	bu.Select(
		"id",
		"type",
		"max(version) AS version",
		"api_version",
		"seller",
		"buyer",
		"last_update",
		"content",
		"created_at",
		"updated_at",
	).From("tmf_object")

	// WHERE: normally we expect the resource name of object to be specified, but we support a query for all object types
	if len(tmfResource) > 0 {
		bu.Where(bu.Equal("type", tmfResource))
	}

	// Build the WHERE by processing the query values specified by the user
	whereClause := sqlb.NewWhereClause()
	cond := sqlb.NewCond()

	for key, values := range queryValues {

		if key == "sort" || key == "fields" {
			// TODO: implement processing for these parameters
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
		// case "lifecycleStatus":
		// 	// Special processing because TMForum allows to specify multiple values
		// 	// in the form 'lifecycleStatus=Launched,Active'
		// 	var vals = []string{}
		// 	// Allow several instances of 'lifecycleStatus' parameter in the query string
		// 	for _, v := range values {
		// 		parts := strings.Split(v, ",")
		// 		// Allow for whitespace surrounding the elements
		// 		for i := range parts {
		// 			parts[i] = strings.TrimSpace(parts[i])
		// 		}
		// 		vals = append(vals, parts...)
		// 	}

		// 	// Use either an equality or an inclusion expression
		// 	if len(vals) == 1 {
		// 		whereClause.AddWhereExpr(
		// 			cond.Args,
		// 			cond.Equal(key, sqlb.List(vals)),
		// 		)
		// 	} else {
		// 		whereClause.AddWhereExpr(
		// 			cond.Args,
		// 			cond.In(key, sqlb.List(vals)),
		// 		)
		// 	}

		case "seller", "buyer":
			// A shortcut for DOME, to simplify life to applications (but can be also done in a TMF-compliant way).
			// Special processing to allow specifying multiple values in the form 'seller=id1,id2,id3'.
			// We also support the standard HTTP query strings like 'seller=id1,id2&seller=id3'
			var vals = []string{}
			// Allow several instances of the key in the query string (as in standard HTTP query strings)
			for _, v := range values {
				// Process each for several comma-separated values in the same key instance
				parts := strings.Split(v, ",")
				// Allow for whitespace surrounding the elements
				for i := range parts {
					parts[i] = strings.TrimSpace(parts[i])
				}
				vals = append(vals, parts...)
			}

			// Use either an equality (when one element) or an inclusion expression (when several)
			if len(vals) == 1 {
				whereClause.AddWhereExpr(
					cond.Args,
					cond.Equal(key, sqlb.List(vals)),
				)
			} else {
				whereClause.AddWhereExpr(
					cond.Args,
					cond.In(key, sqlb.List(vals)),
				)
			}

		default:

			// We assume that the rest of the parameters are not in the fields of the SQL database.
			// We have to use SQLite JSON expressions to search.
			if len(values) == 1 {
				whereClause.AddWhereExpr(
					cond.Args,
					cond.Equal("content->>'$."+key+"'", values[0]),
				)
			} else {
				whereClause.AddWhereExpr(
					cond.Args,
					cond.In("content->>'$."+key+"'", sqlb.List(values)),
				)

			}

		}
	}

	// Add the WHERE to the SELECT
	bu.AddWhereClause(whereClause)

	// We need to GROUP by id, so we can SELECT the record with the latest version from each group
	bu.GroupBy("id")

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
	bu.Limit(limit).Offset(offset)

	// Build the query, with the statement and the arguments to be used
	sql, args := bu.Build()

	return sql, args
}
