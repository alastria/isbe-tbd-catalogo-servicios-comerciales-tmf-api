package repository

import (
	"github.com/hesusruiz/isbetmf/internal/errl"
	"github.com/jmoiron/sqlx"
)

// CreateTMFTableSQL is the table 'tmf_object', the one holding all objects of all types
const CreateTMFTableSQL = `CREATE TABLE IF NOT EXISTS tmf_object (
	"id" TEXT NOT NULL,
	"type" TEXT NOT NULL,
	"version" TEXT,
	"api_version" TEXT,
	"seller" TEXT,
	"seller_operator" TEXT,
	"buyer" TEXT,
	"buyer_operator" TEXT,
	"last_update" TEXT,
	"content" BLOB NOT NULL,
	"random" INTEGER DEFAULT 0,
	"created_at" INTEGER DEFAULT CURRENT_TIMESTAMP,
	"updated_at" INTEGER DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY ("id", "type", "version")
);`

// Used for migrations and to keep the database file compact
const DeleteTMFTableSQL = `DROP TABLE IF EXISTS tmf_object;`
const VacuumSQL = `VACUUM;`

// CreateTables creates the tables in the database if they do not exist.
// It also handles automatic schema/data migration when possible.
func CreateTables(db *sqlx.DB) error {

	if _, err := db.Exec(CreateTMFTableSQL); err != nil {
		return errl.Errorf("failed to create tmf_object table: %w", err)
	}

	if err := RunMigrationsUp(db); err != nil {
		return errl.Error(err)
	}

	// if err := MigrateTables(db, oldMajor, oldMinor); err != nil {
	// 	return errl.Error(err)
	// }

	return nil
}
