package journalCMD

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	Journals    []Journal
	cursor      int
	text_cursor int
	JournalMode string
	Journal     *Journal
	newJournal  Journal
}

func initialModel() model {

	if err := LoadJournals(); err != nil {
		fmt.Printf("Error loading Journals \n%v\n", err)
		return model{
			Journals:    []Journal{},
			JournalMode: "menu",
		}
	}

	return model{
		Journals:    journals,
		JournalMode: "menu",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.JournalMode {
		case "menu":
			switch msg.String() {
			case "a":
				m.JournalMode = "addJournal"
				m.newJournal.Title = ""
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down", "j":
				if m.cursor < len(m.Journals)-1 {
					m.cursor++
				}
			case "d":
				if len(m.Journals) != 0 {
					m.Journals = append(m.Journals[:m.cursor], m.Journals[m.cursor+1:]...)
					if m.cursor >= len(m.Journals) {
						m.cursor = len(m.Journals) - 1
					}
					if m.cursor < 0 {
						m.cursor = 0
					}
				}
			case "e":
				m.JournalMode = "editJournalTitle"
				m.Journal = &m.Journals[m.cursor]
			case "enter":
				m.JournalMode = "editJournalBody"
				m.Journal = &m.Journals[m.cursor]
			case "q", "ctrl+c":
				journals = m.Journals
				if err := SaveJournals(); err != nil {
					fmt.Printf("Error saving to file: %v\n", err)
				}
				return m, tea.Quit
			}
		case "addJournal":
			switch msg.String() {
			case "esc":
				m.JournalMode = "menu"
			case "backspace":
				if len(m.newJournal.Title) > 0 {
					m.newJournal.Title = m.newJournal.Title[:len(m.newJournal.Title)-1]
				}
			case "enter":
				m.JournalMode = "journalBody"
				m.newJournal.Body = ""
				// tasks = m.Journals
				// SaveTasks()
				// m.JournalMode = false
			// case "left":
			// 	if m.text_cursor > 0 {
			// 		m.text_cursor--
			// 	}
			// case "right":
			// 	if m.text_cursor < len(m.Journals[m.cursor].Body)-1 {
			// 		m.text_cursor++
			// 	}
			default:
				m.newJournal.Title += msg.String()
			}
		case "journalBody":
			switch msg.String() {
			case "esc":
				m.JournalMode = "menu"
			case "backspace":
				if len(m.newJournal.Body) > 0 {
					m.newJournal.Body = m.newJournal.Body[:len(m.newJournal.Body)-1]
				}
			case "enter":
				m.JournalMode = "menu"
				m.Journals = append(m.Journals, Journal{ID: len(m.Journals) + 1, Title: m.newJournal.Title, Body: m.newJournal.Body, CreatedAt: time.Now()})
			case "alt+enter":
				m.newJournal.Body += "\n"
			default:
				m.newJournal.Body += msg.String()
			}
		case "editJournalBody":
			switch msg.String() {
			case "esc":
				m.JournalMode = "menu"
			case "backspace":
				if len(m.Journal.Body) > 0 {
					m.Journal.Body = m.Journal.Body[:len(m.Journal.Body)-1]
				}
			case "enter":
				m.JournalMode = "menu"
				// m.Journals = append(m.Journals, Journal{ID: len(m.Journals) + 1, Title: m.newJournal.Title, Body: m.newJournal.Body, CreatedAt: time.Now()})
				m.Journals[m.cursor].Body = m.Journal.Body
			case "alt+enter":
				m.Journal.Body += "\n"
			default:
				m.Journal.Body += msg.String()
			}
		case "editJournalTitle":
			switch msg.String() {
			case "esc":
				m.JournalMode = "menu"
			case "backspace":
				if len(m.Journal.Title) > 0 {
					m.Journal.Title = m.Journal.Title[:len(m.Journal.Title)-1]
				}
			case "enter":
				m.JournalMode = "menu"
				// m.Journals = append(m.Journals, Journal{ID: len(m.Journals) + 1, Title: m.newJournal.Title, Body: m.newJournal.Body, CreatedAt: time.Now()})
				m.Journals[m.cursor].Title = m.Journal.Title
			default:
				m.Journal.Title += msg.String()
			}

		}
	}
	return m, nil
}

func (m model) View() string {
	var s string

	if m.JournalMode == "addJournal" {
		s = fmt.Sprintf("Enter the title: \n%s\n", m.newJournal.Title)
	} else if m.JournalMode == "journalBody" {
		s = fmt.Sprintf("Write Journal: \n%s\n", m.newJournal.Body)
	} else if m.JournalMode == "editJournalBody" {
		s = fmt.Sprintf("Edit Journal: \n%s\n", m.Journal.Body)
	} else if m.JournalMode == "editJournalTitle" {
		s = fmt.Sprintf("Edit Journal Title: \n%s\n", m.Journal.Title)
	} else {
		if len(m.Journals) == 0 {
			s = fmt.Sprintln("List is empty")
			s += fmt.Sprintln("\nPress 'a' to add a Journal Entry.\nPress 'q' to quit.")
			return s
		}
		s = fmt.Sprintln("Journals:\n")

		for i, journal := range m.Journals {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}

			s += fmt.Sprintf("\t %s %d. %s\t\t\t\t%d-%s-%d\n", cursor, journal.ID, journal.Title, journal.CreatedAt.Day(), journal.CreatedAt.Month(), journal.CreatedAt.Year())
		}
		s += fmt.Sprintln("\n\nPress 'a' to add a Journal Entry.\nPress 'e' to edit journal entry title.\nPress 'enter' to edit the jounal entry.\nPress 'd' to delete task.\nPress 'q' to quit.")
	}
	return s
}
