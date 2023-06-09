package tui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mrclrchtr/gh-prTools/git"
	"log"
)

type Model struct {
	commits []string
	cursor  int
}

func InitialModel() Model {
	commits, err := git.GetPrCommits()
	if err != nil {
		log.Fatal("not a git repository", err)
	}

	return Model{
		commits: commits,
	}
}

func (m Model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		//case "up", "k":
		//	if m.cursor > 0 {
		//		m.cursor--
		//	}
		//
		//// The "down" and "j" keys move the cursor down
		//case "down", "j":
		//	if m.cursor < len(m.choices)-1 {
		//		m.cursor++
		//	}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			//_, ok := m.selected[m.cursor]
			//if ok {
			//	delete(m.selected, m.cursor)
			//} else {
			//	m.selected[m.cursor] = struct{}{}
			//}
		}
	}

	// Return the updated Model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m Model) View() string {
	// The header
	s := "What do you want to do?\n\n"

	for _, message := range m.commits {
		s += fmt.Sprintf("%s\n", message)
	}

	//// Iterate over our choices
	//for i, command := range m.choices {
	//
	//	// Is the cursor pointing at this command?
	//	cursor := " " // no cursor
	//	if m.cursor == i {
	//		cursor = ">" // cursor!
	//	}
	//
	//	// Render the row
	//	s += fmt.Sprintf("%s %s\n", cursor, command)
	//}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}
