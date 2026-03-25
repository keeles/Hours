package time

import (
	"fmt"
	"time"

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

	endTime := time.Now().UTC()
	duration := endTime.Sub(timer.StartTime)
	minutes := int(duration.Minutes())

	fmt.Printf("Current Timer \n	Client: %s \n	Running for: %d minutes \n", timer.ClientName, minutes)
	if timer.TaskName != nil {
		fmt.Printf("Task: %s \n", *timer.TaskName)

	}
	return nil
}
