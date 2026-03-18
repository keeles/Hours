package time

import (
	"fmt"

	"github.com/alecthomas/kong"
	db "github.com/keeles/hours/v2/internal/database"
)

func (o Options) Run(ctx *kong.Context) error {
	timer, exists, err := db.GetTimer()
	if err != nil {
		fmt.Println("Error connecting to database")
		return nil
	}

	if !exists {
		fmt.Println("No timer is currently running")
		return nil
	}

	fmt.Printf("Current Timer: %s: %s", timer.ClientName, timer.StartTime)
	return nil
}
