package confirmer

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

type keyMap struct {
	Yes key.Binding
	No  key.Binding
}

var keys = keyMap{
	Yes: key.NewBinding(
		key.WithKeys("y"),
		key.WithHelp("y", "confirm to delete"),
	),
	No: key.NewBinding(
		key.WithKeys("n"),
		key.WithHelp("n", "cancel"),
	),
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part
// of the help.KeyMap interface.
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Yes,
		k.No,
	}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// help.KeyMap interface.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Yes, k.No},
	}
}

var helpView string

func renderHelp() string {
	if len(helpView) == 0 {
		h := help.New()
		h.ShowAll = true
		helpView = h.View(keys)
	}

	return helpView
}
