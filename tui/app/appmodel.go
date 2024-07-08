package app

import (
	"fmt"
	"log/slog"

	"github.com/burkel24/task-app/tui/api"
	"github.com/burkel24/task-app/tui/common"
	"github.com/burkel24/task-app/tui/taskform"
	"github.com/burkel24/task-app/tui/tasklist"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	errorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
)

type AppModel struct {
	appState common.AppState
	error
	taskForm taskform.TaskFormModel
	taskList tasklist.TaskListModel
	quitting bool
}

func NewAppModel() AppModel {
	client := api.NewClient()

	return AppModel{
		appState: common.AppStateTaskList,
		taskForm: taskform.NewTaskFormModel(client),
		taskList: tasklist.NewTaskListModel(client),
	}
}

func (m AppModel) Init() tea.Cmd {
	return tea.Batch(
		m.taskList.Init(),
		m.taskForm.Init(),
	)
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		listCmd tea.Cmd
		formCmd tea.Cmd
	)

	slog.Debug(fmt.Sprintf("AppModel.Update: %T", msg))

	switch msg := msg.(type) {

	case common.ErrorMsg:
		m.error = msg.Err

	case tea.KeyMsg:
		switch m.appState {

		case common.AppStateTaskList:
			m.taskList, listCmd = m.taskList.Update(msg)
			return m, listCmd

		case common.AppStateTaskForm:
			m.taskForm, formCmd = m.taskForm.Update(msg)
			return m, formCmd
		}

	case common.AppStateMsg:
		m.appState = msg.NewState
		return m, nil
	}

	m.taskList, listCmd = m.taskList.Update(msg)
	m.taskForm, formCmd = m.taskForm.Update(msg)

	return m, tea.Batch(listCmd, formCmd)
}

func (m AppModel) View() string {
	if m.quitting {
		return "Bye!\n"
	}

	if m.error != nil {
		return errorStyle.Render(fmt.Sprintf("Error: %s\n", m.error))
	}

	switch m.appState {
	case common.AppStateTaskList:
		return m.taskList.View()

	case common.AppStateTaskForm:
		return m.taskForm.View()
	}

	return "No state"
}
