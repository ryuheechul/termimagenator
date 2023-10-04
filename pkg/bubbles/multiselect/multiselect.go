package multiselect

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("8"))

type styleSelected struct {
	fg lipgloss.TerminalColor
	bg lipgloss.TerminalColor
}

func tableStyle(selected *styleSelected) table.Styles {
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("8")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(selected.fg).
		Background(selected.bg).
		Bold(false)

	return s
}

var focusedStyle = tableStyle(&styleSelected{
	fg: lipgloss.Color("0"),
	bg: lipgloss.Color("4"),
})
var blurredStyle = tableStyle(&styleSelected{
	fg: lipgloss.Color("0"),
	bg: lipgloss.Color("8"),
})

type Model struct {
	selectedTitle   string
	unselectedTitle string
	selectedTable   table.Model
	unselectedTable table.Model
	tfr             TableFocusRotator
	help            help.Model
}

func (m Model) SelectedRows() []table.Row {
	return m.selectedTable.Rows()
}

func (m *Model) ClearSelected() {
	m.selectedTable.SetRows([]table.Row{})
	m.focusThis(&m.unselectedTable)
}

func (m *Model) getFocused() *table.Model {
	focused, _ := m.getStreamEnds()
	return focused
}

func (m *Model) focusThis(this *table.Model) {
	m.tfr.focusThis([]*table.Model{&m.selectedTable, &m.unselectedTable}, this)
}

func (m *Model) focusNext() {
	m.tfr.focusNext([]*table.Model{&m.selectedTable, &m.unselectedTable})
}

func (m *Model) getStreamEnds() (*table.Model, *table.Model) {
	focused := &m.unselectedTable
	unfocused := &m.selectedTable

	if m.selectedTable.Focused() {
		focused = &m.selectedTable
		unfocused = &m.unselectedTable
	}

	return focused, unfocused
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Switch):
			m.focusNext()
			return m, cmd
		case key.Matches(msg, keys.Move):
			from, to := m.getStreamEnds()
			moveRow(from, to)

			if len(from.SelectedRow()) == 0 {
				m.focusNext()
			}
			return m, cmd
		case key.Matches(msg, keys.Bottom):
			m.getFocused().GotoBottom()
			return m, cmd
		case key.Matches(msg, keys.Top):
			m.getFocused().GotoTop()
			return m, cmd
		case key.Matches(msg, keys.AllSelected):
			moveAll(&m.unselectedTable, &m.selectedTable)
			m.focusThis(&m.selectedTable)
			return m, cmd
		case key.Matches(msg, keys.AllUnselected):
			moveAll(&m.selectedTable, &m.unselectedTable)
			m.focusThis(&m.unselectedTable)
			return m, cmd
		}
	}

	m.selectedTable, cmd = m.selectedTable.Update(msg)
	m.unselectedTable, cmd = m.unselectedTable.Update(msg)
	return m, cmd
}

type title struct {
	tag   string
	style lipgloss.Style
}

var selectedTitle = title{
	tag:   "selected",
	style: lipgloss.NewStyle().Bold(true).Background(lipgloss.Color("5")).Foreground(lipgloss.Color("0")),
}

var unselectedTitle = title{
	tag:   "unselected",
	style: lipgloss.NewStyle().Bold(false).Background(lipgloss.Color("8")).Foreground(lipgloss.Color("0")),
}

func renderTitle(text string, t title, tbl table.Model) string {
	highlighted := t.style.Render(fmt.Sprintf(
		" %d %s ",
		len(tbl.Rows()),
		t.tag,
	))

	return fmt.Sprintf(
		" %s %s:",
		highlighted,
		text,
	)
}

func (m Model) View() string {
	selected := renderTitle(m.selectedTitle, selectedTitle, m.selectedTable)
	unselected := renderTitle(m.unselectedTitle, unselectedTitle, m.unselectedTable)

	rendered := strings.Join([]string{
		selected,
		baseStyle.Render(m.selectedTable.View()),
		unselected,
		baseStyle.Render(m.unselectedTable.View()),
	}, "\n")

	return strings.Join([]string{
		rendered,
		renderHelp(),
	}, "\n")
}

func New(
	forCommon []table.Option,
	selectedTitle string,
	forSelected []table.Option,
	unselectedTitle string,
	forUnselected []table.Option,
) Model {

	optsSelected := append([]table.Option{
		table.WithFocused(true),
		table.WithHeight(7),
	}, append(forCommon, forSelected...)...)

	opsUnselected := append([]table.Option{
		table.WithFocused(false),
		table.WithHeight(7),
		table.WithRows([]table.Row{}),
	}, append(forCommon, forUnselected...)...)

	selected := table.New(
		optsSelected...,
	)

	unselected := table.New(
		opsUnselected...,
	)

	tfr := NewTableFocusRotator(&focusedStyle, &blurredStyle)
	h := help.New()
	h.ShowAll = true

	m := Model{selectedTitle, unselectedTitle, selected, unselected, tfr, h}

	if len(selected.Rows()) > 0 {
		m.focusThis(&m.selectedTable)
	} else {
		m.focusThis(&m.unselectedTable)
	}
	return m
}

func moveRow(from *table.Model, to *table.Model) {
	givingRows := from.Rows()
	target := from.SelectedRow()
	cursor := from.Cursor()
	left := givingRows[:cursor]
	right := givingRows[cursor+1:]

	if cursor == len(givingRows)-1 {
		from.SetCursor(cursor - 1)
	}

	from.SetRows(append(left, right...))

	receivingRows := to.Rows()

	receivingRows = append(receivingRows, target)

	to.SetRows(receivingRows)
	to.GotoBottom()
}

func moveAll(from *table.Model, to *table.Model) {
	givingRows := from.Rows()

	from.SetRows([]table.Row{})

	receivingRows := to.Rows()
	receivingRows = append(receivingRows, givingRows...)
	to.SetRows(receivingRows)
	to.GotoBottom()
}
