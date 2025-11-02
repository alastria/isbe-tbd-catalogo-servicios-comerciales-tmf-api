package repository

import (
	"fmt"
	"log/slog"

	"github.com/hesusruiz/isbetmf/internal/errl"
	"github.com/jmoiron/sqlx"
)

// These denote the current version of the database schema and data semantics.
// If we find an old version of the database, we try an automatic migration.
const DBVersionMajor = 2
const DBVersionMinor = 1
const DBAPIVersion = "v4"

// CreateConfigTableSQL is the table 'tmf-config' for our versioning of table schemas and data
const CreateConfigTableSQL = `CREATE TABLE IF NOT EXISTS tmf_config (
	"version" TEXT DEFAULT '0.0',
	"api_version" TEXT DEFAULT 'v4',
	"created_at" DATETIME NOT NULL,
	"updated_at" DATETIME NOT NULL
);`

// CreateTMFTableSQL is the table 'tmf_object', the one holding all objects of all types
const CreateTMFTableSQL = `CREATE TABLE IF NOT EXISTS tmf_object (
	"id" TEXT NOT NULL,
	"type" TEXT NOT NULL,
	"version" TEXT DEFAULT '1.0',
	"api_version" TEXT DEFAULT 'v4',
	"seller" TEXT DEFAULT '',
	"seller_operator" TEXT DEFAULT '',
	"buyer" TEXT DEFAULT '',
	"buyer_operator" TEXT DEFAULT '',
	"last_update" TEXT,
	"content" BLOB NOT NULL,
	"random" INTEGER DEFAULT 0,
	"created_at" DATETIME NOT NULL,
	"updated_at" DATETIME NOT NULL,
	PRIMARY KEY ("id", "type", "version")
);`

// Used for migrations and to keep the database file compact
const DeleteTMFTableSQL = `DROP TABLE IF EXISTS tmf_object;`
const VacuumSQL = `VACUUM;`

// CreateTables creates the tables in the database if they do not exist.
// It also handles automatic schema/data migration when possible.
func CreateTables(db *sqlx.DB) error {

	// Create the config table if it doesn't exist
	if _, err := db.Exec(CreateConfigTableSQL); err != nil {
		return errl.Errorf("failed to create config table: %w", err)
	}

	newVersionStr := fmt.Sprintf("%d.%d", DBVersionMajor, DBVersionMinor)

	// Get the current version from the config table
	var oldVersion string
	err := db.Get(&oldVersion, "SELECT version FROM tmf_config LIMIT 1;")
	if err != nil {
		// If there is an error, it might be because the table is empty, so we insert the initial version
		_, err = db.Exec(
			"INSERT INTO tmf_config (version, api_version, created_at, updated_at) VALUES (?, ?, datetime('now'), datetime('now'));",
			newVersionStr, DBAPIVersion,
		)
		if err != nil {
			return errl.Errorf("failed to insert initial config version: %w", err)
		}
		// We just created the database
		oldVersion = "0.0"
	}

	// Parse the current version as oldMajor.oldMinor
	var oldMajor, oldMinor int
	_, err = fmt.Sscanf(oldVersion, "%d.%d", &oldMajor, &oldMinor)
	if err != nil {
		return errl.Errorf("failed to parse current version: %w", err)
	}

	slog.Info("Database version", slog.String("old", oldVersion), slog.String("new", newVersionStr))

	// Always update the version in the config table
	_, err = db.Exec("UPDATE tmf_config SET version = ?, updated_at = datetime('now');", fmt.Sprintf("%d.%d", DBVersionMajor, DBVersionMinor))
	if err != nil {
		return errl.Errorf("failed to update config version: %w", err)
	}

	if _, err := db.Exec(CreateTMFTableSQL); err != nil {
		return errl.Errorf("failed to create tmf_object table: %w", err)
	}

	if err := MigrateTables(db, oldMajor, oldMinor); err != nil {
		return errl.Error(err)
	}

	return nil
}

func MigrateTables(db *sqlx.DB, oldMajor, oldMinor int) error {

	if oldMajor > DBVersionMajor {
		return errl.Errorf("the existing database version %d is newer than our program %d", oldMajor, DBVersionMajor)
	}
	// if oldMajor == DBVersionMajor {
	// 	// Nothing to be done. Minor versions are just for documentation
	// 	return nil
	// }
	if DBVersionMajor-oldMajor > 1 {
		return errl.Errorf("the existing database version %d is too old to be migrated automatically by this program %d", oldMajor, DBVersionMajor)
	}

	// // Migration: add fields seller_operator and buyer_operator, to improve queries for access control
	// const AddSellerOperatorFieldSQL = `ALTER TABLE tmf_object ADD COLUMN seller_operator TEXT DEFAULT '';`
	// const AddBuyerOperatorFieldSQL = `ALTER TABLE tmf_object ADD COLUMN buyer_operator TEXT DEFAULT '';`

	// if _, err := db.Exec(AddSellerOperatorFieldSQL); err != nil {
	// 	return errl.Errorf("failed to add seller_operator field: %w", err)
	// }

	// if _, err := db.Exec(AddBuyerOperatorFieldSQL); err != nil {
	// 	return errl.Errorf("failed to add buyer_operator field: %w", err)
	// }

	// // Compact the database
	// if _, err := db.Exec(VacuumSQL); err != nil {
	// 	return errl.Errorf("failed to vacuum database: %w", err)
	// }

	// // Always update the version in the config table
	// if _, err := db.Exec("UPDATE tmf_config SET version = ?, updated_at = datetime('now');", fmt.Sprintf("%d.%d", DBVersionMajor, DBVersionMinor)); err != nil {
	// 	return errl.Errorf("failed to update config version: %w", err)
	// }

	tx, err := db.Begin()
	if err != nil {
		return errl.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Scan all existing records to set the new fields to the proper value
	const queryDB = `SELECT id, type, version, api_version, seller, buyer, last_update, content, created_at, updated_at FROM tmf_object`
	rows, err := tx.Query(queryDB)
	if err != nil {
		return errl.Errorf("failed to query tmf_object table: %w", err)
	}
	defer rows.Close()

	// Upgrade each row. If we detect a problem with one row, we just log the error and continue with the rest.
	// But if we get a database error, we stop migration.
	for rows.Next() {
		var obj TMFObject
		if err := rows.Scan(&obj.ID, &obj.Type, &obj.Version, &obj.APIVersion,
			&obj.Seller, &obj.Buyer, &obj.LastUpdate, &obj.Content, &obj.CreatedAt, &obj.UpdatedAt); err != nil {
			return errl.Errorf("iterating over rows in query %s: %w", queryDB, err)
		}

		// Get the map from the content
		objMap, err := NewTMFObjectMapFromRequest(obj.Type, obj.Content)
		if err != nil {
			err = errl.Errorf("failed to parse object content: %w", err)
			slog.Error(err.Error())
			continue
		}

		// Get the seller and sellerOperator info
		seller, sellerOperator, err := objMap.GetSellerInfo(obj.APIVersion)
		if err != nil {
			err = errl.Errorf("object %s, failed to get seller: %w", objMap.ID(), err)
			slog.Error(err.Error())
		}
		obj.Seller = seller
		obj.SellerOperator = sellerOperator

		// Get the buyer and buyerOperator info
		if obj.ID == "urn:ngsi-ld:product-specification:839e37f8-5f9d-48fc-b2c9-3e5a32d90424" {
			fmt.Println(obj.Type)
		}
		buyer, buyerOperator, err := objMap.GetBuyerInfo(obj.APIVersion)
		if err != nil {
			err = errl.Errorf("object %s, failed to get buyer: %w", objMap.ID(), err)
			slog.Error(err.Error())
		}
		obj.Buyer = buyer
		obj.BuyerOperator = buyerOperator

		// Update the database
		if _, err := tx.Exec("UPDATE tmf_object SET seller = ?, seller_operator = ?, buyer = ?, buyer_operator = ? WHERE id = ? AND type = ? AND version = ?",
			obj.Seller, obj.SellerOperator, obj.Buyer, obj.BuyerOperator, obj.ID, obj.Type, obj.Version); err != nil {
			return errl.Errorf("object %s, failed to update tmf_object table: %w", objMap.ID(), err)
		}

	}

	if err := tx.Commit(); err != nil {
		return errl.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
