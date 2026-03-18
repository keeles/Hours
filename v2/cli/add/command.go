package add

import (
	db "github.com/keeles/hours/v2/internal/database"
	"github.com/keeles/hours/v2/internal/logger"

	"github.com/alecthomas/kong"
)

func (o Options) Run(ctx *kong.Context) error {
	err := db.UpdateTaskMinutes(o.Name, o.Task, o.NewMinutes, false)
	if err != nil {
		logger.ErrorWritingFile()
		return nil
	}

	return nil
}
