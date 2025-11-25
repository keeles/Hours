package get

import (
	"fmt"

	"github.com/keeles/hours/internal/lib"
	"github.com/keeles/hours/internal/logger"

	"github.com/alecthomas/kong"
)

func (o Options) Run(ctx *kong.Context) error {
	tasks, err := lib.GetClientTasks(o.Name)
	if err != nil {
		logger.ProjectNotFound(o.Name)
		return nil
	}
	for name, hours := range tasks {
		fmt.Printf("%v: %d hours\n", name, hours)
	}
	return nil
}
