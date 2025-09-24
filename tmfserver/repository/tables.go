package repository

import (
	"fmt"

	"github.com/hesusruiz/isbetmf/internal/errl"
	"github.com/jmoiron/sqlx"
)

const DBVersionMajor = 0
const DBVersionMinor = 1

const CreateConfigTableSQL = `CREATE TABLE IF NOT EXISTS tmf_config (
	"version" TEXT DEFAULT '0.1',
	"api_version" TEXT DEFAULT 'v4',
	"created_at" DATETIME NOT NULL,
	"updated_at" DATETIME NOT NULL
);`

const CreateTMFTableSQL = `CREATE TABLE IF NOT EXISTS tmf_object (
	"id" TEXT NOT NULL,
	"type" TEXT NOT NULL,
	"version" TEXT DEFAULT '1.0',
	"api_version" TEXT DEFAULT 'v4',
	"seller" TEXT DEFAULT '',
	"buyer" TEXT DEFAULT '',
	"last_update" TEXT,
	"content" BLOB NOT NULL,
	"random" INTEGER DEFAULT 0,
	"created_at" DATETIME NOT NULL,
	"updated_at" DATETIME NOT NULL,
	PRIMARY KEY ("id", "type", "version")
);`

const DeleteTMFTableSQL = `DROP TABLE IF EXISTS tmf_object;`
const VacuumTMFTableSQL = `VACUUM;`

func CreateTables(db *sqlx.DB) error {

	// Create the config table if it doesn't exist
	if _, err := db.Exec(CreateConfigTableSQL); err != nil {
		return errl.Errorf("failed to create config table: %w", err)
	}

	ourVersionStr := fmt.Sprintf("%d.%d", DBVersionMajor, DBVersionMinor)

	// Get the current version from the config table
	var currentVersion string
	err := db.Get(&currentVersion, "SELECT version FROM tmf_config LIMIT 1;")
	if err != nil {
		// If there is an error, it might be because the table is empty, so we insert the initial version
		_, err = db.Exec("INSERT INTO tmf_config (version, api_version, created_at, updated_at) VALUES (?, ?, datetime('now'), datetime('now'));", ourVersionStr, "v4")
		if err != nil {
			return errl.Errorf("failed to insert initial config version: %w", err)
		}
		currentVersion = "0.0"
	}

	// Parse the current version as CurrentMajor.CurrentMinor
	// For simplicity, we assume the version is always in the correct format
	// In a real-world scenario, you would want to add error handling here
	var CurrentMajor, CurrentMinor int
	_, err = fmt.Sscanf(currentVersion, "%d.%d", &CurrentMajor, &CurrentMinor)
	if err != nil {
		return errl.Errorf("failed to parse current version: %w", err)
	}

	// Delete the database in case of a major version change
	if CurrentMajor < DBVersionMajor {
		if _, err := db.Exec(DeleteTMFTableSQL); err != nil {
			return errl.Errorf("failed to delete tmf_object table: %w", err)
		}

		if _, err := db.Exec(VacuumTMFTableSQL); err != nil {
			return errl.Errorf("failed to vacuum database: %w", err)
		}

	}

	// Always update the version in the config table
	_, err = db.Exec("UPDATE tmf_config SET version = ?, updated_at = datetime('now');", fmt.Sprintf("%d.%d", DBVersionMajor, DBVersionMinor))
	if err != nil {
		return errl.Errorf("failed to update config version: %w", err)
	}

	if _, err := db.Exec(CreateTMFTableSQL); err != nil {
		return errl.Errorf("failed to create tmf_object table: %w", err)
	}

	return nil
}
