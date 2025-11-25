package new

import (
	"fmt"

	"github.com/keeles/hours/internal/lib"

	"github.com/alecthomas/kong"
)

func (o Options) Run(ctx *kong.Context) error {
	err := lib.AddNewClient(o.Name)
	if err != nil {
		return err
	}

	fmt.Printf("New Project Created: %s\n", o.Name)
	return nil
}
