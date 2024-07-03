package app

import (
	"github.com/burkel24/task-app/tui/tasklist"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	AppStateTaskList = iota
)

type AppModel struct {
	quitting bool
	taskList tasklist.TaskListModel
	appState int
}

func NewAppModel() AppModel {
	return AppModel{
		appState: AppStateTaskList,
		taskList: tasklist.NewTaskListModel(),
	}
}

func (m AppModel) Init() tea.Cmd {
	return tea.Batch(
		m.taskList.Init(),
	)
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		}
	}

	switch m.appState {
	case AppStateTaskList:
		return m.taskList.Update(msg)
	}

	return m, nil
}

func (m AppModel) View() string {
	if m.quitting {
		return "Bye!\n"
	}

	switch m.appState {
	case AppStateTaskList:
		return m.taskList.View()
	}

	return "No state"
}
