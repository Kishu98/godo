package todoCMD

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	todos     []Task
	cursor    int
	inputMode bool
	newTodo   string
}

func initialModel() model {

	if err := LoadTasks(); err != nil {
		fmt.Printf("Error loading tasks \n%v\n", err)
		return model{
			todos: []Task{},
		}
	}

	return model{
		todos: tasks,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.inputMode {
		case false:
			switch msg.String() {
			case "a":
				m.inputMode = true
				m.newTodo = ""
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down", "j":
				if m.cursor < len(m.todos)-1 {
					m.cursor++
				}
			case "d":
				if len(m.todos) != 0 {
					m.todos = append(m.todos[:m.cursor], m.todos[m.cursor+1:]...)
					if m.cursor >= len(m.todos) {
						m.cursor = len(m.todos) - 1
					}
					tasks = m.todos
					if err := SaveTasks(); err != nil {
						fmt.Printf("Error saving to file: %v\n", err)
					}
				}
			case "q", "ctrl+c":
				return m, tea.Quit
			}
		case true:
			switch msg.String() {
			case "esc":
				m.inputMode = false
			case "backspace":
				if len(m.newTodo) > 0 {
					m.newTodo = m.newTodo[:len(m.newTodo)-1]
				}
			case "enter":
				m.todos = append(m.todos, Task{ID: len(m.todos) + 1, Name: m.newTodo})
				tasks = m.todos
				SaveTasks()
				m.inputMode = false
			default:
				m.newTodo += msg.String()
			}

		}
	}
	return m, nil
}

func (m model) View() string {
	var s string

	if m.inputMode {
		s = fmt.Sprintf("Enter task to add: %s\n", m.newTodo)
	} else {

		if len(m.todos) == 0 {
			s = fmt.Sprintln("List is empty. Press 'q' to exit")
			return s
		}
		s = fmt.Sprintln("Your list:")

		for i, todo := range m.todos {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}

			s += fmt.Sprintf("%s %s\n", cursor, todo.Name)
		}

		s += fmt.Sprintln("\n\nPress 'q' to quit.")

	}
	return s
}
