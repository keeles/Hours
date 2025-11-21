package new

import (
	"fmt"
	"strings"

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
	_, exists := data[strings.ToLower(o.Name)]
	if exists {
		logger.ProjectAlreadyExists(o.Name)
		return nil
	}
	err = lib.AddNewProject(data, o.Name)
	if err != nil {
		logger.ErrorWritingFile()
		return nil
	}
	fmt.Printf("New Project Created: %s\n", o.Name)
	return nil
}