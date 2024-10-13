package journalCMD

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	journals    []Journal
	cursor      int
	journalMode int
	newJournal  string
}

func initialModel() model {

	if err := LoadJournals(); err != nil {
		fmt.Printf("Error loading journals \n%v\n", err)
		return model{
			journals: []Journal{},
		}
	}

	return model{
		journals: journals,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.journalMode {
		case 0:
			switch msg.String() {
			case "a":
				m.journalMode = 1
				m.newJournal = ""
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down", "j":
				if m.cursor < len(m.journals)-1 {
					m.cursor++
				}
			}
		}
	}
}
