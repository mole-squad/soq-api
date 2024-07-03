package main

import (
	"fmt"
	"os"

	"github.com/burkel24/task-app/tui/app"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m := app.NewAppModel()

	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
