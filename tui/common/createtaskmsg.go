package common

import tea "github.com/charmbracelet/bubbletea"

type CreateTaskMsg struct {
}

func NewCreateTaskMsg() tea.Cmd {
	return func() tea.Msg {
		return CreateTaskMsg{}
	}
}
