package todoCMD

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var TodoCMD = &cobra.Command{
	Use:   "todo",
	Short: "Manage to-do list",
	Long:  "Add, list or delete in your to-do list",
	Run: func(c *cobra.Command, args []string) {

		if c.Flags().Changed("add") || list || c.Flags().Changed("delete") {

			// Checking if add flag used
			if c.Flags().Changed("add") {
				if strings.HasPrefix(add, "-") {
					fmt.Println("Invalid task provided.")
					return
				}
				if add == "" {
					fmt.Println("No task provided")
					return
				}
				addingTask()
			}

			// checking if delete flag is used
			if c.Flags().Changed("delete") {
				if delete == 0 {
					fmt.Println("Invalid task number. Check task list.")
					return
				}
				deleteTask()
			}

			// checking if list flag is used
			if list {
				listTasks()
			}

		} else { // If no flags then run TUI

			p := tea.NewProgram(initialModel())

			if _, err := p.Run(); err != nil {
				fmt.Printf("Error: Something went wrong! \n%v\n", err)
				return
			}
		}

	},
}

var list bool
var add string
var delete int

func init() {
	TodoCMD.Flags().BoolVarP(&list, "list", "l", false, "Show only list of tasks")
	TodoCMD.Flags().StringVarP(&add, "add", "a", "", "Add tasks to the list")
	TodoCMD.Flags().IntVarP(&delete, "delete", "d", len(tasks), "Delete tasks from the list")
}
