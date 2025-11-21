package list

import (
	"fmt"

	"github.com/keeles/hours/internal/lib"
	"github.com/keeles/hours/internal/logger"

	"github.com/alecthomas/kong"
)

func (o Options) Run (ctx *kong.Context) error {
	data, err := lib.ReadFile()
	if err != nil {
		logger.FileNotExists()
		return nil
	}

	for _, project := range data {
		fmt.Printf("%v: %d hours\n", project.Name, project.Hours)
	}
	return nil
}