package remover

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

type keyMap struct {
	Dismiss key.Binding
}

var keys = keyMap{
	Dismiss: key.NewBinding(
		key.WithKeys("q", "enter", " "),
		key.WithHelp("enter/space/q", "press any of these to dismiss"),
	),
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part
// of the help.KeyMap interface.
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Dismiss,
	}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// help.KeyMap interface.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Dismiss},
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
