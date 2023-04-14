package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upWidgets, downWidgets)
}

func upWidgets(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return nil
}

func downWidgets(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
