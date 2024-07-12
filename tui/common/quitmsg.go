package common

import tea "github.com/charmbracelet/bubbletea"

type QuitMsg struct{}

func NewQuitMsg() tea.Msg {

	return QuitMsg{}
}
