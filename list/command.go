package list

import (
	"fmt"

	"github.com/keeles/hours/internal/lib"
	"github.com/keeles/hours/internal/logger"

	"github.com/alecthomas/kong"
)

func (o Options) Run(ctx *kong.Context) error {
	data, err := lib.GetAll()
	if err != nil {
		fmt.Printf("Error: %v \n", err)
		logger.FileNotExists()
		return nil
	}

	for client, tasks := range data {
		fmt.Printf("%v: \n", client)
		for task, hours := range tasks {
			fmt.Printf(" |- %v: %d hours\n", task, hours)
		}
	}
	return nil
}
