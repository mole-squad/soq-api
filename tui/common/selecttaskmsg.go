package common

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mole-squad/soq-api/pkg/tasks"
)

type SelectTaskMsg struct {
	Task tasks.TaskDTO
}

func NewSelectTaskMsg(task tasks.TaskDTO) tea.Cmd {
	return func() tea.Msg {
		return SelectTaskMsg{Task: task}
	}
}
