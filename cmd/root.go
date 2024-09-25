package cmd

import (
	"fmt"

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
		fmt.Printf("Error: %v", err)
		return
	}
}
