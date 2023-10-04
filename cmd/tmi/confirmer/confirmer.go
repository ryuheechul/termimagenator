package confirmer

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ryuheechul/termimagenator/cmd/tmi/plan"
)

type Model struct {
	Plan plan.Plan
}

func (m Model) Init() tea.Cmd { return nil }

func InitialModel(plan plan.Plan) Model {
	return Model{plan}
}

type ConfirmedMsg struct {
	Answer      bool
	WhenSuccess plan.Plan
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Yes):
			return m, func() tea.Msg { return ConfirmedMsg{true, m.Plan} }
		case key.Matches(msg, keys.No):
			return m, func() tea.Msg { return ConfirmedMsg{false, m.Plan} }
		}
	}

	return m, cmd
}

func (m Model) View() string {
	return strings.Join([]string{
		fmt.Sprintf(" 🗑️  going to delete %s", m.Plan.Desc),
		"",
		" ✅ choose what to do",
		"",
		renderHelp(),
	}, "\n")
}
