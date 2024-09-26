package todoCMD

import "fmt"

// import (
// 	"fmt"

// 	"github.com/spf13/cobra"
// )

// var addCMD = &cobra.Command{
// 	Use:   "add",
// 	Short: "Add a new task to the tO-do list",
// 	Run: func(c *cobra.Command, args []string) {
// 		if len(args) == 0 {
// 			fmt.Println("Error: No task provided")
// 			return
// 		}

// 		task := args[0]

// 		if err := LoadTasks(); err != nil {
// 			fmt.Printf("Error: Loading tasks\n%v\n", err)
// 			return
// 		}

// 		newTask := Task{
// 			ID:   len(tasks) + 1,
// 			Name: task,
// 		}

// 		tasks = append(tasks, newTask)
// 		if err := SaveTasks(); err != nil {
// 			fmt.Printf("Error: Saving to file\n%v\n", err)
// 			return
// 		}

// 		fmt.Printf("Task added: %s\n", task)
// 	},
// }

func addingTask() {
	if len(add) == 0 {
		fmt.Println("No task provided to add to list")
		return
	}

	if err := LoadTasks(); err != nil {
		fmt.Printf("Error loading tasks %v\n", err)
		return
	}

	newTask := Task{
		ID:   len(tasks) + 1,
		Name: add,
	}

	tasks = append(tasks, newTask)
	if err := SaveTasks(); err != nil {
		fmt.Printf("Error adding task: %v\n", err)
		return
	}

	fmt.Printf("Task added: %s\n", add)
}
