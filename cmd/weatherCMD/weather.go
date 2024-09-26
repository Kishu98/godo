package weatherCMD

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

const apiUrl = "https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric"

type WeatherData struct {
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Name string `json:"name"`
}

type model struct {
	city string
}

var WeatherCMD = &cobra.Command{
	Use:   "weather [city]",
	Short: "Get the current weather for specified city",

	Run: func(c *cobra.Command, args []string) {
		if len(args) == 0 {
			m := model{
				city: "",
			}
			p := tea.NewProgram(m)
			if _, err := p.Run(); err != nil {
				fmt.Println("error", err)
				return
			}
		} else {
			city := args[0]
			getWeather(city)
		}
	},
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			getWeather(m.city)
			return m, tea.Quit
		case "backspace":
			m.city = m.city[:len(m.city)-1]
		case "q", "ctrl+c":
			return m, tea.Quit
		default:
			m.city += msg.String()
		}
	}
	return m, nil
}

func (m model) View() string {

	s := fmt.Sprintf("Enter the city: %v", m.city)

	return s
}

func getWeather(city string) {

	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Error loading environment variables", err)
		return
	}

	var apiKey = os.Getenv("API_KEY")
	url := fmt.Sprintf(apiUrl, city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching weather data:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var weatherData WeatherData
	if err := json.Unmarshal(body, &weatherData); err != nil {
		fmt.Println("Error parsing weather data:", err)
		return
	}

	fmt.Printf("Weather in %s: %s, %g°C\n", weatherData.Name, weatherData.Weather[0].Description, weatherData.Main.Temp)
}
