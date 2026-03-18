package config

import (
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	db "github.com/keeles/hours/v2/internal/database"
	"github.com/keeles/hours/v2/internal/lib"
	"github.com/keeles/hours/v2/internal/logger"
)

func (o AddDirectoryOptions) Run(ctx *kong.Context) error {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error reading current working directory")
		return err
	}

	exists, err := db.ClientExists(o.Client)
	if err != nil {
		logger.FileNotExists()
		return err
	}

	if !exists {
		fmt.Printf("Client with name %s not found.", o.Client)
		return nil
	}

	return lib.AppendDirectory(o.Client, dir)
}

func (o RemoveDirectoryOptions) Run(ctx *kong.Context) error {
	if o.Client == "" {
		dir, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("Error reading current working directory: %w", err)
		}

		return lib.RemoveDirectory(dir)
	}

	return lib.RemoveDirectoryOfClient(o.Client)
}

func (o ListOptions) Run(ctx *kong.Context) error {
	return nil
}
