package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ryuheechul/termimagenator/cmd/tmi/confirmer"
	"github.com/ryuheechul/termimagenator/cmd/tmi/loader"
	"github.com/ryuheechul/termimagenator/cmd/tmi/remover"
	"github.com/ryuheechul/termimagenator/cmd/tmi/shipper"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	loader    loader.Model
	shipper   *shipper.Model
	confirmer *confirmer.Model
	remover   *remover.Model
}

func (m Model) Init() tea.Cmd {
	return m.loader.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.remover == nil {
			// skip processing key press while removing
			switch {
			case key.Matches(msg, keys.Quit):
				return m, tea.Quit
			}
		}
	case remover.Announce:
		m.remover = nil
		m.shipper.ClearSelected()
	case confirmer.ConfirmedMsg:
		m.confirmer = nil
		if c := confirmer.ConfirmedMsg(msg); c.Answer == true {
			remover := remover.InitialModel(c.WhenSuccess)
			m.remover = &remover
			return m, m.remover.Init()
		}
	case shipper.DeletionReqMsg:
		confirmer := confirmer.InitialModel(shipper.DeletionReqMsg(msg).Plan)
		m.confirmer = &confirmer
	case loader.FetchedMsg:
		if m.shipper == nil {
			shipper := shipper.InitialModel(m.loader.Fetched.Columns, m.loader.Fetched.Rows)
			m.shipper = &shipper
		}
	}

	if m.remover != nil {
		return m.updateRemover(msg)
	}

	if m.confirmer != nil {
		return m.updateConfirmer(msg)
	}

	if m.shipper != nil {
		return m.updateShipper(msg)
	}

	return m.updateLoader(msg)
}

func (m Model) updateLoader(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.loader, cmd = m.loader.Update(msg)
	return m, cmd
}

func (m Model) updateRemover(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var remover remover.Model
	remover, cmd = m.remover.Update(msg)
	m.remover = &remover
	return m, cmd
}

func (m Model) updateConfirmer(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var confirmer confirmer.Model
	confirmer, cmd = m.confirmer.Update(msg)
	m.confirmer = &confirmer
	return m, cmd
}

func (m Model) updateShipper(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var shipper shipper.Model
	shipper, cmd = m.shipper.Update(msg)
	m.shipper = &shipper
	return m, cmd
}

func (m Model) View() string {
	if m.remover != nil {
		return strings.Join([]string{
			"", // margin at the top
			m.remover.View(),
		}, "\n")
	}

	var v string

	if m.confirmer != nil {
		v = m.confirmer.View()
	} else if m.shipper != nil {
		v = m.shipper.View()
	} else {
		v = m.loader.View()
	}

	return strings.Join([]string{
		"", // margin at the top
		v,
		renderHelp(),
		"", // to prevent the last line from disappearing
	}, "\n")
}

func main() {
	m := &Model{loader.InitialModel(), nil, nil, nil}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
