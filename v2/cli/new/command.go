package new

import (
	"fmt"

	db "github.com/keeles/hours/v2/internal/database"

	"github.com/alecthomas/kong"
)

func (o Options) Run(ctx *kong.Context) error {
	err := db.AddNewClient(o.Name)
	if err != nil {
		return err
	}

	fmt.Printf("New Project Created: %s\n", o.Name)
	return nil
}
