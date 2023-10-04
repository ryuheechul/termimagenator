package shipper

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

type keyMap struct {
	Delete key.Binding
}

var keys = keyMap{
	Delete: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "delete selected"),
	),
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part
// of the help.KeyMap interface.
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Delete,
	}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// help.KeyMap interface.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Delete},
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
