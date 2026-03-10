package ui

import (

	tea "github.com/charmbracelet/bubbletea"
)

func (m SystemModel) Init() tea.Cmd {
	return nil
}

func (m SystemModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			m.Quitting = true
			return m, tea.Quit
		}
	}

	return m, nil
}

