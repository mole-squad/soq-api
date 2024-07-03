package main

import (
	"fmt"
	"os"

	"github.com/burkel24/task-app/tui/tasklist"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	AppStateTaskList = iota
)

type Model struct {
	quitting bool
	taskList tasklist.TaskListModel
	appState int
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.taskList.Init(),
	)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m Model) View() string {
	if m.quitting {
		return "Bye!\n"
	}

	switch m.appState {
	case AppStateTaskList:
		return m.taskList.View()
	}

	return "No state"
}

func main() {
	m := Model{
		appState: AppStateTaskList,
		taskList: tasklist.NewTaskListModel(),
	}

	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
