package models

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/pedrolucaspalma/notes-go/constants"
	"github.com/pedrolucaspalma/notes-go/sequence"
)

type fret struct {
	Number      int
	Note        *sequence.Note
	DisplayType constants.FretDisplayType
}

func NewFret(
	number int,
	note *sequence.Note,
	displayType constants.FretDisplayType,
) fret {
	return fret{
		Number:      number,
		Note:        note,
		DisplayType: displayType,
	}
}

func (f fret) String() string {
	color := lipgloss.Color(constants.COMPONENTS_COLORS.NOTE_TEXT)
	colorStyle := lipgloss.NewStyle().Foreground(color).Padding(0).Margin(0)

	stringColor := lipgloss.Color(constants.COMPONENTS_COLORS.STRING)
	stringChar := lipgloss.NewStyle().Foreground(stringColor).Render("━")

	switch f.DisplayType {
	case constants.DISPLAY_AS_FLAT:
		return fmt.Sprintf("%s%s %s %s┿",
			stringChar,
			stringChar,
			colorStyle.Render(fmt.Sprintf("%-2s", f.Note.NameFlat)), // Attention here, the note name is being written BEFORE applying color to preserve alignment
			stringChar,
		)
	case constants.DISPLAY_AS_SHARP:
		return fmt.Sprintf("%s%s %s %s┿",
			stringChar,
			stringChar,
			colorStyle.Render(fmt.Sprintf("%-2s", f.Note.NameSharp)), // Attention here, the note name is being written BEFORE applying color to preserve alignment
			stringChar,
		)

	case constants.DISPLAY_CIRCLE:
		return fmt.Sprintf("━━ %-2s ━┿", "o")
	case constants.DISPLAY_EMPTY_FRET:
		return fmt.Sprintf("%s┿", strings.Repeat(stringChar, 7))
	default:
		return fmt.Sprintf("%s┿", strings.Repeat(stringChar, 7))
	}
}
