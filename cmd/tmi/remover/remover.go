package remover

import (
	"fmt"
	"strings"

	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ryuheechul/termimagenator/cmd/tmi/plan"
	"github.com/ryuheechul/termimagenator/pkg/image/rm"
)

type Announce struct {
	success bool
}

type result struct {
	untagged []string
	deleted  []string
}

type Model struct {
	What      plan.Plan
	Conducted result
	err       error
	spinner   spinner.Model
}

func announce() tea.Msg {
	return Announce{true}
}

func (m Model) Init() tea.Cmd {
	return tea.Sequence(m.spinner.Tick, m.actuallyRemove)
}

func (m Model) actuallyRemove() tea.Msg {
	time.Sleep(2000 * time.Millisecond)

	untagged, deleted, err := rm.Remove(m.What.Images)

	if err != nil {
		panic(err)
	}

	return result{untagged, deleted}
}

type errMsg error

func InitialModel(what plan.Plan) Model {
	s := spinner.New()
	s.Spinner = spinner.Points
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return Model{What: what, spinner: s}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.isJobDone() {
			switch {
			case key.Matches(msg, keys.Dismiss):
				return m, announce
			}
		}
	case result:
		r := result(msg)
		m.Conducted = r
		m.What.UpdateResult(r.untagged, r.deleted)

		return m, nil
	case errMsg:
		m.err = msg
		return m, nil
	}

	var cmd tea.Cmd

	m.spinner, cmd = m.spinner.Update(msg)
	return m, cmd
}

func (m Model) isJobDone() bool {
	return len(m.Conducted.untagged) > 0 || len(m.Conducted.deleted) > 0
}

func (m Model) View() string {
	sv := m.spinner.View()
	title := fmt.Sprintf(" %s deleting", sv)

	help := ""
	if m.isJobDone() {
		title = "  ✅ deleted"
		help = renderHelp()
	}

	return strings.Join([]string{
		strings.Join([]string{
			title,
			m.desc(),
		}, " "),
		help,
	}, "\n")
}

func (m Model) desc() string {
	if len(m.What.Result) > 0 {

		s := lipgloss.NewStyle().MarginLeft(5)

		return strings.Join([]string{
			m.What.Desc,
			"",
			s.Render(m.What.Result),
		}, "\n")
	}

	return m.What.Desc
}
