package selectinput

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type SelectInputModel struct {
	selectInput textinput.Model

	list list.Model
}

type SelectOption interface {
	Label() string
	Value() interface{}
}

func NewSelectInputModel(title string) *SelectInputModel {
	input := textinput.New()
	input.Prompt = ""

	list := list.New([]list.Item{}, itemDelegate{}, 0, 0)
	list.Title = title

	list.SetShowHelp(false)
	list.SetShowPagination(false)
	list.SetShowFilter(false)
	list.SetShowStatusBar(false)

	return &SelectInputModel{
		selectInput: input,
		list:        list,
	}
}

func (m *SelectInputModel) Init() tea.Cmd {
	return nil
}

func (m *SelectInputModel) Update(msg tea.Msg) (SelectInputModel, tea.Cmd) {
	var (
		listCmd  tea.Cmd
		inputCmd tea.Cmd
	)

	selected := m.SelectedItem()
	if selected != nil {
		m.selectInput.SetValue(selected.Label())
	}

	m.list, listCmd = m.list.Update(msg)
	m.selectInput, inputCmd = m.selectInput.Update(msg)

	return *m, tea.Batch(listCmd, inputCmd)
}

// Renders the content for the main select input
func (m *SelectInputModel) View() string {
	return m.selectInput.View()
}

// Renders the content for the select panel when it is open
func (m *SelectInputModel) ViewSelectPanel() string {
	return m.list.View()
}

func (m *SelectInputModel) SelectedItem() SelectOption {
	opt := m.list.SelectedItem()
	option, ok := opt.(SelectListOption)
	if !ok {
		return nil
	}

	return option.opt
}

func (m *SelectInputModel) Value() interface{} {
	return m.SelectedItem().Value()
}

func (m *SelectInputModel) SetOptions(opts []SelectOption) {
	newItems := make([]list.Item, len(opts))

	for i, opt := range opts {
		newItems[i] = SelectListOption{opt: opt}
	}

	m.list.SetItems(newItems)
}

func (m *SelectInputModel) SetSelected(selected interface{}) {
	for i, opt := range m.list.Items() {
		option, ok := opt.(SelectListOption)
		if !ok {
			return
		}

		if option.opt.Value() == selected {
			m.selectInput.SetValue(option.opt.Label())
			m.list.Select(i)
		}
	}
}

func (m *SelectInputModel) Focus() {

}

func (m *SelectInputModel) Blur() {
	m.selectInput.Blur()
}

func (m *SelectInputModel) SetSize(width, height int) {
	m.list.SetSize(width, height)
}
