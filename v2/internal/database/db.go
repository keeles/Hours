package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"time"

	_ "modernc.org/sqlite"
)

type Projects map[string]Project

type Project struct {
	Name  string `json:"Name"`
	Hours int    `json:"Hours"`
}

type Tasks map[string]int

type Timer struct {
	ClientName string
	TaskName   *string // pointer = optional
	StartTime  time.Time
}

const SchemaVersion = 2

func GetDBPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to determine home directory - please set $HOME environment variable")
	}

	// TODO: remove v2 when publishing final version
	configDir := filepath.Join(home, ".config", "hours", "v2")

	err = os.MkdirAll(configDir, 0755)
	if err != nil {
		return "", fmt.Errorf("failed to create $HOME/.config/hours directory for database storage")
	}

	return filepath.Join(configDir, "hours.db"), nil
}

func InitDb() (*sql.DB, error) {

	dbPath, err := GetDBPath()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`PRAGMA foreign_keys = ON;`)
	if err != nil {
		return nil, err
	}

	// Ensure core tables exist
	schema := `
	CREATE TABLE IF NOT EXISTS schema_version (
		version INTEGER NOT NULL
	);

	CREATE TABLE IF NOT EXISTS clients (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT UNIQUE NOT NULL
	);

	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		client_id INTEGER NOT NULL,
		name TEXT NOT NULL,
		minutes INTEGER DEFAULT 0,
		FOREIGN KEY (client_id) REFERENCES clients(id) ON DELETE CASCADE,
		UNIQUE(id, name)
	);

	-- Only ONE active timer allowed
	CREATE TABLE IF NOT EXISTS active_timer (
		id INTEGER PRIMARY KEY CHECK (id = 1),
		client_id INTEGER NOT NULL,
		task_id INTEGER,
		start_time TEXT NOT NULL,
		FOREIGN KEY (client_id) REFERENCES clients(id) ON DELETE CASCADE,
		FOREIGN KEY (task_id) REFERENCES tasks(id) ON DELETE SET NULL
	);

	CREATE TABLE IF NOT EXISTS time_entries (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		task_id INTEGER NOT NULL,
		start_time TEXT NOT NULL,
		end_time TEXT NOT NULL,
		minutes INTEGER NOT NULL,
		FOREIGN KEY (task_id) REFERENCES tasks(id) ON DELETE CASCADE
	);
	`

	_, err = db.Exec(schema)
	if err != nil {
		return nil, err
	}

	err = ensureSchemaVersion(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}
