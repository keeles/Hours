package add

import (
	"github.com/keeles/hours/internal/lib"
	"github.com/keeles/hours/internal/logger"

	"github.com/alecthomas/kong"
)

func (o Options) Run(ctx *kong.Context) error {
	err := lib.UpdateTaskHours(o.Name, o.Task, o.NewHours, false)
	if err != nil {
		logger.ErrorWritingFile()
		return nil
	}

	return nil
}
