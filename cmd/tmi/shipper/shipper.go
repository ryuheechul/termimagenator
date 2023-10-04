package shipper

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ryuheechul/termimagenator/cmd/tmi/plan"
	"github.com/ryuheechul/termimagenator/pkg/bubbles/multiselect"
	"github.com/samber/lo"
)

type Model struct {
	ms      multiselect.Model
	errDesc string
}

func (m Model) Init() tea.Cmd { return nil }

type DeletionReqMsg struct {
	Plan plan.Plan
}

func (m Model) requestDeletion() tea.Msg {
	rows := m.ms.SelectedRows()

	return DeletionReqMsg{
		plan.InitialPlan(lo.Map(rows, func(row table.Row, index int) string {
			imageId, repoAndTag := row[0], row[1]

			if strings.Contains(repoAndTag, "<none>") {
				return imageId
			}

			return repoAndTag
		})),
	}
}

func (m *Model) ClearSelected() {
	m.ms.ClearSelected()
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Delete):
			if len(m.ms.SelectedRows()) == 0 {

				m.errDesc = "! nothing selected to delete"
				return m, cmd
			}
			return m, m.requestDeletion
		}
	}

	m.errDesc = ""
	m.ms, cmd = m.ms.Update(msg)
	return m, cmd
}

var errStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("9"))

func (m Model) View() string {
	v := strings.Join([]string{
		m.ms.View(),
		renderHelp(),
	}, "\n")

	if len(m.errDesc) > 0 {
		return strings.Join([]string{
			fmt.Sprintf(" %s", errStyle.Render(m.errDesc)),
			v,
		}, "\n")
	}

	return v
}

func InitialModel(columns []table.Column, rows []table.Row) Model {
	return Model{multiselect.New(
		[]table.Option{
			table.WithColumns(columns),
		},
		"images to be deleted or untagged",
		[]table.Option{
			// table.WithRows(rows),
		},
		"images to be NOT deleted",
		[]table.Option{
			table.WithRows(rows),
		},
	), ""}
}
