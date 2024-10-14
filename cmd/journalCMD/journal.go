package journalCMD

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var JournalCMD = &cobra.Command{
	Use:   "journal",
	Short: "Create Journals",
	Long:  "Create and write journals in terminal",
	Run: func(c *cobra.Command, args []string) {
		p := tea.NewProgram(initialModel())

		if _, err := p.Run(); err != nil {
			fmt.Printf("Error donkey: \n%v\n", err)
			return
		}
	},
}
