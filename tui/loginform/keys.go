package loginform

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Next   key.Binding
	Quit   key.Binding
	Submit key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Next, k.Submit, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Next, k.Submit, k.Quit},
	}
}

var keys = keyMap{
	Next: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "next"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "esc"),
		key.WithHelp("esc", "quit"),
	),
	Submit: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "submit"),
	),
}
