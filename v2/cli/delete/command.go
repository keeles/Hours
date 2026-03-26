package delete

import (
	"github.com/alecthomas/kong"
	db "github.com/keeles/hours/v2/internal/database"
	"github.com/keeles/hours/v2/internal/logger"
)

func (o Options) Run(ctx *kong.Context) error {
	err := db.DeleteClient(o.Name, o.Force)
	if err != nil {
		logger.ErrorWritingFile()
		return nil
	}

	return nil
}
