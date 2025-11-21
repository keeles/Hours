package version

import (
	"github.com/keeles/hours/internal/logger"

	"github.com/alecthomas/kong"
)

func (o Options) Run (ctx * kong.Context) error {
	current := ctx.Model.Vars()["versionNumber"]
	logger.PrintVersion(current)
	return nil
}