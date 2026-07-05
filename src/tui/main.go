package tui

import (
	"fmt"
	"os"
	"xidots-cli/src/tui/actions"
	"xidots-cli/src/tui/logo"
	"xidots-cli/src/tui/theme"
	"xidots-cli/src/tui/tuihelp"

	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	tint "github.com/lrstanley/bubbletint/v2"
)

type model struct {
	keys tuihelp.KeyMap
	help help.Model

	width  int
	height int

	spinner spinner.Model
	logo    *logo.Model
	actions *actions.Model

	quitting bool
}

func initialModel() *model {
	m := &model{
		keys:    tuihelp.Keys,
		help:    help.New(),
		spinner: spinner.New(spinner.WithSpinner(spinner.Dot)),
		logo:    logo.New(),
		actions: actions.New(),
	}

	return m
}

func (m model) propagate(msg tea.Msg) tea.Cmd {
	cmds := []tea.Cmd{
		m.logo.Update(msg),
		m.actions.Update(msg),
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
		m.actions.SetSize(
			msg.Width,
			msg.Height-m.logo.GetHeight()-lipgloss.Height(m.help.View(m.keys)),
		)

	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, m.keys.Left):
			tint.PreviousTint()
			return m, theme.ThemeChanged
		case key.Matches(msg, m.keys.Right):
			tint.NextTint()
			return m, theme.ThemeChanged
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
			m.actions.SetSize(
				m.width,
				m.height-m.logo.GetHeight()-lipgloss.Height(m.help.View(m.keys)),
			)
		case key.Matches(msg, m.keys.Quit):
			m.quitting = true
			return m, tea.Quit
		}

	case spinner.TickMsg:
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

	logo := m.logo.View()
	actions := m.actions.View()
	helpView := m.help.View(m.keys)

	content := lipgloss.JoinVertical(lipgloss.Left, logo, actions, helpView)

	// height := m.height - 1 - strings.Count(content, "\n") - strings.Count(helpView, "\n")
	view.SetContent(content)
	return view
}

func Run() {
	theme.Init()

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
