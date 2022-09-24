package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices  []string         // items in the to-do list
	cursor   int              // which item is our cursor pointing at
	selected map[int]struct{} // which items are selected
}

func initialModel() model {
	return model{
		// Our shopping list is a grocery list
		choices: []string{"buy carrots", "Buy celery", "Buy tomato"},

		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}
