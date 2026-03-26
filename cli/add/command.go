package add

import (
	db "github.com/keeles/hours/internal/database"
	"github.com/keeles/hours/internal/logger"

	"github.com/alecthomas/kong"
)

func (o Options) Run(ctx *kong.Context) error {
	amount := o.Amount

	if o.Hours {
		amount = amount * 60
	}

	err := db.UpdateTaskMinutes(o.Name, o.Task, amount, false)

	if err != nil {
		logger.ErrorWritingFile()
		return nil
	}

	return nil
}
