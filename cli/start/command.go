package start

import (
	"fmt"

	"github.com/alecthomas/kong"
	db "github.com/keeles/hours/internal/database"
	"github.com/keeles/hours/internal/logger"
)

func (o Options) Run(ctx *kong.Context) error {
	timer, exists, err := db.GetTimer()
	if err != nil {
		fmt.Printf("Database connection error: %s", err)
		return nil
	}

	if exists {
		fmt.Println("Cannot have more that one timer running.")
		logger.PrintTimer(timer.ClientName, timer.StartTime, timer.TaskName)
		return nil
	}

	client, err := resolveClient(o.Client)
	if err != nil {
		fmt.Printf("Error Starting Timer, no client: %s", err)
		return nil
	}

	err = db.StartTimer(client, o.Task)
	if err != nil {
		fmt.Printf("Error starting timer: %s \n", err)
		return nil
	}

	fmt.Printf("Starting timer for %s \n", client)

	return nil
}
