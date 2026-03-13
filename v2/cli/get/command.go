package get

import (
	"fmt"

	db "github.com/keeles/hours/v2/internal/database"
	"github.com/keeles/hours/v2/internal/logger"

	"github.com/alecthomas/kong"
)

func (o Options) Run(ctx *kong.Context) error {
	tasks, err := db.GetClientTasks(o.Name)
	if err != nil {
		logger.ProjectNotFound(o.Name)
		return nil
	}
	for name, hours := range tasks {
		fmt.Printf("%v: %d hours\n", name, hours)
	}
	return nil
}
