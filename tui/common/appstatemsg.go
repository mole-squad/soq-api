package common

import (
	tea "github.com/charmbracelet/bubbletea"
)

type AppStateMsg struct {
	NewState AppState
}

func AppStateCmd(newState AppState) tea.Cmd {
	return func() tea.Msg {
		return AppStateMsg{NewState: newState}
	}
}
