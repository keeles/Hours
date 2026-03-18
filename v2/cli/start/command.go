package start

import (
	"fmt"

	"github.com/alecthomas/kong"
	db "github.com/keeles/hours/v2/internal/database"
	"github.com/keeles/hours/v2/internal/logger"
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

	err = db.StartTimer(o.Client, o.Task)
	if err != nil {
		fmt.Printf("Error starting timer: %s \n", err)
		return nil
	}

	return nil
}
