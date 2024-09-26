package weatherCMD

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var WeatherCMD = &cobra.Command{
	Use:   "weather",
	Short: "Get the current weather for specified city",
	Run: func(c *cobra.Command, args []string) {
		if len(place) == 0 {
			m := model{
				city: "",
			}
			p := tea.NewProgram(m)
			if _, err := p.Run(); err != nil {
				fmt.Println("error", err)
				return
			}
		} else {
			// city := args[0]
			getWeather(place)
		}
	},
}

var place string
var units string

func init() {
	WeatherCMD.Flags().StringVarP(&place, "place", "p", "", "Show weather of specific place")
	WeatherCMD.Flags().StringVarP(&units, "units", "u", "metric", "specify units")
}
