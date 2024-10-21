package newsCMD

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const apiURL = "https://newsapi.org/v2/top-headlines?category=technology&apiKey=7a5a6dcd9d5d44ebb6ec9d425a73e3cc"

type News struct {
	Status   string
	Articles []struct {
		Author  string
		Title   string
		Url     string
		Content string
	}
}

func getNews() {

	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error fetching News data:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error fetching News data:", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var news News
	if err := json.Unmarshal(body, &news); err != nil {
		fmt.Println("Error parsing news data:", err)
		return
	}

	for i, article := range news.Articles {
		fmt.Printf("%d. %s\n%s\n\n", i+1, article.Title, article.Url)
	}
}
