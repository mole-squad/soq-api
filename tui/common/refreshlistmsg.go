package common

import tea "github.com/charmbracelet/bubbletea"

type RefreshListMsg struct {
}

func NewRefreshListMsg() tea.Cmd {
	return func() tea.Msg {
		return RefreshListMsg{}
	}
}
