package actions

import (
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type Model struct {
	frame lipgloss.Style
	list  list.Model
}

func (m *Model) setStyles() {
	m.frame = lipgloss.NewStyle().MarginTop(1)
}

func (m *Model) SetSize(width, height int) {
	w, h := m.frame.GetFrameSize()
	m.list.SetSize(width-w, height-h)
}

func New() *Model {
	items := []list.Item{
		item{
			title: "Sync repo",
			desc:  "Run git commands to get latest changes palle palle palle palle palle palle",
		},
		item{title: "Sync repo"},
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.SetShowHelp(false)
	l.SetShowTitle(false)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	m := &Model{
		list: l,
	}

	m.setStyles()

	return m
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return cmd
}

func (m *Model) View() string {
	return m.frame.Render(m.list.View())
}
