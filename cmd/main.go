package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kong"
)

func main() {
	hours := &Hours{}
	ctx := kong.Parse(
		hours,
		kong.Description("A tool for tracking hours via the command line."),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact:             true,
			Summary:             false,
			NoExpandSubcommands: true,
		}),
		kong.Vars{
			"versionNumber": "1.0.1",
		},
	)
	if err := ctx.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
