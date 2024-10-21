package cmd

import (
	"fmt"

	"github.com/Kishu98/godo/cmd/journalCMD"
	"github.com/Kishu98/godo/cmd/newsCMD"
	"github.com/Kishu98/godo/cmd/todoCMD"
	"github.com/Kishu98/godo/cmd/weatherCMD"
	"github.com/spf13/cobra"
)

var RootCMD = &cobra.Command{
	Use:   "godo",
	Short: "It will do everything & anything (Hopefully)",
	Long:  "A CLI app which is able to perform any task, made using go.",
	Run: func(c *cobra.Command, args []string) {
		fmt.Println("Checking godo")
	},
}

func Execute() {
	if err := RootCMD.Execute(); err != nil {
		fmt.Printf("Error: %v \n", err)
		return
	}
}

func init() {
	RootCMD.AddCommand(todoCMD.TodoCMD)
	RootCMD.AddCommand(weatherCMD.WeatherCMD)
	RootCMD.AddCommand(journalCMD.JournalCMD)
    RootCMD.AddCommand(newsCMD.NewsCMD)
}
