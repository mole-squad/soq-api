package tasklist

import (
	"context"
	"fmt"
	"time"

	"github.com/burkel24/task-app/pkg/tasks"
	"github.com/burkel24/task-app/tui/api"
	"github.com/burkel24/task-app/tui/common"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	docStyle = lipgloss.NewStyle().Margin(1, 2)
)

type taskLoadMsg struct {
	tasks []tasks.TaskDTO
}

type listKeyMap struct {
	New    key.Binding
	Edit   key.Binding
	Delete key.Binding
}

func newListKeyMap() listKeyMap {
	return listKeyMap{
		New: key.NewBinding(
			key.WithKeys("n"),
			key.WithHelp("n", "new task"),
		),
		Edit: key.NewBinding(
			key.WithKeys("e"),
			key.WithHelp("e", "edit task"),
		),
		Delete: key.NewBinding(
			key.WithKeys("d"),
			key.WithHelp("d", "delete task"),
		),
	}
}

type TaskListModel struct {
	client *api.Client
	tasks  []tasks.TaskDTO
	keys   listKeyMap
	list   list.Model
}

func NewTaskListModel(client *api.Client) TaskListModel {
	listKeys := newListKeyMap()

	list := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	list.Title = "My Tasks"

	list.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			listKeys.New,
		}
	}

	return TaskListModel{
		client: client,
		keys:   listKeys,
		list:   list,
	}
}

func (m TaskListModel) Init() tea.Cmd {
	return nil
}

func (m TaskListModel) Update(msg tea.Msg) (TaskListModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case taskLoadMsg:
		m.tasks = msg.tasks

		newItems := make([]list.Item, len(m.tasks))

		for i, task := range m.tasks {
			newItems[i] = TaskListItem{task: task}
		}

		m.list.SetItems(newItems)

	case common.RefreshListMsg:
		return m, m.getTasks

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.New):
			return m, tea.Sequence(
				common.NewCreateTaskMsg(),
				common.AppStateCmd(common.AppStateTaskForm),
			)

		case key.Matches(msg, m.keys.Edit):
			selected := m.list.SelectedItem()
			if selected == nil {
				return m, common.NewErrorMsg(fmt.Errorf("no task selected"))
			}

			taskItem, ok := selected.(TaskListItem)
			if !ok {
				return m, common.NewErrorMsg(fmt.Errorf("unexpected task item type"))
			}

			return m, tea.Sequence(
				common.NewRefreshFocusAreasMsg(),
				common.NewSelectTaskMsg(taskItem.task),
				common.AppStateCmd(common.AppStateTaskForm),
			)

		case key.Matches(msg, m.keys.Delete):
			selected := m.list.SelectedItem()
			if selected == nil {
				return m, common.NewErrorMsg(fmt.Errorf("no task selected"))
			}

			taskItem, ok := selected.(TaskListItem)
			if !ok {
				return m, common.NewErrorMsg(fmt.Errorf("unexpected task item type"))
			}

			err := m.client.DeleteTask(context.Background(), taskItem.task.ID)
			if err != nil {
				return m, common.NewErrorMsg(fmt.Errorf("failed to delete task: %w", err))
			}

			return m, tea.Sequence(
				common.NewRefreshListMsg(),
			)
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	return m, cmd
}

func (m TaskListModel) View() string {
	return docStyle.Render(m.list.View())
}

func (m TaskListModel) getTasks() tea.Msg {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	tasks, err := m.client.ListTasks(ctx)
	if err != nil {
		return common.ErrorMsg{Err: err}
	}

	return taskLoadMsg{tasks: tasks}
}
