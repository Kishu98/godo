package todoCMD

import (
	"fmt"

	"github.com/spf13/cobra"
)

var tasks []Task

var TodoCMD = &cobra.Command{
	Use:   "todo",
	Short: "Manage to-do list",
	Long:  "Add, list or delete in your to-do list",
	Run: func(c *cobra.Command, args []string) {
		fmt.Println("To-do list command: Use 'add' , 'list', 'delete'.")
	},
}

func init() {
	TodoCMD.AddCommand(deleteCMD)
	TodoCMD.AddCommand(addCMD)
	TodoCMD.AddCommand(listCMD)
}
