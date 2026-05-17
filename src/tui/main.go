package tui

import (
	"fmt"
	"os"

	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	tint "github.com/lrstanley/bubbletint/v2"
)

type ThemeChangedMsg struct{}

type model struct {
	width  int
	height int

	spinner spinner.Model
	logo    *logo
}

func initialModel() *model {
	m := &model{
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

	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "l":
			tint.NextTint()
			return m, func() tea.Msg { return ThemeChangedMsg{} }
		case "h":
			tint.PreviousTint()
			return m, func() tea.Msg { return ThemeChangedMsg{} }
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
	var view tea.View

	if m.width == 0 {
		return view
	}

	loadingTitle := lipgloss.JoinHorizontal(lipgloss.Left, m.spinner.View(), "Loading...")

	content := lipgloss.JoinVertical(lipgloss.Left, m.logo.View(), loadingTitle)
	view.SetContent(content)

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
