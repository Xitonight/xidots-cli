package theme

import (
	tea "charm.land/bubbletea/v2"
	tint "github.com/lrstanley/bubbletint/v2"
)

type ThemeChangedMsg struct{}

func Init() {
	tint.NewDefaultRegistry()
	tint.SetTint(tint.TintCatppuccinMocha)
}

func ThemeChanged() tea.Msg {
	return ThemeChangedMsg{}
}
