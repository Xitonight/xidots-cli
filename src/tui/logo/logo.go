package logo

import (
	"image/color"
	"strings"
	"xidots-cli/src/tui/theme"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	tint "github.com/lrstanley/bubbletint/v2"
)

var logoStr = strings.TrimSpace(`
▄▄ ▄▄ ▄▄ ▄▄▄▄   ▄▄▄ ▄▄▄▄▄▄ ▄▄▄▄      ▄▄▄▄ ▄▄    ▄▄
▀█▄█▀ ██ ██▀██ ██▀██  ██  ███▄▄ ▄▄▄ ██▀▀▀ ██    ██
██ ██ ██ ████▀ ▀███▀  ██  ▄▄██▀     ▀████ ██▄▄▄ ██
`)

type Model struct {
	frame    lipgloss.Style
	gradient []color.Color

	width  int
	height int
}

func (m *Model) GetWidth() int {
	return m.width
}

func (m *Model) GetHeight() int {
	return m.height
}

func New() *Model {
	w, h := lipgloss.Size(strings.TrimSpace(logoStr))
	l := &Model{
		width:  w,
		height: h,
	}
	l.setStyles()
	return l
}

func (m *Model) setStyles() {
	m.frame = lipgloss.NewStyle().MarginBottom(1)
	m.gradient = lipgloss.Blend2D(
		m.width,
		m.height,
		180,
		tint.Current().Purple,
		tint.Current().Blue,
	)
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) tea.Cmd {
	switch msg.(type) {
	case theme.ThemeChangedMsg:
		m.setStyles()
	}
	return nil
}

func (m *Model) View() string {
	lines := strings.Split(logoStr, "\n")

	buf := strings.Builder{}
	for y := range m.height {
		runes := []rune(lines[y])
		for x := range m.width {
			buf.WriteString(
				lipgloss.NewStyle().
					Foreground(m.gradient[y*m.width+x]).
					Render(string(runes[x])),
			)
		}
		if y < m.height-1 { // End of row.
			buf.WriteString("\n")
		}
	}

	return m.frame.Render(buf.String())
}
