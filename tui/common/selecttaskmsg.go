package common

import (
	"github.com/burkel24/task-app/pkg/tasks"
	tea "github.com/charmbracelet/bubbletea"
)

type SelectTaskMsg struct {
	Task tasks.TaskDTO
}

func SelectTaskCmd(task tasks.TaskDTO) tea.Cmd {
	return func() tea.Msg {
		return SelectTaskMsg{Task: task}
	}
}

func NewTaskCmd() tea.Cmd {
	return func() tea.Msg {
		newTask := tasks.TaskDTO{
			Summary: "",
			Notes:   "",
		}

		return SelectTaskMsg{Task: newTask}
	}
}
