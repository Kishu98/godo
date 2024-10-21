package newsCMD

import (
	"github.com/spf13/cobra"
)

var NewsCMD = &cobra.Command{
	Use:   "news",
	Short: "Get the latest news and articles",
	Run: func(cmd *cobra.Command, args []string) {
		getNews()
	},
}
