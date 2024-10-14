package journalCMD

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	Journals    []Journal
	cursor      int
	JournalMode int
	newJournal  Journal
}

func initialModel() model {

	if err := LoadJournals(); err != nil {
		fmt.Printf("Error loading Journals \n%v\n", err)
		return model{
			Journals: []Journal{},
		}
	}

	return model{
		Journals: journals,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.JournalMode {
		case 0:
			switch msg.String() {
			case "a":
				m.JournalMode = 1
				m.newJournal.Title = ""
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down", "j":
				if m.cursor < len(m.Journals)-1 {
					m.cursor++
				}
			case "q":
				journals = m.Journals
				if err := SaveJournals(); err != nil {
					fmt.Printf("Error saving to file: %v\n", err)
				}
				return m, tea.Quit
			}
		case 1:
			switch msg.String() {
			case "esc":
				m.JournalMode = 0
			case "backspace":
				if len(m.newJournal.Title) > 0 {
					m.newJournal.Title = m.newJournal.Title[:len(m.newJournal.Title)-1]
				}
			case "enter":
				m.JournalMode = 2
				m.newJournal.Body = ""
				// tasks = m.Journals
				// SaveTasks()
				// m.JournalMode = false
			default:
				m.newJournal.Title += msg.String()
			}
		case 2:
			switch msg.String() {
			case "esc":
				m.JournalMode = 0
			case "backspace":
				if len(m.newJournal.Body) > 0 {
					m.newJournal.Body = m.newJournal.Body[:len(m.newJournal.Body)-1]
				}
			case "enter":
				m.JournalMode = 0
				m.Journals = append(m.Journals, Journal{ID: len(m.Journals) + 1, Title: m.newJournal.Title, Body: m.newJournal.Body})
			default:
				m.newJournal.Body += msg.String()
			}

		}
	}
	return m, nil
}

func (m model) View() string {
	var s string

	if m.JournalMode == 1 {
		s = fmt.Sprintf("Enter the title: %s\n", m.newJournal.Title)
	} else if m.JournalMode == 2 {
		s = fmt.Sprintf("Write Journal: \n%s\n", m.newJournal.Body)
	} else {
		if len(m.Journals) == 0 {
			s = fmt.Sprintln("List is empty")
			return s
		}
		s = fmt.Sprintln("Journals:\n")

		for i, journal := range m.Journals {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}

			s += fmt.Sprintf("\t %s %s\n", cursor, journal.Title)
		}

	}
	return s
}
