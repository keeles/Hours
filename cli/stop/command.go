package stop

import (
	"fmt"

	"github.com/alecthomas/kong"
	db "github.com/keeles/hours/v2/internal/database"
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

	client, task, err := db.StopTimer()
	if err != nil {
		fmt.Printf("Error stopping timer: %s\n", err)
		return nil
	}

	fmt.Printf("Timer for %s has been stopped. Time allocated to %s\n", client, task)

	return nil
}
