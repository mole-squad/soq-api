package common

import (
	"github.com/burkel24/task-app/pkg/tasks"
	tea "github.com/charmbracelet/bubbletea"
)

type SelectTaskMsg struct {
	Task tasks.TaskDTO
}

func NewSelectTaskMsg(task tasks.TaskDTO) tea.Cmd {
	return func() tea.Msg {
		return SelectTaskMsg{Task: task}
	}
}
