package weatherCMD

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

const apiUrl = "https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric"

type WeatherData struct {
	Weather []struct {
		Desciption string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Name string `json:"name"`
}

var WeatherCMD = &cobra.Command{
	Use:   "weather [city]",
	Short: "Get the current weather for specified city",
	Args:  cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		city := args[0]
		getWeather(city)
	},
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

	fmt.Printf("Weather in %s: %s, %g°C\n", weatherData.Name, weatherData.Weather[0].Desciption, weatherData.Main.Temp)
}
