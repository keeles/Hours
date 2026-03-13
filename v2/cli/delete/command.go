package delete

import (
	"fmt"

	"github.com/alecthomas/kong"
	db "github.com/keeles/hours/v2/internal/database"
	"github.com/keeles/hours/v2/internal/logger"
)

func (o Options) Run(ctx *kong.Context) error {
	err := db.DeleteClient(o.Name)
	if err != nil {
		logger.ErrorWritingFile()
		return nil
	}

	fmt.Printf("Deleted Client: %v \n", o.Name)
	return nil
}
