package weatherCMD

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const apiUrl = "https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=%s"

type WeatherData struct {
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Name string `json:"name"`
}

func getWeather(city string) {

	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Error loading environment variables", err)
		return
	}

	var apiKey = os.Getenv("API_KEY")
	url := fmt.Sprintf(apiUrl, city, apiKey, units)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching weather data:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error fetching weather data: %s\n", resp.Status)
		return
	}

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

	if units != "imperial" {
		fmt.Printf("Weather in %s: %s, %g°C\n", weatherData.Name, weatherData.Weather[0].Description, weatherData.Main.Temp)
	} else {
		fmt.Printf("Weather in %s: %s, %g°F\n", weatherData.Name, weatherData.Weather[0].Description, weatherData.Main.Temp)
	}
}
