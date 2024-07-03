package tasklist

import (
	"encoding/json"
	"io"
	"net/http"
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

type taskLoadMsg struct {
	tasks []tasks.TaskDTO
}

type errMsg struct {
	error
}

type TaskListModel struct {
	error
	tasks []tasks.TaskDTO
	list  list.Model
}

func NewTaskListModel() TaskListModel {
	list := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	list.Title = "My Tasks"

	return TaskListModel{
		list: list,
	}
}

func (t TaskListModel) getTasks() tea.Msg {
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

	return taskLoadMsg{tasks: tasksResp}
}

func (t TaskListModel) Init() tea.Cmd {
	return t.getTasks
}

func (t TaskListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		t.list.SetSize(msg.Width-h, msg.Height-v)

	case taskLoadMsg:
		t.tasks = msg.tasks

		newItems := make([]list.Item, len(t.tasks))

		for i, task := range t.tasks {
			newItems[i] = TaskListItem{task: task}
		}

		t.list.SetItems(newItems)

	case errMsg:
		t.error = msg.error

		return t, nil
	}

	var cmd tea.Cmd
	t.list, cmd = t.list.Update(msg)

	return t, cmd
}

func (t TaskListModel) View() string {
	if t.error != nil {
		return helpStyle.Render(t.error.Error())
	}

	return docStyle.Render(t.list.View())
}
