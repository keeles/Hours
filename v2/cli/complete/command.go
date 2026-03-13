package complete

import (
	"fmt"

	"github.com/alecthomas/kong"
	db "github.com/keeles/hours/v2/internal/database"
	"github.com/keeles/hours/v2/internal/logger"
)

func (o Options) Run(ctx *kong.Context) error {
	err := db.DeleteTask(o.Client, o.Task)
	if err != nil {
		logger.ProjectNotFound(o.Task)
		return nil
	}

	fmt.Printf("Deleted Task: %v \n", o.Task)
	return nil
}
