package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var tasks []Task

const taskFile = "tasks.json"

func loadTasks() error {
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

func saveTasks() error {
	data, err := json.Marshal(tasks)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return err
	}
	return os.WriteFile(taskFile, data, 0644)
}

var todoCMD = &cobra.Command{
	Use:   "todo",
	Short: "Manage to-do list",
	Long:  "Add, list or delete in your to-do list",
	Run: func(c *cobra.Command, args []string) {
		fmt.Println("To-do list command: Use 'add' , 'list', 'delete'.")
	},
}

var addCMD = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to the tO-do list",
	Run: func(c *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: No task provided")
			return
		}

		task := args[0]

		if err := loadTasks(); err != nil {
			fmt.Printf("Error: Loading tasks \n %v", err)
			return
		}

		newTask := Task{
			ID:   len(tasks) + 1,
			Name: task,
		}

		tasks = append(tasks, newTask)
		if err := saveTasks(); err != nil {
			fmt.Printf("Error: Saving to file \n %v", err)
			return
		}

		fmt.Printf("Task added: %s\n", task)
	},
}

var listCMD = &cobra.Command{
	Use:   "list",
	Short: "List all the commands in to-do list",
	Run: func(c *cobra.Command, args []string) {

		if err := loadTasks(); err != nil {
			fmt.Printf("Error: Loading tasks \n %v", err)
			return
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks in list")
			return
		}

		fmt.Println("Tasks:")
		for _, task := range tasks {
			fmt.Printf("%d %s \n", task.ID, task.Name)
		}
	},
}

var deleteCMD = &cobra.Command{
	Use:   "delete",
	Short: "Delete any task from the to-do list",
	Run: func(c *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No task provided")
			return
		}
		fmt.Println("Delete cmd is work in progress")
	},
}

func init() {
	RootCMD.AddCommand(todoCMD)
	todoCMD.AddCommand(addCMD)
	todoCMD.AddCommand(listCMD)
	todoCMD.AddCommand(deleteCMD)
}
