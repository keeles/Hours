package remove

import (
	"fmt"

	db "github.com/keeles/hours/v2/internal/database"
	"github.com/keeles/hours/v2/internal/logger"

	"github.com/alecthomas/kong"
)

func (o Options) Run(ctx *kong.Context) error {
	err := db.UpdateTaskMinutes(o.Name, o.Task, o.MinutesToRemove, true)
	if err != nil {
		logger.ErrorWritingFile()
		return nil
	}

	fmt.Printf("Removed %v minutes from %s", o.MinutesToRemove, o.Task)
	return nil
}
