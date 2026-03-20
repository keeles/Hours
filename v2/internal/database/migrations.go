package database

import (
	"database/sql"
	"fmt"
)

func RunMigrations(db *sql.DB, current int) error {
	if current < 2 {
		// v1 stored hours instead of minutes
		_, err := db.Exec(`
			ALTER TABLE tasks ADD COLUMN minutes INTEGER DEFAULT 0;
		`)
		if err != nil {
			return err
		}

		_, err = db.Exec(`
			UPDATE tasks SET minutes = hours * 60;
		`)
		if err != nil {
			return err
		}

		_, err = db.Exec(`UPDATE schema_version SET version = 2`)
		if err != nil {
			return err
		}
	}

	return nil
}

func EnsureSchemaVersion(db *sql.DB) error {
	var version int

	err := db.QueryRow(`SELECT version FROM schema_version LIMIT 1`).Scan(&version)

	if err == sql.ErrNoRows {
		// Fresh install
		_, err := db.Exec(`INSERT INTO schema_version(version) VALUES (?)`, SchemaVersion)
		return err

	}

	if err != nil {
		return err
	}

	if version < SchemaVersion {
		fmt.Println("Old database version detected, running migrations")
		fmt.Println(version)
		return RunMigrations(db, version)
	}

	return nil
}
