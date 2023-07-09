package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// Defining the mode l

type model struct {
	choices  []string         //items in the list
	cursor   int              // cursor for choices
	selected map[int]struct{} // which item is selected
}

func initialModel() model {
	return model{
		choices:  []string{"Car ", "Motorbike ", "SpaceShip "},
		selected: map[int]struct{}{},
	}
}

// Initializing the model which is defined in L17

func (m model) Init() tea.Cmd {
	return nil
}

// Update messsage updates the ui in some way - the following code looks like its handling imput basically

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", "":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	return m, nil
}

// View funtion

func (m model) View() string {
	s := `What drive its
What is the main 
++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

`

	for i, choice := range m.choices {

		cursor := "  " // no cursor if not selected
		if m.cursor == i {
			cursor = "👉"
		}

		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "✅"
		}

		// rendering the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)

	}

	s += "\n Press q to quits.\n"
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Panic: %v\n", err)
		os.Exit(1)
	}
}
