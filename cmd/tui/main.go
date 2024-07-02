package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/burkel24/task-app/pkg/tasks"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	docStyle  = lipgloss.NewStyle().Margin(1, 2)
)

const taskUrl = "http://localhost:3000/tasks"

type Model struct {
	quitting bool
	err      error
	taskList []tasks.TaskDTO
	list     list.Model
}

type taskMsg struct {
	tasks []tasks.TaskDTO
}

type errMsg struct {
	error
}

type item struct {
	task tasks.TaskDTO
}

func (i item) Title() string       { return i.task.Summary }
func (i item) Description() string { return "TODO" }
func (i item) FilterValue() string { return i.task.Summary }

func getTasks() tea.Msg {
	c := &http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := c.Get(taskUrl)
	if err != nil {
		return errMsg{error: err}
	}

	defer res.Body.Close()

	respBody, _ := io.ReadAll(res.Body)

	var tasksResp []tasks.TaskDTO
	if err = json.Unmarshal(respBody, &tasksResp); err != nil {
		return errMsg{error: err}
	}

	return taskMsg{tasks: tasksResp}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		getTasks,
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

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case taskMsg:
		m.taskList = msg.tasks
		newItems := make([]list.Item, len(m.taskList))

		for i, task := range m.taskList {
			newItems[i] = item{task: task}
		}

		m.list.SetItems(newItems)

	case errMsg:
		m.err = msg.error
		return m, nil
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if m.quitting {
		return "Bye!\n"
	}

	if m.err != nil {
		text := fmt.Sprintf("Something went wrong: %s", m.err.Error())
		return helpStyle.Render(text)
	}

	return docStyle.Render(m.list.View())
}

func main() {
	m := Model{list: list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "My Tasks"

	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
