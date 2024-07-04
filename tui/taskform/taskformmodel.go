package taskform

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/burkel24/task-app/pkg/tasks"
	"github.com/burkel24/task-app/tui/api"
	"github.com/burkel24/task-app/tui/common"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	hotPink  = lipgloss.Color("#FF06B7")
	darkGray = lipgloss.Color("#767676")
)

var (
	docStyle = lipgloss.NewStyle().Margin(1, 2)

	inputLabelStyle = lipgloss.NewStyle().
			Foreground(hotPink)

	inputStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(hotPink)

	formFieldWrapperStyle = lipgloss.NewStyle().Padding(0)
)

const (
	summaryInputIdx = iota
	notesInputIdx
)

type formKeyMap struct {
	Help key.Binding
	Exit key.Binding
	Save key.Binding
	Next key.Binding
}

func (k formKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Save, k.Exit}
}

func (k formKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Help, k.Save, k.Exit},
	}
}

var keys = formKeyMap{
	Help: key.NewBinding(
		key.WithKeys("ctrl+h"),
		key.WithHelp("ctrl+h", "help"),
	),
	Exit: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "exit"),
	),
	Save: key.NewBinding(
		key.WithKeys("ctrl+s"),
		key.WithHelp("ctrl+s", "save"),
	),
	Next: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "next field"),
	),
}

type TaskFormModel struct {
	client    *api.Client
	keys      formKeyMap
	help      help.Model
	height    int
	isNewTask bool
	summary   textarea.Model
	notes     textarea.Model
	focused   int
	task      tasks.TaskDTO
}

func newFormField(label string, height int) textarea.Model {
	input := textarea.New()
	input.Placeholder = label
	input.ShowLineNumbers = false
	input.Prompt = ""

	input.MaxWidth = 0
	input.FocusedStyle.CursorLine = lipgloss.NewStyle()

	input.SetWidth(50)
	input.SetHeight(height)

	return input
}

func NewTaskFormModel(client *api.Client) TaskFormModel {
	summaryInput := newFormField("Summary", 2)
	notesInput := newFormField("Notes", 5)

	return TaskFormModel{
		client:  client,
		keys:    keys,
		help:    help.New(),
		summary: summaryInput,
		height:  0,
		notes:   notesInput,
		focused: summaryInputIdx,
	}
}

func (m *TaskFormModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *TaskFormModel) Update(msg tea.Msg) (TaskFormModel, tea.Cmd) {
	var (
		summaryCmd tea.Cmd
		notesCmd   tea.Cmd
	)

	slog.Info("TaskFormModel.Update", "msg", msg)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		_, v := docStyle.GetFrameSize()
		m.height = msg.Height - v

		slog.Info("TaskFormModel.Update: tea.WindowSizeMsg")
		m.help.Width = msg.Width

	case tea.KeyMsg:
		switch {

		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll

		case key.Matches(msg, m.keys.Exit):
			return *m, common.AppStateCmd(common.AppStateTaskList)

		case key.Matches(msg, m.keys.Save):
			return *m, m.submitTask()

		case key.Matches(msg, m.keys.Next):
			if m.focused == summaryInputIdx {
				m.summary.Blur()
				m.notes.Focus()
				m.focused = notesInputIdx
			} else {
				m.notes.Blur()
				m.summary.Focus()
				m.focused = summaryInputIdx
			}
		}

	case common.CreateTaskMsg:
		m.onTaskCreate()

	case common.SelectTaskMsg:
		m.onTaskSelect(msg.Task)
		slog.Info("TaskFormModel.Update: common.SelectTaskMsg", "state", m)
	}

	m.summary, summaryCmd = m.summary.Update(msg)
	m.notes, notesCmd = m.notes.Update(msg)

	return *m, tea.Batch(
		summaryCmd,
		notesCmd,
	)
}

func (m *TaskFormModel) View() string {
	summary := lipgloss.JoinVertical(
		lipgloss.Left,
		inputLabelStyle.Width(30).Render("Summary"),
		inputStyle.Render(m.summary.View()),
	)

	notes := lipgloss.JoinVertical(
		lipgloss.Left,
		inputLabelStyle.Width(30).Render("Notes"),
		inputStyle.Render(m.notes.View()),
	)

	form := lipgloss.JoinVertical(
		lipgloss.Left,
		formFieldWrapperStyle.Render(summary),
		formFieldWrapperStyle.Render(notes),
	)

	help := m.help.View(m.keys)
	availHeight := m.height - lipgloss.Height(form) - lipgloss.Height(help)
	slog.Info(fmt.Sprintf("%d", availHeight))

	return docStyle.Render(lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.NewStyle().Height(availHeight).Render(form),
		help,
	))
}

func (m *TaskFormModel) onTaskCreate() {
	slog.Info("Creating new task")

	m.isNewTask = true
	m.task = tasks.TaskDTO{
		Summary: "",
		Notes:   "",
	}

	m.setFormStateFromModel()
}

func (m *TaskFormModel) onTaskSelect(task tasks.TaskDTO) {
	slog.Info("Editing task", "task", task)

	m.isNewTask = false
	m.task = task

	m.setFormStateFromModel()
}

func (m *TaskFormModel) setFormStateFromModel() {
	m.summary.SetValue(m.task.Summary)
	m.notes.SetValue(m.task.Notes)

	m.notes.Blur()
	m.summary.Focus()
}

func (m *TaskFormModel) submitTask() tea.Cmd {
	summary := m.summary.Value()
	notes := m.notes.Value()

	// TODO validation

	var err error
	if m.isNewTask {
		err = m.createTask(summary, notes)
	} else {
		err = m.updateTask(summary, notes)
	}

	if err != nil {
		return common.NewErrorMsg(err)
	}

	return tea.Sequence(
		common.NewTaskCmd(),
		common.NewRefreshListMsg(),
		common.AppStateCmd(common.AppStateTaskList),
	)
}

func (m *TaskFormModel) createTask(summary, notes string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	dto := tasks.CreateTaskRequestDto{
		Summary: summary,
		Notes:   notes,
	}

	_, err := m.client.CreateTask(ctx, &dto)
	if err != nil {
		return fmt.Errorf("error creating task: %w", err)
	}

	return nil
}

func (m *TaskFormModel) updateTask(summary, notes string) error {
	return fmt.Errorf("not implemented")
}
