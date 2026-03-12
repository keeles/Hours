package remove

import (
	"fmt"

	"github.com/keeles/hours/internal/lib"
	"github.com/keeles/hours/internal/logger"

	"github.com/alecthomas/kong"
)

func (o Options) Run(ctx *kong.Context) error {
	err := lib.UpdateTaskHours(o.Name, o.Task, o.HoursToRemove, true)
	if err != nil {
		logger.ErrorWritingFile()
		return nil
	}

	fmt.Printf("Removed %v hours from %s", o.HoursToRemove, o.Task)
	return nil
}
