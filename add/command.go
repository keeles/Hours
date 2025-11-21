package add

import (
	"strings"

	"github.com/keeles/hours/internal/lib"
	"github.com/keeles/hours/internal/logger"

	"github.com/alecthomas/kong"
)

func (o Options) Run (ctx * kong.Context) error {
	data, err := lib.ReadFile()
	if err != nil {
		logger.FileNotExists()
		return nil
	}

	project := data[strings.ToLower(o.Name)]
	if (lib.Project{} == project) {
		logger.ProjectNotFound(o.Name)
		return nil
	}

	project.Hours += o.NewHours 
	data[strings.ToLower(o.Name)] = project
	err = lib.WriteFile(data)
	if err != nil {
		logger.ErrorWritingFile()
		return nil
	}

	return nil
}