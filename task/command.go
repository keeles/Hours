package task

import (
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/keeles/hours/internal/lib"
	"github.com/keeles/hours/internal/logger"
)

func (o Options) Run(ctx *kong.Context) error {
	err := lib.AddNewTask(o.Client, o.Task)
	if err != nil {
		fmt.Printf("%v", err)
		logger.ErrorWritingFile()
		return nil
	}

	return nil
}
