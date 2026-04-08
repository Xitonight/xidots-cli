package logo

import (
	"math/rand"
	"strings"

	"charm.land/lipgloss/v2"
)

type letterform func(bool) string

// Opts are the options for rendering the title art.
type Opts struct {
	Spacing int
}

// Render returns the stylized logo.
func Render(o Opts) string {
	letterforms := []letterform{
		letterX,
		letterI,
		letterD,
		letterO,
		letterT,
		letterS,
		letterHyphen,
		letterC,
		letterL,
		letterI,
	}

	stretchIndex := rand.Intn(len(letterforms))
	return renderWord(o.Spacing, stretchIndex, letterforms...)
}

func renderWord(spacing int, stretchIndex int, letterforms ...letterform) string {
	if spacing < 0 {
		spacing = 0
	}

	renderedLetterforms := make([]string, len(letterforms))
	for i, letter := range letterforms {
		renderedLetterforms[i] = letter(i == stretchIndex)
	}

	res := renderedLetterforms[0]
	for i := 1; i < len(renderedLetterforms); i++ {
		if spacing > 0 {
			spaceCol := strings.Repeat(" ", spacing)
			spacePart := spaceCol + "\n" + spaceCol + "\n" + spaceCol
			res = lipgloss.JoinHorizontal(lipgloss.Top, res, spacePart)
		}
		res = lipgloss.JoinHorizontal(lipgloss.Top, res, renderedLetterforms[i])
	}
	return strings.TrimSpace(res)
}

func joinLetterform(letters ...string) string {
	res := letters[0]
	for i := 1; i < len(letters); i++ {
		res = lipgloss.JoinHorizontal(lipgloss.Top, res, letters[i])
	}
	return res
}

type letterformProps struct {
	width      int
	minStretch int
	maxStretch int
	stretch    bool
}

func stretchLetterformPart(s string, p letterformProps) string {
	if p.maxStretch < p.minStretch {
		p.minStretch, p.maxStretch = p.maxStretch, p.minStretch
	}
	n := p.width
	if p.stretch {
		n = rand.Intn(p.maxStretch-p.minStretch+1) + p.minStretch
	}

	lines := strings.Split(s, "\n")
	res := make([]string, len(lines))
	for i, line := range lines {
		res[i] = strings.Repeat(line, n)
	}
	return strings.Join(res, "\n")
}

func letterX(stretch bool) string {
	return "█ █\n▄▀▄\n▀ ▀"
}

func letterI(stretch bool) string {
	return "█\n█\n▀"
}

func letterD(stretch bool) string {
	left := "█\n█\n▀"
	center := "▀\n \n▀"
	right := "▄\n█\n "
	return joinLetterform(
		left,
		stretchLetterformPart(center, letterformProps{
			stretch:    stretch,
			width:      2,
			minStretch: 4,
			maxStretch: 8,
		}),
		right,
	)
}

func letterO(stretch bool) string {
	side := "▄\n█\n "
	center := "▀\n \n▀"
	return joinLetterform(
		side,
		stretchLetterformPart(center, letterformProps{
			stretch:    stretch,
			width:      2,
			minStretch: 4,
			maxStretch: 8,
		}),
		side,
	)
}

func letterT(stretch bool) string {
	top := "▀\n \n "
	center := "█\n█\n▀"
	return joinLetterform(
		stretchLetterformPart(top, letterformProps{
			stretch:    stretch,
			width:      2,
			minStretch: 3,
			maxStretch: 5,
		}),
		center,
		stretchLetterformPart(top, letterformProps{
			stretch:    stretch,
			width:      2,
			minStretch: 3,
			maxStretch: 5,
		}),
	)
}

func letterS(stretch bool) string {
	left := "▄\n▀\n▀"
	center := "▀\n▀\n▀"
	right := "▀\n█\n "
	return joinLetterform(
		left,
		stretchLetterformPart(center, letterformProps{
			stretch:    stretch,
			width:      3,
			minStretch: 5,
			maxStretch: 8,
		}),
		right,
	)
}

func letterHyphen(stretch bool) string {
	center := " \n▀\n "
	return stretchLetterformPart(center, letterformProps{
		stretch:    stretch,
		width:      3,
		minStretch: 4,
		maxStretch: 6,
	})
}

func letterC(stretch bool) string {
	left := "▄\n█\n "
	right := "▀\n \n▀"
	return joinLetterform(
		left,
		stretchLetterformPart(right, letterformProps{
			stretch:    stretch,
			width:      3,
			minStretch: 5,
			maxStretch: 8,
		}),
	)
}

func letterL(stretch bool) string {
	left := "█\n█\n▀"
	right := " \n \n▀"
	return joinLetterform(
		left,
		stretchLetterformPart(right, letterformProps{
			stretch:    stretch,
			width:      3,
			minStretch: 5,
			maxStretch: 8,
		}),
	)
}
