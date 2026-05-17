package tui

import (
	"fmt"
	"os"
	"strings"

	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	tint "github.com/lrstanley/bubbletint/v2"
)

type ThemeChangedMsg struct{}

type model struct {
	keys keyMap
	help help.Model

	width  int
	height int

	spinner spinner.Model
	logo    *logo

	quitting bool
}

func initialModel() *model {
	m := &model{
		keys:    keys,
		help:    help.New(),
		spinner: spinner.New(spinner.WithSpinner(spinner.Dot)),
		logo:    newLogo(),
	}

	return m
}

func (m model) propagate(msg tea.Msg) tea.Cmd {
	cmds := []tea.Cmd{
		m.logo.Update(msg),
	}
	return tea.Batch(cmds...)
}

func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.help.SetWidth(msg.Width)

	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, m.keys.Left):
			tint.PreviousTint()
			return m, func() tea.Msg { return ThemeChangedMsg{} }
		case key.Matches(msg, m.keys.Right):
			tint.NextTint()
			return m, func() tea.Msg { return ThemeChangedMsg{} }
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Quit):
			m.quitting = true
			return m, tea.Quit
		}

	case ThemeChangedMsg:
		return m, m.propagate(msg)

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	return m, m.propagate(msg)
}

func (m model) View() tea.View {
	if m.quitting {
		return tea.NewView("Quitting...\n")
	}

	var view tea.View
	view.AltScreen = true

	if m.width == 0 {
		return view
	}

	logo := lipgloss.NewStyle().PaddingBottom(1).Render(m.logo.View())
	loadingTitle := lipgloss.JoinHorizontal(lipgloss.Left, m.spinner.View(), "Loading...")
	helpView := m.help.View(m.keys)

	content := lipgloss.JoinVertical(lipgloss.Left, logo, loadingTitle)

	height := m.height - 1 - strings.Count(content, "\n") - strings.Count(helpView, "\n")
	view.SetContent(content + strings.Repeat("\n", height) + helpView)

	return view
}

func Run() {
	tint.NewDefaultRegistry()
	tint.SetTint(tint.TintDracula)

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
