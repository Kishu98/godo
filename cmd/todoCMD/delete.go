package todoCMD

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCMD = &cobra.Command{
	Use:   "delete",
	Short: "Delete any task from the to-do list",
	Run: func(c *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Println("No task provided")
			return
		}

		if err := LoadTasks(); err != nil {
			fmt.Printf("Error: Loading tasks \n%v\n", err)
			return
		}

		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Error: Invalid task ID \n%v\n", err)
			return
		}

		for i, task := range tasks {
			if task.ID == taskID {
				tasks = append(tasks[:i], tasks[i+1:]...)
				if err := SaveTasks(); err != nil {
					fmt.Printf("Error: Saving to file \n%v\n", err)
					return
				}
				fmt.Printf("Deleted task: %d\n", taskID)
				return
			}
		}

		fmt.Printf("Task ID %d not found \n", taskID)

	},
}
