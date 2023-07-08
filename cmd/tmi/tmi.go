package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ryuheechul/termimagenator/pkg/image/ls"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type result string

func list() tea.Msg {
	return result(strings.Join(ls.ListWithDefaultFormat(), "\n"))
}

func off(p *tea.Program) {
	p.Send(list())
}

type model struct {
	spinner spinner.Model
	result  string
	err     error
}

type errMsg error

func initialModel() model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return model{spinner: s}
}

func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case result:
		m.result = string(msg)
		return m, tea.Quit
	case errMsg:
		m.err = msg
		return m, nil
	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m model) View() string {
	if m.err != nil {
		return m.err.Error()
	}

	if m.result != "" {
		return m.result
	}

	return fmt.Sprintf("\n\n %s listing images", m.spinner.View())
}

func main() {
	p := tea.NewProgram(initialModel())

	go off(p)
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
