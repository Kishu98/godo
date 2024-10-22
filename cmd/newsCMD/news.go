package newsCMD

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var NewsCMD = &cobra.Command{
	Use:   "news",
	Short: "Get the latest news and articles",
	Run: func(cmd *cobra.Command, args []string) {
        p := tea.NewProgram(initialModel())

        if _, err:= p.Run(); err != nil {
            fmt.Printf("Error: Something went wrong! \n%v\n", err)
            return
        }
	},
}

var list int

func init() {
    NewsCMD.Flags().IntVarP(&list, "list", "l", 5, "List specified number of articles")
}

