package task

import (
	"fmt"

	"github.com/alecthomas/kong"
	db "github.com/keeles/hours/internal/database"
	"github.com/keeles/hours/internal/logger"
)

func (o Options) Run(ctx *kong.Context) error {
	err := db.AddNewTask(o.Client, o.Task)
	if err != nil {
		fmt.Printf("%v", err)
		logger.ErrorWritingFile()
		return nil
	}

	return nil
}
