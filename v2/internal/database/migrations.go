package database

import "database/sql"

func runMigrations(db *sql.DB, current int) error {

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

func ensureSchemaVersion(db *sql.DB) error {

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
		return runMigrations(db, version)
	}

	return nil
}
