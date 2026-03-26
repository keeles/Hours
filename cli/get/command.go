package get

import (
	"fmt"

	db "github.com/keeles/hours/internal/database"
	"github.com/keeles/hours/internal/lib"
	"github.com/keeles/hours/internal/logger"

	"github.com/alecthomas/kong"
)

func (o Options) Run(ctx *kong.Context) error {
	tasks, err := db.GetClientTasks(o.Name)
	if err != nil {
		logger.ProjectNotFound(o.Name)
		return nil
	}

	for name, minutes := range tasks {
		hours := lib.MinutesToRoundedHours(minutes)
		fmt.Printf("%v: %d minutes (%.2f hours)\n", name, minutes, hours)
	}

	return nil
}
