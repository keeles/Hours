package version

import (
	"github.com/keeles/hours/v2/internal/logger"

	"github.com/alecthomas/kong"
)

func (o Options) Run(ctx *kong.Context) error {
	current := ctx.Model.Vars()["versionNumber"]
	logger.PrintVersion(current)

	return nil
}
