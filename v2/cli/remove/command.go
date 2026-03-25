package remove

import (
	"fmt"

	db "github.com/keeles/hours/v2/internal/database"
	"github.com/keeles/hours/v2/internal/logger"

	"github.com/alecthomas/kong"
)

func (o Options) Run(ctx *kong.Context) error {
	amount := o.Amount

	if o.Hours {
		amount = amount * 60
	}

	err := db.UpdateTaskMinutes(o.Name, o.Task, amount, true)
	if err != nil {
		fmt.Println(err)
		logger.ErrorWritingFile()
		return nil
	}

	fmt.Printf("Removed %v minutes from %s", amount, o.Task)
	return nil
}
