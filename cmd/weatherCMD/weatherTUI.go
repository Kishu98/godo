package weatherCMD

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	city string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if m.city != "" {
				getWeather(m.city)
				return m, tea.Quit
			}
		case "backspace":
			if len(m.city) > 0 {
				m.city = m.city[:len(m.city)-1]
			}
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
