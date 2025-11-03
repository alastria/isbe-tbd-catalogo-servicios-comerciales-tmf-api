package migrations

import (
	"log/slog"

	"github.com/hesusruiz/isbetmf/tmfserver/repository"
	"github.com/jmoiron/sqlx"
)

func init() {
	repository.RegisterMigration("20251102T203201", migration_up_20251102T203201, nil)
}

func migration_up_20251102T203201(db *sqlx.DB) error {
	slog.Info("Hello from migration_up_20251102T203201")

	// Here there would be the migration logic
	// ...

	// If we return an error, the migration transaction will be cancelled.
	// Otherwise, the transaction is committed and the next migration (if any) is run.
	return nil
}
