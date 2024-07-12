package common

import tea "github.com/charmbracelet/bubbletea"

type AuthMsg struct {
	Token string
}

func NewAuthMsg(token string) tea.Cmd {
	return func() tea.Msg {
		return AuthMsg{Token: token}
	}
}
