package list

import (
	"fmt"

	db "github.com/keeles/hours/internal/database"
	"github.com/keeles/hours/internal/lib"
	"github.com/keeles/hours/internal/logger"

	"github.com/alecthomas/kong"
)

func (o Options) Run(ctx *kong.Context) error {
	data, err := db.GetAll()
	if err != nil {
		fmt.Printf("Error: %v \n", err)
		logger.FileNotExists()
		return nil
	}

	for client, tasks := range data {
		fmt.Printf("%v: \n", client)
		for task, minutes := range tasks {
			hours := lib.MinutesToRoundedHours(minutes)
			fmt.Printf("  %v: %d minutes (%.2f hours)\n", task, minutes, hours)
		}
		fmt.Println("")
	}

	return nil
}
