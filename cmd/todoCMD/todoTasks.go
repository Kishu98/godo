package todoCMD

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

const taskFile = "tasks.json"

func LoadTasks() error {
	file, err := os.ReadFile(taskFile)
	if err != nil {
		if os.IsNotExist(err) {
			tasks = []Task{}
			return nil
		}
		return err
	}
	return json.Unmarshal(file, &tasks)
}

func SaveTasks() error {
	data, err := json.Marshal(tasks)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}
	return os.WriteFile(taskFile, data, 0644)
}