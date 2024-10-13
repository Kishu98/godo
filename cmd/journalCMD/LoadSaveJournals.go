package journalCMD

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Journal struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
}

var journals []Journal

const journalFile = "Journals.json"

func LoadJournals() error {
	file, err := os.ReadFile(journalFile)
	if err != nil {
		if os.IsNotExist(err) {
			journals = []Journal{}
			return nil
		}
		return err
	}
	return json.Unmarshal(file, &journals)
}

func SaveJournals() error {
	data, err := json.Marshal(journals)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}
	return os.WriteFile(journalFile, data, 0644)
}
