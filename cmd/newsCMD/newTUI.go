package newsCMD

import (
	"fmt"
	// "io"
	// "net/http"

	"github.com/gocolly/colly/v2"
)

const apiURL = "https://dev.to/top/week"

type News struct {
	Articles []struct {
		Author  string
		Title   string
		Url     string
		Content string
	}
}

type article struct {
	Title string
	Link  string
}

func getNews() {

	// resp, err := http.Get(apiURL)
	// if err != nil {
	// 	fmt.Println("Error fetching News data:", err)
	// 	return
	// }
	// defer resp.Body.Close()
	//
	// if resp.StatusCode != http.StatusOK {
	// 	fmt.Println("Error fetching News data:", resp.StatusCode)
	// 	return
	// }

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("Error reading response:", err)
	// 	return
	// }

	// var articles []struct {
	//
	// }

	var articles []article

	c := colly.NewCollector()

	c.OnHTML("a.crayons-story__hidden-navigation-link", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		link = "https://dev.to" + link
		title := e.Text
		article := article{
			Title: title,
			Link:  link,
		}
		// article {
		//     Title: title,
		//     Link: link,
		// }
		articles = append(articles, article)

	})

	c.Visit("https://dev.to/top/week")

	for i, article := range articles {
        if i >= 5 {
            break
        }
		fmt.Printf("\n%d. %s\n%s\n", i+1, article.Title, article.Link)
	}

	// var news News
	// if err := json.Unmarshal(body, &news); err != nil {
	// 	fmt.Println("Error parsing news data:", err)
	// 	return
	// }

	// for i, article := range news.Articles {
	// 	fmt.Printf("%d. %s\n%s\n\n", i+1, article.Title, article.Url)
	// }

}
