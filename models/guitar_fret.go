package models

import (
	"fmt"

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
	switch f.DisplayType {
	case constants.DISPLAY_AS_FLAT:
		return fmt.Sprintf("━━ %-2s ━┿", f.Note.NameFlat)
	case constants.DISPLAY_AS_SHARP:
		return fmt.Sprintf("━━ %-2s ━┿", f.Note.NameSharp)
	case constants.DISPLAY_CIRCLE:
		return fmt.Sprintf("━━ %-2s ━┿", "o")
	case constants.DISPLAY_EMPTY_FRET:
		return "━━━━━━━┿"
	default:
		return "━━━━━━━┿"
	}
}
