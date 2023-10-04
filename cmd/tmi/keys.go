package main

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

type keyMap struct {
	Quit key.Binding
}

var keys = keyMap{
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part
// of the help.KeyMap interface.
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Quit,
	}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// help.KeyMap interface.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Quit},
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
