package delete

import (
	"fmt"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/keeles/hours/internal/lib"
	"github.com/keeles/hours/internal/logger"
)

func (o Options) Run (ctx *kong.Context) error {
	data, err := lib.ReadFile()
	if err != nil {
		logger.FileNotExists()
		return nil
	}

	projectName := strings.ToLower(o.Name)
	project := data[strings.ToLower(projectName)]
	if (lib.Project{} == project) {
		logger.ProjectNotFound(o.Name)
		return nil
	}

	if !o.Force {
		fmt.Printf("Delete project %s with %d hours? (y/N): ", o.Name, project.Hours)
		var response string
		fmt.Scanln(&response)

		if strings.ToLower(response) != "y" {
			fmt.Println("Deletion cancelled")
			return nil
		}
	}
	delete(data, projectName)

	err = lib.WriteFile(data)
	if err != nil {
		logger.ErrorWritingFile()
		return nil
	}

	fmt.Printf("Deleted project: %v \n", o.Name)
	return nil
}