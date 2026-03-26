package new

import (
	"fmt"

	db "github.com/keeles/hours/internal/database"

	"github.com/alecthomas/kong"
)

func (o Options) Run(ctx *kong.Context) error {
	err := db.AddNewClient(o.Name)
	if err != nil {
		return err
	}

	fmt.Printf("New Client/Category Created: %s\n", o.Name)

	return nil
}
