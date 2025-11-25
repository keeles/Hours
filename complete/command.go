package complete

import (
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/keeles/hours/internal/lib"
	"github.com/keeles/hours/internal/logger"
)

func (o Options) Run(ctx *kong.Context) error {
	err := lib.DeleteTask(o.Client, o.Task)
	if err != nil {
		logger.ProjectNotFound(o.Task)
		return nil
	}

	fmt.Printf("Deleted Task: %v \n", o.Task)
	return nil
}
