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
	client *api.Client

	isNewTask bool
	task      tasks.TaskDTO

	height int
	width  int

	focused int
	summary textarea.Model
	notes   textarea.Model

	keys formKeyMap
	help help.Model
}

func newFormField(label string, height int) textarea.Model {
	input := textarea.New()
	input.Placeholder = label
	input.ShowLineNumbers = false
	input.Prompt = ""

	input.MaxWidth = 0
	input.FocusedStyle.CursorLine = lipgloss.NewStyle()

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
		width:   0,
	}
}

func (m *TaskFormModel) Init() tea.Cmd {
	return textarea.Blink
}

func (m *TaskFormModel) Update(msg tea.Msg) (TaskFormModel, tea.Cmd) {
	var (
		summaryCmd tea.Cmd
		notesCmd   tea.Cmd
	)

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.onWindowSize(msg.Width, msg.Height)

	case tea.KeyMsg:
		switch {

		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll

		case key.Matches(msg, m.keys.Exit):
			return *m, common.AppStateCmd(common.AppStateTaskList)

		case key.Matches(msg, m.keys.Save):
			return *m, m.submitTask()

		case key.Matches(msg, m.keys.Next):
			m.onNextField()
		}

	case common.CreateTaskMsg:
		m.onTaskCreate()

	case common.SelectTaskMsg:
		m.onTaskSelect(msg.Task)
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

func (m *TaskFormModel) onWindowSize(width, height int) {
	h, v := docStyle.GetFrameSize()

	m.height = height - v

	availWidth := width - h

	m.summary.SetWidth(availWidth)
	m.notes.SetWidth(availWidth)

	m.width = availWidth
	m.help.Width = availWidth
}

func (m *TaskFormModel) onNextField() {
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
		slog.Error("Error submitting task", "error", err)
		return common.NewErrorMsg(err)
	}

	return tea.Sequence(
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	dto := tasks.UpdateTaskRequestDto{
		Summary: summary,
		Notes:   notes,
	}

	_, err := m.client.UpdateTask(ctx, m.task.ID, &dto)
	if err != nil {
		return fmt.Errorf("error updating task: %w", err)
	}

	return nil
}
