package delete

import (
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/keeles/hours/internal/lib"
	"github.com/keeles/hours/internal/logger"
)

func (o Options) Run(ctx *kong.Context) error {
	err := lib.DeleteClient(o.Name)
	if err != nil {
		logger.ErrorWritingFile()
		return nil
	}

	fmt.Printf("Deleted Client: %v \n", o.Name)
	return nil
}
