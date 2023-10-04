package loader

import (
	"fmt"
	"strings"

	"github.com/ryuheechul/termimagenator/pkg/image/ls"
	"github.com/samber/lo"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type result struct {
	Columns []table.Column
	Rows    []table.Row
}

func load() tea.Msg {
	images, err := ls.ListWithDefaultFormat()
	if err != nil {
		panic(err)
	}

	columns := []table.Column{
		{Title: "id", Width: 12},
		{Title: "image", Width: 100},
	}

	rows := lo.Map(images, func(text string, index int) table.Row {
		return lo.Reverse(strings.Split(text, " "))
	})

	return result{Columns: columns, Rows: rows}
}

type Model struct {
	spinner spinner.Model
	Fetched result
	err     error
}

type errMsg error

func InitialModel() Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return Model{spinner: s}
}

// this is only needed when loader.Model is at top level
func (m Model) Init() tea.Cmd {
	return tea.Sequence(m.spinner.Tick, load)
}

type FetchedMsg struct{}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case result:
		m.Fetched = result(msg)
		return m, func() tea.Msg {
			return FetchedMsg{}
		}
	case errMsg:
		m.err = msg
		return m, nil
	}
	var cmd tea.Cmd

	m.spinner, cmd = m.spinner.Update(msg)

	return m, cmd
}

func (m Model) View() string {
	if m.err != nil {
		return m.err.Error()
	}

	if m.Fetched.Columns != nil {
		return "fetched"
	}

	return fmt.Sprintf(" %slisting images", m.spinner.View())
}
