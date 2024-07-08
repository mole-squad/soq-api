package selectinput

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(SelectListOption)
	if !ok {
		return
	}

	if index == m.Index() {
		fmt.Fprintf(w, "> %s", i.Label())
	} else {
		fmt.Fprintf(w, "  %s", i.Label())
	}
}
