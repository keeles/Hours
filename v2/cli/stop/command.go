package stop

import (
	"fmt"

	"github.com/alecthomas/kong"
	db "github.com/keeles/hours/v2/internal/database"
	"github.com/keeles/hours/v2/internal/lib"
)

func (o Options) Run(ctx *kong.Context) error {
	_, exists, err := db.GetTimer()
	if err != nil {
		fmt.Printf("Database connection error: %s", err)
		return nil
	}

	if !exists {
		fmt.Println("No timer is currently running")
		return nil
	}

	client := o.Client
	if client == "" {
		clientName, exists := lib.GetClientNameForCurrentDirectory()
		if !exists {
			fmt.Println("No client associated with current directory")
			return nil
		}
		client = clientName
	}

	err = db.StopTimer(client, o.Task)
	if err != nil {
		fmt.Print(err)
		return nil
	}

	fmt.Printf("Timer for %s has been stopped.", client)
	return nil
}
