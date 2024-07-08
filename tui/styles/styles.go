package styles

import "github.com/charmbracelet/lipgloss"

var (
	PageWrapperStyle = lipgloss.NewStyle().Margin(1, 2)

	InputLabelStyle = lipgloss.NewStyle().
			Foreground(HotPink)

	InputStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(HotPink)

	FormFieldWrapperStyle = lipgloss.NewStyle().Padding(0).Margin(0)

	BorderStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(HotPink)
)
