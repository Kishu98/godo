package todoCMD

import "fmt"

// import (
// 	"fmt"

// 	"github.com/spf13/cobra"
// )

// var listCMD = &cobra.Command{
// 	Use:   "list",
// 	Short: "List all the commands in to-do list",
// 	Run: func(c *cobra.Command, args []string) {

// 		if err := LoadTasks(); err != nil {
// 			fmt.Printf("Error: Loading tasks\n%v\n", err)
// 			return
// 		}

// 		if len(tasks) == 0 {
// 			fmt.Println("No tasks in list")
// 			return
// 		}

//			fmt.Println("Tasks:")
//			for _, task := range tasks {
//				fmt.Printf("%d %s \n", task.ID, task.Name)
//			}
//		},
//	}
func listTasks() {

	if err := LoadTasks(); err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks in the list")
		return
	}

	for i, task := range tasks {
		fmt.Printf("%d %s\n", i+1, task.Name)
	}
}
