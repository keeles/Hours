package start

import (
	db "github.com/keeles/hours/v2/internal/database"
	"github.com/keeles/hours/v2/internal/lib"
)

func resolveClient(input string) (string, error) {
	if exists, _ := db.ClientExists(input); exists {
		return input, nil
	}

	if clientName, exists := lib.GetClientNameForCurrentDirectory(); exists {
		return clientName, nil
	}

	selected, err := db.SelectClientForTimer()
	if err != nil {
		return "", nil
	}

	return selected, nil
}
