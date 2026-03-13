package logger

import (
	"fmt"
	"os"

	db "github.com/keeles/hours/v2/internal/database"
)

func ProjectNotFound(name string) {
	fmt.Fprintf(os.Stderr, "Error, could not find project %v\n", name)
}

func FileNotExists() {
	pathname, err := db.GetDBPath()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
	}
	fmt.Fprintf(os.Stderr, "Error, could not find database at %v", pathname)
}

func ErrorWritingFile() {
	pathname, err := db.GetDBPath()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
	}
	fmt.Fprintf(os.Stderr, "Error, could not write to database at %v", pathname)
}

func ProjectAlreadyExists(name string) {
	fmt.Fprintf(os.Stderr, "Error, project with name %v already exists", name)
}

func PrintVersion(version string) {
	fmt.Printf("Current Hours version: %v \n", version)
}
