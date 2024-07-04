package main

import (
	"fmt"
	"os"

	"github.com/burkel24/task-app/tui/app"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m := app.NewAppModel()

	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "")
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}

		defer f.Close()
	}

	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
