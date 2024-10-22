package newsCMD

import (
	"fmt"
	// "io"
	// "net/http"

	"github.com/PuerkitoBio/goquery"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gocolly/colly/v2"
)

const apiURL = "https://dev.to/top/week"

// type News struct {
// 	Articles []struct {
// 		Author  string
// 		Title   string
// 		Url     string
// 		Content string
// 	}
// }

// Making the articles variables global so that I can use it everywhere and get the articles stored in it.
var articles []article

type model struct {
	Articles []article
	cursor   int
	ViewMode string
}

type article struct {
	Title string
	Link  string
}

func initialModel() model {
	if err := getNews(); err != nil {
		fmt.Printf("Error getting the articles: \n%v\n", err)
		return model{
			Articles: []article{},
			ViewMode: "list",
		}
	}

	return model{
		Articles: articles,
		ViewMode: "list",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.ViewMode {
		case "list":
			switch msg.String() {
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down", "j":
				if m.cursor < len(m.Articles)-1 {
					m.cursor++
				}
			case "q", "ctrl+c":
				return m, tea.Quit
			case "return", "enter":
				m.ViewMode = "read"
			}
		case "read":
			switch msg.String() {
			case "q":
				m.ViewMode = "list"
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	var s string

	if m.ViewMode == "list" {
		for i, article := range m.Articles {
			// if i > list {
			//
			// }
			cursor := " "
			// fmt.Println("Cursor -->", m.cursor)
			// fmt.Println("I -->", i)
			if m.cursor == i {
				cursor = ">"
			}

			s += fmt.Sprintf("%s %d. %s\n\t%s\n", cursor, i+1, article.Title, article.Link)
		}
	} else if m.ViewMode == "read" {
        s = getArticles(m)
	} else {
        s = fmt.Sprintln("No idea")
    }
    // why the getArticles not working????
	return s
}

func getArticles(m model) string {
	var s string
    fmt.Println("Starting to fetch...")
	c := colly.NewCollector()
	c.OnHTML("div.article-body", func(h *colly.HTMLElement) {
		h.DOM.Children().Each(func(i int, elem *goquery.Selection) {
			switch goquery.NodeName(elem) {
			case "p":
                fmt.Println("Para --> ", elem.Text())
				s += fmt.Sprint(elem.Text(), "\n")
            case "h1", "h2", "h3", "h4", "h5", "h6":
                fmt.Println("Para --> ", elem.Text())
                s += fmt.Sprint("\n\t", elem.Text(), "\n")
            default:
                fmt.Println("Para --> ", elem.Text())
                s += fmt.Sprint("\n\t\t", elem.Text(), "\n")
			}
		})
	})
    fmt.Print(s)
    // Did forget to add the c.Visit, but its still not working
    c.Visit(m.Articles[m.cursor].Link)
    return s
}

func getNews() error {

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
		if len(articles) == list {
			return
		}

		articles = append(articles, article)

	})

	c.Visit("https://dev.to/top/week")

	return nil

	// This below is to show the articles when not using tui.
	// for i, article := range articles {
	// 	if i >= list {
	// 		break
	// 	}
	// 	fmt.Printf("\n%d. %s\n%s\n", i+1, article.Title, article.Link)
	// }

	// var news News
	// if err := json.Unmarshal(body, &news); err != nil {
	// 	fmt.Println("Error parsing news data:", err)
	// 	return
	// }

	// for i, article := range news.Articles {
	// 	fmt.Printf("%d. %s\n%s\n\n", i+1, article.Title, article.Url)
	// }

}
