package forms

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/lipgloss"
)

func NewFormField(label string, height int) textarea.Model {
	input := textarea.New()
	input.Placeholder = label
	input.ShowLineNumbers = false
	input.Prompt = ""

	input.MaxWidth = 0
	input.FocusedStyle.CursorLine = lipgloss.NewStyle()

	input.SetHeight(height)

	return input
}
