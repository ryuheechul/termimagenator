package multiselect

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

type keyMap struct {
	Up            key.Binding
	Down          key.Binding
	Top           key.Binding
	Bottom        key.Binding
	AllSelected   key.Binding
	AllUnselected key.Binding
	Switch        key.Binding
	Move          key.Binding
}

var keys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	Top: key.NewBinding(
		key.WithKeys("t"),
		key.WithHelp("t", "move cursor to top"),
	),
	Bottom: key.NewBinding(
		key.WithKeys("b"),
		key.WithHelp("b", "move cursor to bottom"),
	),
	AllSelected: key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "move everything to selected"),
	),
	AllUnselected: key.NewBinding(
		key.WithKeys("u"),
		key.WithHelp("u", "move everything to unselected"),
	),
	Switch: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "rotate focus"),
	),
	Move: key.NewBinding(
		key.WithKeys("enter", " "),
		key.WithHelp("enter/space", "move a row at cursor to (un)selected"),
	),
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part
// of the help.KeyMap interface.
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Move,
		k.Switch,
		k.Top,
		k.Bottom,
	}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// help.KeyMap interface.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Top, k.Bottom},
		{k.Move, k.Switch, k.AllSelected, k.AllUnselected},
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
