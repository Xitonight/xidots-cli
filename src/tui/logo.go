package tui

import (
	"image/color"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	tint "github.com/lrstanley/bubbletint/v2"
)

var logoStr = strings.TrimSpace(`
▄▄ ▄▄ ▄▄ ▄▄▄▄   ▄▄▄ ▄▄▄▄▄▄ ▄▄▄▄      ▄▄▄▄ ▄▄    ▄▄
▀█▄█▀ ██ ██▀██ ██▀██  ██  ███▄▄ ▄▄▄ ██▀▀▀ ██    ██
██ ██ ██ ████▀ ▀███▀  ██  ▄▄██▀     ▀████ ██▄▄▄ ██
`)

type logo struct {
	gradient []color.Color

	width  int
	height int
}

func newLogo() *logo {
	w, h := lipgloss.Size(strings.TrimSpace(logoStr))
	l := &logo{
		width:  w,
		height: h,
	}
	l.setStyles()
	return l
}

func (l *logo) setStyles() {
	l.gradient = lipgloss.Blend2D(
		l.width,
		l.height,
		180,
		tint.Current().Purple,
		tint.Current().Blue,
	)
}

func (l *logo) Init() tea.Cmd {
	return nil
}

func (l *logo) Update(msg tea.Msg) tea.Cmd {
	switch msg.(type) {
	case ThemeChangedMsg:
		l.setStyles()
	}
	return nil
}

func (l *logo) View() string {
	lines := strings.Split(logoStr, "\n")

	buf := strings.Builder{}
	for y := range l.height {
		runes := []rune(lines[y])
		for x := range l.width {
			buf.WriteString(
				lipgloss.NewStyle().
					Foreground(l.gradient[y*l.width+x]).
					Render(string(runes[x])),
			)
		}
		if y < l.height-1 { // End of row.
			buf.WriteString("\n")
		}
	}

	return buf.String()
}
