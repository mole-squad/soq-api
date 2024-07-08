package common

import tea "github.com/charmbracelet/bubbletea"

type RefreshFocusAreasMsg struct {
}

func NewRefreshFocusAreasMsg() tea.Cmd {
	return func() tea.Msg {
		return RefreshFocusAreasMsg{}
	}
}
