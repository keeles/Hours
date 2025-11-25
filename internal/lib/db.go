package lib

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

type Projects map[string]Project

type Project struct {
	Name  string `json:"Name"`
	Hours int    `json:"Hours"`
}

type Tasks map[string]int

func GetDBPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to determine home directory - please set $HOME environment variable")
	}

	configDir := filepath.Join(home, ".config", "hours")
	err = os.MkdirAll(configDir, 0755)
	if err != nil {
		return "", fmt.Errorf("failed to create $HOME/.config/hours directory for database storage")
	}

	return filepath.Join(configDir, "hours.db"), nil
}

func InitDb() (*sql.DB, error) {
	dbPath, err := GetDBPath()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		fmt.Println("Failing Here")
		return nil, err
	}

	schema := `
	CREATE TABLE IF NOT EXISTS clients (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT UNIQUE NOT NULL	
	);
	
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		client_id INTEGER NOT NULL,
		name TEXT NOT NULL,
		hours INTEGER DEFAULT 0,
		FOREIGN KEY (client_id) REFERENCES clients(id) ON DELETE CASCADE,
		UNIQUE(client_id, name)
	);
	`

	_, err = db.Exec(schema)
	if err != nil {
		return nil, err
	}

	return db, nil
}
